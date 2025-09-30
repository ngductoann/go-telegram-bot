package repository

import (
	"context"
	"time"

	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
)

type UserRepository interface {
	// Define methods for user repository
	GetByTelegramUserID(ctx context.Context, telegramUserID types.TelegramUserID) (*entity.User, error)
	GetByUUID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
	// Add other necessary methods like Create, Update, Delete, etc.
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, telegramUserID types.TelegramUserID) error
	UpdareLastSeen(ctx context.Context, telegramUserID types.TelegramUserID, timestamp time.Time) error
	DeactivateUser(ctx context.Context, telegramUserID types.TelegramUserID) error
	ActivateUser(ctx context.Context, telegramUserID types.TelegramUserID) error
}
