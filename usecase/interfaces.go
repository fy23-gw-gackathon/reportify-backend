package usecase

import (
	"reportify-backend/entity"

	"golang.org/x/net/context"
)

type UserRepo interface{}

type OrganizationRepo interface {
	GetOrganizations(ctx context.Context, limit *int, offset *int) ([]*entity.Organization, error)
	GetMVV(ctx context.Context, organizationId string) (*entity.MVV, error)
}

type ConversationRepo interface {
	GetPrevConversations(ctx context.Context, userID string) ([]*entity.Conversation, error)
	SendReport(ctx context.Context, userID string, mvv *entity.MVV, prevMessages []*entity.Conversation, message string) (*entity.Conversation, error)
}