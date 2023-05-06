package persistence

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"net/http"
	"reportify-backend/entity"
	"reportify-backend/infrastructure/driver"
	"reportify-backend/infrastructure/persistence/model"
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

func (p UserPersistence) GetUsers(ctx context.Context, organizationCode string) ([]*entity.User, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var record *model.Organization
	if err := db.
		Where("code = ?", organizationCode).
		Preload("Users").
		Find(&record).Error; err != nil {
		return nil, err
	}
	var users []*entity.User
	for _, user := range record.Users {
		users = append(users, user.ToEntity())
	}
	return users, nil
}

func (p UserPersistence) GetOrganizationUserRole(ctx context.Context, organizationCode string, userID, email *string) (*entity.OrganizationUser, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var record *model.OrganizationUser
	if email == nil && userID == nil {
		return nil, entity.NewError(http.StatusNotFound, errors.New("email or userID must be provided"))
	}
	if email != nil {
		db = db.Preload("User", "email = ?", *email)
	}
	if userID != nil {
		db = db.Preload("User", "id = ?", *userID)
	}
	if err := db.Preload("Organization", "code = ?", organizationCode).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.NewError(http.StatusNotFound, err)
		}
		return nil, err
	}
	return &entity.OrganizationUser{
		UserID:         record.UserID,
		OrganizationID: record.OrganizationID,
		IsAdmin:        record.Role == 1,
	}, nil
}

func (p UserPersistence) GetUserIDFromToken(ctx context.Context, token string) (*string, error) {
	user, err := p.CognitoClient.GetUserFromToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return &user.ID, nil
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
			cognitoUser, err := p.CognitoClient.CreateUser(ctx, userID, userName, email)
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
