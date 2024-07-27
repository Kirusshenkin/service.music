package database

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"service.music/internal/config"
)

var ctx = context.Background()

func InitRedisClient(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}
