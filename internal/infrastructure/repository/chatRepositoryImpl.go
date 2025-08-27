package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngductoann/go-telegram-bot/internal/domain/entity"
	"github.com/ngductoann/go-telegram-bot/internal/domain/errors"
	"github.com/ngductoann/go-telegram-bot/internal/domain/repository"
	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
	"gorm.io/gorm"
)

// chatRepository implements the ChatRepository interface using GORM for database operations.
type chatRepository struct {
	db *gorm.DB
}

// NewChatRepository creates a new instance of chatRepository.
func NewChatRepository(db *gorm.DB) repository.ChatRepository {
	return &chatRepository{db: db}
}

// GetByID retrieves a chat by its Telegram chat ID.
func (cr *chatRepository) GetByTelegramChatID(ctx context.Context, telegramChatID types.TelegramChatID) (
	*entity.Chat, error,
) {
	var chat entity.Chat
	err := cr.db.WithContext(ctx).
		Where("telegram_chat_id = ?", telegramChatID).
		First(&chat).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrChatNotFound
		}
		return nil, err
	}

	return &chat, err
}

// GetByTelegramChatID retrieves a chat by its UUID.
func (cr *chatRepository) GetByID(ctx context.Context, chatID uuid.UUID) (
	*entity.Chat, error,
) {
	var chat entity.Chat
	err := cr.db.WithContext(ctx).
		Where("id = ?", chatID).
		First(&chat).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrChatNotFound
		}
		return nil, err
	}

	return &chat, err
}

// Create adds a new chat to the database.
func (cr *chatRepository) Create(ctx context.Context, chat *entity.Chat) error {
	return cr.db.WithContext(ctx).Create(chat).Error
}

// Update modifies an existing chat in the database.
func (cr *chatRepository) Update(ctx context.Context, chat *entity.Chat) error {
	return cr.db.WithContext(ctx).Save(chat).Error
}

// DeleteByTelegramChatID removes a chat from the database using its Telegram chat ID.
func (cr *chatRepository) DeleteByTelegramChatID(ctx context.Context, telegramChatID types.TelegramChatID) error {
	return cr.db.WithContext(ctx).Where("telegram_chat_id = ?", telegramChatID).Delete(&entity.Chat{}).Error
}

// DeleteByID removes a chat from the database using its UUID.
func (cr *chatRepository) DeleteByID(ctx context.Context, chatID uuid.UUID) error {
	return cr.db.WithContext(ctx).Where("id = ?", chatID).Delete(&entity.Chat{}).Error
}

// GetActiveChats retrieves a list of active chats, limited by the specified number.
func (cr *chatRepository) GetActiveChats(ctx context.Context, limit int) ([]*entity.Chat, error) {
	var chats []*entity.Chat

	err := cr.db.WithContext(ctx).
		Where("is_active = ?", true).
		Limit(limit).
		Find(&chats).Error
	if err != nil {
		return nil, err
	}

	return chats, nil
}
