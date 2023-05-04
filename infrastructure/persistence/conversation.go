package persistence

import (
	"fmt"
	"reportify-backend/entity"
	"reportify-backend/infrastructure/driver"
	"reportify-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

var (
	// AIをメンターにさせるためのプロンプト
	reportSystemPromptTemplate = `あなたはとある会社のチームに所属し、新人教育を担当しています。あなたのチームには新人が何人かおり、これらの新人のメンターとなりました。
	新人がチームでうまく成長できる為に、新人がその日に行ったことと、昔と比べてどのように成長したかを把握する必要があります。あなたは、新人に日報を書いてもらい、メンターである自分からレビューしてフィードバックを返す手段をとりました。
	
	新人は下記の形式で日報を書きます:
	
	# 日報 YYYY-MM-DD
	
	## 今日やったこと
	
	## 学んだこと、感じたこと
	
	## 明日やること
	
	YYYY-MM-DDにはその日の日付が入ります。あなたは最新の日報、過去の日報、過去の日報へのフィードバックとチームのMission、Vision、Valueをもとに新人にフィードバックを返します。
	そしてあなたが所属しているチームのMission、Vision、Valueは下記のとおりです。
	Mission: %s
	Vision: %s
	Value: %s
	
	フィードバックの内容として、下記の内容を含めてください。わかりやすいようにMarkdown形式で箇条書きしてください。
	- 成長点
	- 課題点
	- 今後の成長のために何すれば良いのか
	- 明日は特にどこを注意すれば良いのか
	- 日報として何を気をつけたほうが良いのか
	
	成長点については、過去データが無い場合はなしで大丈夫です。`
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
		systemPrompt := fmt.Sprintf(reportSystemPromptTemplate, mvv.Mission, mvv.Vision, mvv.Value)
		resp, err := u.GptDriver.RequestMessage(systemPrompt, prevMessages, message)
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
