package usecase

import (
	"errors"
	"golang.org/x/net/context"
	"net/http"
	"reportify-backend/entity"
)

type UserUseCase struct {
	UserRepo
	OrganizationRepo
}

func NewUserUseCase(userRepo UserRepo, orgRepo OrganizationRepo) *UserUseCase {
	return &UserUseCase{userRepo, orgRepo}
}

func (u UserUseCase) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	return u.UserRepo.GetUser(ctx, userID)
}

func (u UserUseCase) GetUsers(ctx context.Context, organizationCode string) ([]*entity.User, error) {
	return u.UserRepo.GetUsers(ctx, organizationCode)
}

func (u UserUseCase) InviteUser(ctx context.Context, email, organizationCode, userID string) (*entity.User, error) {
	ou, err := u.UserRepo.GetOrganizationUserRole(ctx, organizationCode, nil, &email)
	if err != nil {
		return nil, err
	}

	// 管理者でない場合はエラー
	if ou == nil || ou.IsAdmin == false {
		return nil, entity.NewError(http.StatusForbidden, errors.New("you are not admin"))
	}

	// Organization取得
	org, err := u.OrganizationRepo.GetOrganization(ctx, organizationCode, userID)
	if err != nil {
		return nil, err
	}

	// User作成
	user, err := u.UserRepo.CreateUser(ctx, email, org.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
