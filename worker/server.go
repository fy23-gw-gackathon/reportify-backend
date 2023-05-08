package main

import (
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/driver"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
	"os"
)

type PubSubClient interface {
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
}

// redis.Clientを隠蔽してサーバー起動するところ
type server struct {
	PubSubClient
}

func newServer(rdb PubSubClient) *server {
	return &server{rdb}
}

func (s server) Run(f func(ctx context.Context, payload string) error) {
	ctx := context.Background()
	pubSub := s.PubSubClient.Subscribe(ctx, driver.JobQueueKey)
	defer pubSub.Close()
	ch := pubSub.Channel()
	for msg := range ch {
		log.Println("...received")
		err := f(ctx, msg.Payload)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("success")
		}
	}
	os.Exit(0)
}
