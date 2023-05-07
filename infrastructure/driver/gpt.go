package driver

import (
	"fmt"
	"github.com/fy23-gw-gackathon/reportify-backend/config"
	"math"
	"time"

	"github.com/sashabaranov/go-openai"
	"golang.org/x/net/context"
)

const retryLimit = 5

type GptDriver struct {
	*openai.Client
}

func NewGptDriver(cfg config.Config) *GptDriver {
	return &GptDriver{
		openai.NewClient(cfg.OpenAI.Key),
	}
}

func (d *GptDriver) RequestMessage(systemPrompt string, userPrompt string) (openai.ChatCompletionResponse, error) {
	// バックオフリトライ
	retryCnt := 0
	for {
		messages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userPrompt,
		})
		resp, err := d.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo0301,
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

const (
	// ReportSystemPromptTemplate AIをメンターにさせるためのプロンプト
	ReportSystemPromptTemplate = `あなたはとある会社のチームに所属し、新人教育を担当しています。あなたのチームには新人が何人かおり、これらの新人のメンターとなりました。
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
