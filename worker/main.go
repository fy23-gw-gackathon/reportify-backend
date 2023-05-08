package main

import (
	"github.com/fy23-gw-gackathon/reportify-backend/config"
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/driver"
	"golang.org/x/net/context"
)

func main() {
	cfg := config.Load()
	rdb := driver.NewRedisClient(cfg.Datastore.Address)
	client := driver.NewHttp()
	gpt := driver.NewGptDriver(cfg)
	h := newHandler(client, gpt)
	s := newServer(rdb)
	s.Run(func(ctx context.Context, payload string) error {
		return h.CommentReport(payload)
	})
}
