package driver

import (
	"fmt"
	"math"
	"os"
	"reportify-backend/entity"
	"time"

	"github.com/sashabaranov/go-openai"
	"golang.org/x/net/context"
)

const retryLimit = 5

var (
	openAIKey = os.Getenv("OPENAI_KEY")
)

type GptDriver struct {
	*openai.Client
}

func NewGptDriver() *GptDriver {
	return &GptDriver{
		openai.NewClient(openAIKey),
	}
}

func (d *GptDriver) RequestMessage(systemPrompt string, prevMessages []*entity.Conversation, userPrompt string) (openai.ChatCompletionResponse, error) {
	// バックオフリトライ
	retryCnt := 0
	for {
		messages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
		}
		// 同じセッションの前の会話を復元する。現状この方法が一般的のようだが、計算量O(N)かかるので会話が長くなるほど遅くなるのでキャッシュとか活用したい
		for _, v := range prevMessages {
			if v.IsFromAI {
				messages = append(messages, openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleAssistant,
					Content: v.Content,
				})
			}else {
				messages = append(messages, openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleUser,
					Content: v.Content,
				})
			}
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userPrompt,
		})
		resp, err := d.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo0301,
				Messages: messages,
			},
		)
		if err != nil && retryCnt < retryLimit {
			// error チェック
			fmt.Println(err)
		} else {
			fmt.Println(resp.Choices[0].Message.Content)
			return resp, nil
		}
		time.Sleep(time.Duration(math.Pow(2, float64(retryCnt))) * time.Second)
		retryCnt++
		fmt.Println("retrying...")
	}
}
