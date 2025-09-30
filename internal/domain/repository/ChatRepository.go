package repository

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
)

type ChatRepository interface {
	GetByUUID(ctx context.Context, id uuid.UUID) (*entity.Chat, error)
	GetByTelegramChatID(ctx context.Context, chatID types.TelegramChatID) (*entity.Chat, error)
	// Add other nested repository methods here
	Create(ctx context.Context, chat *entity.Chat) error
	Update(ctx context.Context, chat *entity.Chat) error
	Delete(ctx context.Context, telegramChatID types.TelegramChatID) error
	GetActiveChats(ctx context.Context, limit int) ([]*entity.Chat, error)
}
