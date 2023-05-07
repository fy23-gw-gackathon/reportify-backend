package main

import (
	"encoding/json"
	"fmt"
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/driver"
	"golang.org/x/net/context"
	"os"
)

func main() {
	rdb := driver.NewRedisClient("redis:6379")

	pubSub := rdb.Subscribe(context.Background(), driver.JobQueueKey)
	defer pubSub.Close()
	ch := pubSub.Channel()
	for msg := range ch {
		m := getMessage(msg.Payload)
		fmt.Println(m)
	}
	os.Exit(0)
}

func getMessage(payload string) *driver.Message {
	b := []byte(payload)
	var msg *driver.Message
	if err := json.Unmarshal(b, &msg); err != nil {
		panic(err)
	}
	return msg
}
