package main

import (
	"encoding/json"
	"fmt"
	"github.com/fy23-gw-gackathon/reportify-backend/config"
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/driver"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
	"os"
)

func main() {
	cfg := config.Load()
	rdb := driver.NewRedisClient(cfg.Datastore.Address)
	client := driver.NewHttp()

	pubSub := rdb.Subscribe(context.Background(), driver.JobQueueKey)
	defer pubSub.Close()

	ch := pubSub.Channel()
	for msg := range ch {
		log.Println(handler(msg, client))
	}
	os.Exit(0)
}

func getMessage(payload string) *entity.PubSubMessage {
	b := []byte(payload)
	var msg *entity.PubSubMessage
	if err := json.Unmarshal(b, &msg); err != nil {
		panic(err)
	}
	return msg
}

func handler(msg *redis.Message, client *driver.Http) error {
	m := getMessage(msg.Payload)

	// TODO: ChatGPTの処理を書く

	d, err := json.Marshal(entity.ReviewReportRequest{ReviewBody: m.Body})
	if err != nil {
		return err
	}
	if _, err = client.Put(fmt.Sprintf("http://backend:8080/reports/%s", m.ID), d); err != nil {
		return err
	}
	return nil
}
