package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ngductoann/go-telegram-bot/internal/domain/entity"
	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
)

// SessionRepository defines the interface for session repository
type SessionRepository interface {
	GetByUserAndChat(ctx context.Context, userID uuid.UUID, chatID types.TelegramChatID) (*entity.UserSession, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.UserSession, error)
	Create(ctx context.Context, session *entity.UserSession) error
	Update(ctx context.Context, session *entity.UserSession) error
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteExpired(ctx context.Context) error
	ExtendSession(ctx context.Context, id uuid.UUID, duration time.Duration) error
}
