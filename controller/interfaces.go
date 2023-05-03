package controller

import (
	"golang.org/x/net/context"
	"reportify-backend/entity"
)

type UserUseCase interface{}

type OrganizationUseCase interface {
	GetOrganizations(ctx context.Context, limit *int, offset *int) ([]*entity.Organization, error)
}
