package connection

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "monpass",
		DB:       0,
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}
