package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngductoann/go-telegram-bot/internal/domain/entity"
	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
)

type ChatRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Chat, error)
	GetByTelegramChatID(ctx context.Context, chatID types.TelegramChatID) (*entity.Chat, error)
	Create(ctx context.Context, chat *entity.Chat) error
	Update(ctx context.Context, chat *entity.Chat) error
	DeleteByTelegramChatID(ctx context.Context, chatID types.TelegramChatID) error
	DeleteByID(ctx context.Context, uuid uuid.UUID) error
	GetActiveChats(ctx context.Context, limit int) ([]*entity.Chat, error)
}
