package controller

import (
	"reportify-backend/entity"

	"golang.org/x/net/context"
)

type UserUseCase interface{}

type OrganizationUseCase interface {
	GetOrganizations(ctx context.Context, limit *int, offset *int) ([]*entity.Organization, error)
	GetMVV(ctx context.Context, organizationId string) (*entity.MVV, error)
}

type ConversationUseCase interface {
	SendReport(ctx context.Context, userID string, mvv *entity.MVV, message string) (*entity.Conversation, error)
}

