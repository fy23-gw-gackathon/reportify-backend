package usecase

import (
	"golang.org/x/net/context"
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

func (u *OrganizationUseCase) GetOrganization(ctx context.Context, organizationID string) (*entity.Organization, error) {
	return u.OrganizationRepo.GetOrganization(ctx, organizationID)
}

func (u *OrganizationUseCase) UpdateOrganization(ctx context.Context, organizationID, organizationName, organizationCode, mission, vision, value string) (*entity.Organization, error) {
	return u.OrganizationRepo.UpdateOrganization(ctx, organizationID, organizationName, organizationCode, mission, vision, value)
}
