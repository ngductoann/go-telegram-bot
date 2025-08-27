package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngductoann/go-telegram-bot/internal/domain/entity"
	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
)

// MessageRepository defines the interface for message repository
type MessageRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Message, error)
	GetByChatID(ctx context.Context, chatID types.TelegramChatID, limit int) ([]*entity.Message, error)
	Create(ctx context.Context, message *entity.Message) error
	Update(ctx context.Context, message *entity.Message) error
	Delete(ctx context.Context, id uuid.UUID) error
	Search(ctx context.Context, query string, limit int) ([]*entity.Message, error)
}
