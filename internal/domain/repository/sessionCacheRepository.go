package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ngductoann/go-telegram-bot/internal/domain/entity"
	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
)

type SessionCacheRepository interface {
	CacheRepository
	GetUserSession(ctx context.Context, userID uuid.UUID, chatID types.TelegramChatID) (*entity.UserSession, error)
	SetUserSession(ctx context.Context, session *entity.UserSession, expiration time.Duration) error
	DeleteUserSession(ctx context.Context, userID uuid.UUID, chatID types.TelegramChatID) error
	GetActiveSessionCount(ctx context.Context, userID uuid.UUID) (int64, error)
}
