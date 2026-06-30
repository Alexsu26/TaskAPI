package cache

import (
	"context"
	"fmt"
	"time"

	"taskapi/internal/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(config *config.Config) (*redis.Client, error) {
	addr := config.Redis.Host + ":" + config.Redis.Port

	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, fmt.Errorf("ping redis error: %w", err)
	}
	return client, nil
}
