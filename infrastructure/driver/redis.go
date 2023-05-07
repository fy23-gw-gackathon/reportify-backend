package driver

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
)

const JobQueueKey = "job_queue"

func NewRedisClient(addr string) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	log.Println("[REDIS] run a connection test with redis")
	log.Println("[REDIS] ping ...")
	if err := cli.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("[Redis] cannot connect to redis / ", err)
	}
	log.Println("[REDIS] ... pong connection ok!")
	return cli
}
