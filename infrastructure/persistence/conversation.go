package persistence

import (
	"reportify-backend/entity"
	"reportify-backend/infrastructure/driver"
	"reportify-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ConversationPersistence struct {
	GptDriver *driver.GptDriver
}

func NewConversationPersistence(GptDriver *driver.GptDriver) *ConversationPersistence {
	return &ConversationPersistence{
		GptDriver: GptDriver,
	}
}

func (u *ConversationPersistence) GetPrevConversations(ctx context.Context, userID string) ([]*entity.Conversation, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var records []*model.Conversation
	if err :=  db.Table("conversations").Where("user_id = ?", userID).Find(&records).Error; err != nil {
		return nil, err
	}
	var conversations []*entity.Conversation
	for _, record := range records {
		conversations = append(conversations, &entity.Conversation{
			ID:   record.ID,
			UserID: record.UserID,
			IsFromAI: record.IsFromAI,
			Content: record.Content,
			CreatedAt: record.CreatedAt,
		})
	}
	return conversations, nil
}

func (u *ConversationPersistence) SendReport(ctx context.Context, userID string,  mvv *entity.MVV, prevMessages []*entity.Conversation, message string) (*entity.Conversation, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var result entity.Conversation
	db.Transaction(func(tx *gorm.DB) error {
		resp, err := u.GptDriver.RequestMessage(prevMessages, message)
		if err != nil {
			return err
		}
		u1, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		u2, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		conversations := []model.Conversation{
			{
				ID:  u1.String(),
				UserID: userID,
				IsFromAI: false,
				Content: message,
			},
			{
				ID:  u2.String(),
				UserID: userID,
				IsFromAI: true,
				Content: resp.Choices[0].Message.Content,
			},
		}
		if err := tx.Table("conversations").Create(&conversations).Error; err != nil {
			return err
		}
		result = entity.Conversation{
			UserID: userID,
			IsFromAI: false,
			Content: resp.Choices[0].Message.Content,
		}
		return nil
	})
	return &result, nil
}
