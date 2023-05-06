package usecase

import (
	"errors"
	"golang.org/x/net/context"
	"net/http"
	"reportify-backend/entity"
)

type OrganizationUseCase struct {
	OrganizationRepo
	UserRepo
}

func NewOrganizationUseCase(organizationRepo OrganizationRepo, userRepo UserRepo) *OrganizationUseCase {
	return &OrganizationUseCase{organizationRepo, userRepo}
}

func (u *OrganizationUseCase) GetOrganizations(ctx context.Context, userID string) ([]*entity.Organization, error) {
	return u.OrganizationRepo.GetOrganizations(ctx, userID)
}

func (u *OrganizationUseCase) GetOrganization(ctx context.Context, organizationCode, userID string) (*entity.Organization, error) {
	return u.OrganizationRepo.GetOrganization(ctx, organizationCode, userID)
}

func (u *OrganizationUseCase) UpdateOrganization(ctx context.Context, oldOrganizationCode, userID, organizationName, organizationCode, mission, vision, value string) (*entity.Organization, error) {
	ou, err := u.UserRepo.GetOrganizationUserRole(ctx, oldOrganizationCode, &userID, nil)
	if err != nil {
		return nil, err
	}

	// 管理者でない場合はエラー
	if ou == nil || ou.IsAdmin == false {
		return nil, entity.NewError(http.StatusForbidden, errors.New("you are not admin"))
	}

	return u.OrganizationRepo.UpdateOrganization(ctx, oldOrganizationCode, organizationName, organizationCode, mission, vision, value)
}
