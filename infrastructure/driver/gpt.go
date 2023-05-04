package driver

import (
	"fmt"
	"math"
	"os"
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

func (d *GptDriver) RequestMessage(input string) (openai.ChatCompletionResponse, error) {
	// バックオフリトライ
	retryCnt := 0
	for {
		resp, err := d.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo0301,
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
			fmt.Println(resp.Choices[0].Message.Content)
			return resp, nil
		}
		time.Sleep(time.Duration(math.Pow(2, float64(retryCnt))) * time.Second)
		retryCnt++
		fmt.Println("retrying...")
	}
}
