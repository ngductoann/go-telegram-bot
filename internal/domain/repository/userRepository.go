package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ngductoann/go-telegram-bot/internal/domain/entity"
	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
)

// UserRepository defines the interface for user repository
type UserRepository interface {
	GetByTelegramUserID(ctx context.Context, userID types.TelegramUserID) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, userID types.TelegramUserID) error
	UpdateLastSeen(ctx context.Context, userID types.TelegramUserID, timestamp time.Time) error
	GetActiveUsers(ctx context.Context, limit int) ([]*entity.User, error)
}
