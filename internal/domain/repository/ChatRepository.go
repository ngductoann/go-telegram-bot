package repository

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
)

type ChatRepository interface {
	FindByUUID(ctx context.Context, id uuid.UUID) (*entity.Chat, error)
	FindByTelegramChatID(ctx context.Context, chatID types.TelegramChatID) (*entity.Chat, error)
	// Add other nested repository methods here
	Create(ctx context.Context, chat *entity.Chat) error
	Update(ctx context.Context, chat *entity.Chat) error
	Delete(ctx context.Context, chat *entity.Chat) error
	GetOrCreate(ctx context.Context, chatID types.TelegramChatID, chatTypes types.ChatType) (*entity.Chat, error)
	GetActiveChats(ctx context.Context, limit int) ([]*entity.Chat, error)
}
