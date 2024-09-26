package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/aparnasukesh/movies-booking-svc/config"
	"github.com/go-redis/redis/v8"
)

func NewRedis(cfg config.Config) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // No password
		DB:       0,  // Default DB
	})

	// Test Redis connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return redisClient, nil
}
