package repository

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
)

type MessageRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Message, error)
	GetByTelegramID(ctx context.Context, telegramMessageID int64) (*entity.Message, error)
	GetByTelegramUserID(ctx context.Context, telegramUserID types.TelegramUserID) ([]*entity.Message, error)
	GetByTelegramChatID(ctx context.Context, telegramChatID types.TelegramChatID) ([]*entity.Message, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.Message, error)
	GetByChatID(ctx context.Context, chatID uuid.UUID) ([]*entity.Message, error)
	Create(ctx context.Context, message *entity.Message) error
	Update(ctx context.Context, message *entity.Message) error
	Delete(ctx context.Context, telegramMessageID int64) error
}
