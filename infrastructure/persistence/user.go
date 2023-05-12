package persistence

import (
	"errors"
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/driver"
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/persistence/model"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type UserPersistence struct {
	*driver.CognitoClient
}

func NewUserPersistence(client *driver.CognitoClient) *UserPersistence {
	return &UserPersistence{client}
}

func (p UserPersistence) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var record *model.User
	if err := db.Preload("OrganizationUser").First(&record, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return record.ToEntity(), nil
}

func (p UserPersistence) GetUsers(ctx context.Context, organizationID string) ([]*entity.User, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var record *model.Organization
	if err := db.
		Where("id = ?", organizationID).
		Preload("Users.OrganizationUsers").
		Find(&record).Error; err != nil {
		return nil, err
	}
	var users []*entity.User
	for _, user := range record.Users {
		users = append(users, user.ToEntity())
	}
	return users, nil
}

func (p UserPersistence) GetOrganizationUser(ctx context.Context, organizationCode string, userID string) (*entity.OrganizationUser, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var record *model.OrganizationUser
	if err := db.Where("user_id = ?", userID).Preload("User").Preload("Organization", "code = ?", organizationCode).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.NewError(http.StatusNotFound, err)
		}
		return nil, err
	}
	if record.Organization == nil {
		return nil, entity.NewError(http.StatusNotFound, errors.New("organization not found"))
	}
	return &entity.OrganizationUser{
		UserID:         record.UserID,
		UserName:       record.User.Name,
		OrganizationID: record.Organization.ID,
		IsAdmin:        record.Role == 1,
	}, nil
}

func (p UserPersistence) GetUserIDFromToken(ctx context.Context, token string) (*string, error) {
	cognitoUser, err := p.CognitoClient.GetUserFromToken(ctx, token)
	if err != nil {
		return nil, entity.NewError(http.StatusUnauthorized, err)
	}
	return &cognitoUser.UserID, nil
}

func (p UserPersistence) CreateUser(ctx context.Context, email, organizationID string) (*entity.User, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	names := strings.Split(email, "@")
	var userName string
	if len(names) >= 2 {
		userName = names[0]
	}

	var user *model.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userID := generateID().String()
			// CognitoID取得
			cognitoUser, err := p.CognitoClient.CreateUser(ctx, userID, email)
			if err != nil {
				return nil, entity.NewError(http.StatusInternalServerError, err)
			}

			// User作成
			user = &model.User{
				ID:        userID,
				Name:      userName,
				Email:     email,
				CognitoID: cognitoUser.ID,
				OrganizationUsers: []*model.OrganizationUser{
					{
						UserID:         userID,
						OrganizationID: organizationID,
						Role:           0,
					},
				},
			}
			if err := db.Create(&user).Error; err != nil {
				var mysqlErr *mysql.MySQLError
				if errors.As(err, &mysqlErr) && mysqlErr.Number == driver.ErrDuplicateEntryNumber {
					return nil, entity.NewError(http.StatusConflict, err)
				}
				return nil, err
			}
			return user.ToEntity(), nil
		}
		return nil, err
	}

	// Organizationとの紐付け
	ou := &model.OrganizationUser{
		UserID:         user.ID,
		OrganizationID: organizationID,
		Role:           0,
	}
	if err := db.Create(&ou).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == driver.ErrDuplicateEntryNumber {
			return nil, entity.NewError(http.StatusConflict, err)
		}
		return nil, err
	}

	user.OrganizationUsers = append(user.OrganizationUsers, ou)

	return user.ToEntity(), nil
}

func (p UserPersistence) UpdateUserRole(ctx context.Context, organizationID, userID string, role bool) error {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var record *model.OrganizationUser
	if err := db.Where("user_id = ? AND organization_id = ?", userID, organizationID).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.NewError(http.StatusNotFound, err)
		}
		return err
	}
	roleInt := 0
	if role {
		roleInt = 1
	}
	if err := db.Model(&record).Where("user_id = ? AND organization_id = ?", userID, organizationID).Update("role", roleInt).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == driver.ErrDuplicateEntryNumber {
			return entity.NewError(http.StatusConflict, err)
		}
		return err
	}
	return nil
}

func (p UserPersistence) DeleteUser(ctx context.Context, organizationID, userID string) error {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var record *model.OrganizationUser
	if err := db.Where("user_id = ? AND organization_id = ?", userID, organizationID).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.NewError(http.StatusNotFound, err)
		}
		return err
	}
	if err := db.Where("user_id = ? AND organization_id = ?", userID, organizationID).Delete(&record).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == driver.ErrDuplicateEntryNumber {
			return entity.NewError(http.StatusConflict, err)
		}
		return err
	}
	return nil
}
