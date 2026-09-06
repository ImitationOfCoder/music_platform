package redis_client

import (
	"api_gateway_microservice/config"
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})

	cmd := client.Ping(context.Background())
	if cmd.Err() != nil {
		client.Close()
		fmt.Println("Failed to ping redis:", cmd.Err())
		os.Exit(1)
	}

	return client
}
