package driver

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"golang.org/x/net/context"
	"math"
	"time"
)

const retryLimit = 5

type GptDriver struct {
	*openai.Client
}

func NewGptDriver() *GptDriver {

	return &GptDriver{
		openai.NewClient("your token"),
	}
}

func (d *GptDriver) RequestMessage(input string) (openai.ChatCompletionResponse, error) {
	// バックオフリトライ
	retryCnt := 0.0
	for {
		resp, err := d.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleSystem,
						Content: "",
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Content: input,
					},
				},
			},
		)
		if err != nil && retryCnt < retryLimit {
			// error チェック
			fmt.Println(err)
		} else {
			return resp, nil
		}
		time.Sleep(time.Duration(math.Pow(2, retryCnt)) * time.Second)
		retryCnt++
		fmt.Println("retrying...")
	}
}
