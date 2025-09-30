package repository

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/errors"
	"go-telegram-bot/internal/domain/repository"
	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

// NewChatRepository creates a new instance of ChatRepository.
func NewChatRepository(db *gorm.DB) repository.ChatRepository {
	return &chatRepository{db: db}
}

// GetByUUID retrieves a chat by its UUID.
func (r *chatRepository) GetByUUID(
	ctx context.Context, id uuid.UUID,
) (*entity.Chat, error) {
	var chat entity.Chat
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&chat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrChatNotFound
		}
		return nil, err
	}

	return &chat, nil
}

// GetByTelegramChatID retrieves a chat by its TelegramChatID.
func (r *chatRepository) GetByTelegramChatID(
	ctx context.Context, telegramChatID types.TelegramChatID,
) (*entity.Chat, error) {
	var chat entity.Chat
	if err := r.db.WithContext(ctx).
		Where("telegram_chat_id = ?", telegramChatID).
		First(&chat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrChatNotFound
		}
		return nil, err
	}

	return &chat, nil
}

// GetAll retrieves all chats from the database.
func (r *chatRepository) GetAll(
	ctx context.Context,
) ([]*entity.Chat, error) {
	var chats []*entity.Chat

	if err := r.db.WithContext(ctx).Find(&chats).Error; err != nil {
		return nil, err
	}

	return chats, nil
}

// Create inserts a new chat into the database.
func (r *chatRepository) Create(
	ctx context.Context, chat *entity.Chat,
) error {
	return r.db.WithContext(ctx).Create(chat).Error
}

// Update modifies an existing chat in the database.
func (r *chatRepository) Update(
	ctx context.Context, chat *entity.Chat,
) error {
	return r.db.WithContext(ctx).Save(chat).Error
}

// Update modifies an existing chat in the database.
func (r *chatRepository) Delete(
	ctx context.Context, telegramChatID types.TelegramChatID,
) error {
	return r.db.WithContext(ctx).
		Where("telegram_chat_id = ?", telegramChatID).
		Delete(&entity.Chat{}).Error
}

// GetActiveChats retrieves a list of active chats up to the specified limit.
func (r *chatRepository) GetActiveChats(
	ctx context.Context, limit int,
) ([]*entity.Chat, error) {
	var chats []*entity.Chat
	if err := r.db.WithContext(ctx).
		Where("is_active = ?", true).
		Limit(limit).
		Find(&chats).Error; err != nil {
		return nil, err
	}

	return chats, nil
}
