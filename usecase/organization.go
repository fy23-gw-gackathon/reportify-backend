package usecase

import (
	"golang.org/x/net/context"
	"reportify-backend/entity"
)

type OrganizationUseCase struct {
	OrganizationRepo
}

func NewOrganizationUseCase(organizationRepo OrganizationRepo) *OrganizationUseCase {
	return &OrganizationUseCase{organizationRepo}
}

func (u *OrganizationUseCase) GetOrganizations(ctx context.Context, limit *int, offset *int) ([]*entity.Organization, error) {
	return u.OrganizationRepo.GetOrganizations(ctx, limit, offset)
}
