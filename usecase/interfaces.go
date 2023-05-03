package usecase

import (
	"golang.org/x/net/context"
	"reportify-backend/entity"
)

type UserRepo interface{}

type OrganizationRepo interface {
	GetOrganizations(ctx context.Context, limit *int, offset *int) ([]*entity.Organization, error)
}
