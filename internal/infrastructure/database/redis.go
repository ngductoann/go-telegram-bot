package database

import (
	"context"
	"fmt"

	domainService "github.com/ngductoann/go-telegram-bot/internal/domain/service"
	"github.com/redis/go-redis/v9"
)

// NewRedisConnection creates a new Redis client based on the provided configuration.
func NewRedisConnection(redisURL, host, port, password string, db int, logger domainService.Logger) (*redis.Client, error) {
	var client *redis.Client

	if redisURL != "" && redisURL != ":" {
		opts, err := redis.ParseURL(redisURL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
		}
		client = redis.NewClient(opts)
	} else {
		addr := ""
		if host == "" {
			addr = "localhost:6379"
		} else {
			addr = host + ":" + port
		}
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})
	}

	if err := RedisHealthCheck(client); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return client, nil
}

// RedisHealthCheck checks the health of the Redis connection by sending a PING command.
func RedisHealthCheck(client *redis.Client) error {
	ctx := context.Background()
	return client.Ping(ctx).Err()
}
