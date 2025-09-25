package repository

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
)

type UserRepository interface {
	// Define methods for user repository
	FindByTelegramUserID(ctx context.Context, telegramUserID types.TelegramUserID) (*entity.User, error)
	FindByUUID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	// Add other necessary methods like Create, Update, Delete, etc.
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, user *entity.User) error
	GetOrCreate(ctx context.Context, telegramUserID types.TelegramUserID, firstName string) (*entity.User, error)
}
