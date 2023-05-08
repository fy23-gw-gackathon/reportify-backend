package usecase

import (
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"golang.org/x/net/context"
)

type UserUseCase struct {
	UserRepo
	OrganizationRepo
}

func NewUserUseCase(userRepo UserRepo, orgRepo OrganizationRepo) *UserUseCase {
	return &UserUseCase{userRepo, orgRepo}
}

func (u UserUseCase) GetUserFromToken(ctx context.Context, token string) (*entity.User, error) {
	userID, err := u.UserRepo.GetUserIDFromToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return u.UserRepo.GetUser(ctx, *userID)
}

func (u UserUseCase) GetUsers(ctx context.Context, organizationID string) ([]*entity.User, error) {
	return u.UserRepo.GetUsers(ctx, organizationID)
}

func (u UserUseCase) InviteUser(ctx context.Context, email, organizationID string) (*entity.User, error) {
	// User作成
	user, err := u.UserRepo.CreateUser(ctx, email, organizationID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserUseCase) UpdateUserRole(ctx context.Context, organizationID, userID string, role bool) error {
	return u.UserRepo.UpdateUserRole(ctx, organizationID, userID, role)
}

func (u UserUseCase) DeleteUser(ctx context.Context, organizationID, userID string) error {
	return u.UserRepo.DeleteUser(ctx, organizationID, userID)
}
