package repository

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
)

type MessageRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Message, error)
	FindByTelegramUserID(ctx context.Context, telegramUserID types.TelegramUserID) ([]*entity.Message, error)
	FindByTelegramChatID(ctx context.Context, telegramChatID types.TelegramChatID) ([]*entity.Message, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.Message, error)
	FindByChatID(ctx context.Context, chatID uuid.UUID) ([]*entity.Message, error)
	Create(ctx context.Context, message *entity.Message) error
	Update(ctx context.Context, message *entity.Message) error
	Delete(ctx context.Context, message *entity.Message) error
}
