package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ngductoann/go-telegram-bot/internal/domain/entity"
	"github.com/ngductoann/go-telegram-bot/internal/domain/repository"
	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
	"github.com/redis/go-redis/v9"
)

// redisCacheRepository implements CacheRepository using Redis
type redisCacheRepository struct {
	client *redis.Client
}

// NewRedisCacheRepository creates a new Redis cache repository
func NewRedisCacheRepository(client *redis.Client) repository.CacheRepository {
	return &redisCacheRepository{
		client: client,
	}
}

// Set stores a value in Redis with expiration
func (r *redisCacheRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value from Redis
func (r *redisCacheRepository) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Delete removes a key from Redis
func (r *redisCacheRepository) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// Exists checks if a key exists in Redis
func (r *redisCacheRepository) Exists(ctx context.Context, key string) (bool, error) {
	count, err := r.client.Exists(ctx, key).Result()
	return count > 0, err
}

// SetNX sets a key only if it doesn't exist
func (r *redisCacheRepository) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.client.SetNX(ctx, key, value, expiration).Result()
}

// Increment increments a numeric value
func (r *redisCacheRepository) Increment(ctx context.Context, key string) (int64, error) {
	return r.client.Incr(ctx, key).Result()
}

// Decrement decrements a numeric value
func (r *redisCacheRepository) Decrement(ctx context.Context, key string) (int64, error) {
	return r.client.Decr(ctx, key).Result()
}

// SetWithExpiration is an alias for Set
func (r *redisCacheRepository) SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Set(ctx, key, value, expiration)
}

// redisSessionCacheRepository implements SessionCacheRepository
type redisSessionCacheRepository struct {
	repository.CacheRepository
	client *redis.Client
}

// NewRedisSessionCacheRepository creates a new Redis session cache repository
func NewRedisSessionCacheRepository(client *redis.Client) repository.SessionCacheRepository {
	return &redisSessionCacheRepository{
		CacheRepository: NewRedisCacheRepository(client),
		client:          client,
	}
}

// GetUserSession retrieves a user session from cache
func (r *redisSessionCacheRepository) GetUserSession(ctx context.Context, userID uuid.UUID, chatID types.TelegramChatID) (*entity.UserSession, error) {
	key := fmt.Sprintf("session:%s:%d", userID.String(), int64(chatID))
	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // Session not found
		}
		return nil, err
	}

	var session entity.UserSession
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return &session, nil
}

// SetUserSession stores a user session in cache
func (r *redisSessionCacheRepository) SetUserSession(ctx context.Context, session *entity.UserSession, expiration time.Duration) error {
	key := fmt.Sprintf("session:%s:%d", session.UserID.String(), session.ChatID)

	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	return r.client.Set(ctx, key, data, expiration).Err()
}

// DeleteUserSession removes a user session from cache
func (r *redisSessionCacheRepository) DeleteUserSession(ctx context.Context, userID uuid.UUID, chatID types.TelegramChatID) error {
	key := fmt.Sprintf("session:%s:%d", userID.String(), int64(chatID))
	return r.client.Del(ctx, key).Err()
}

// GetActiveSessionCount returns the number of active sessions for a user
func (r *redisSessionCacheRepository) GetActiveSessionCount(ctx context.Context, userID uuid.UUID) (int64, error) {
	pattern := fmt.Sprintf("session:%s:*", userID.String())
	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		return 0, err
	}
	return int64(len(keys)), nil
}
