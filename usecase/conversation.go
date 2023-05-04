package usecase

import (
	"reportify-backend/entity"

	"golang.org/x/net/context"
)

type ConversationUseCase struct {
	ConversationRepo
}

func NewConversationUseCase(conversationRepo ConversationRepo) *ConversationUseCase {
	return &ConversationUseCase{
		ConversationRepo: conversationRepo,
	}
}

func(u *ConversationUseCase) SendReport(ctx context.Context, userID string, mvv *entity.MVV, message string) (*entity.Conversation, error){
	prevMessages, err := u.ConversationRepo.GetPrevConversations(ctx, userID)
	if err != nil {
		return nil, err
	}
	return u.ConversationRepo.SendReport(ctx, userID, mvv, prevMessages, message)
}