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

// FindByUUID retrieves a chat by its UUID.
func (r *chatRepository) FindByUUID(ctx context.Context, id uuid.UUID) (*entity.Chat, error) {
	var chat entity.Chat
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&chat).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrChatNotFound
		}
		return nil, err
	}

	return &chat, err
}

// FindByTelegramChatID retrieves a chat by its TelegramChatID.
func (r *chatRepository) FindByTelegramChatID(
	ctx context.Context,
	telegramChatID types.TelegramChatID,
) (*entity.Chat, error) {
	var chat entity.Chat
	err := r.db.WithContext(ctx).
		Where("telegram_chat_id = ?", telegramChatID).
		First(&chat).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrChatNotFound
		}
		return nil, err
	}

	return &chat, nil
}

// FindAll retrieves all chats from the database.
func (r *chatRepository) FindAll(ctx context.Context) ([]*entity.Chat, error) {
	var chats []*entity.Chat

	err := r.db.WithContext(ctx).Find(&chats).Error
	if err != nil {
		return nil, err
	}

	return chats, nil
}

// Create inserts a new chat into the database.
func (r *chatRepository) Create(ctx context.Context, chat *entity.Chat) error {
	existsChat, err := r.FindByTelegramChatID(ctx, chat.TelegramChatID)
	if existsChat != nil || err == nil {
		return errors.ErrChatAlreadyExists
	}
	return r.db.WithContext(ctx).Create(chat).Error
}

// Update modifies an existing chat in the database.
func (r *chatRepository) Update(ctx context.Context, chat *entity.Chat) error {
	_, err := r.FindByTelegramChatID(ctx, chat.TelegramChatID)
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Save(chat).Error
}

// Update modifies an existing chat in the database.
func (r *chatRepository) Delete(ctx context.Context, chat *entity.Chat) error {
	_, err := r.FindByTelegramChatID(ctx, chat.TelegramChatID)
	if err != nil {
		return nil
	}
	return r.db.WithContext(ctx).Delete(chat).Error
}

// GetOrCreate retrieves a chat by its TelegramChatID or creates it if it doesn't exist.
func (r *chatRepository) GetOrCreate(
	ctx context.Context,
	chatID types.TelegramChatID,
	chatType types.ChatType,
) (*entity.Chat, error) {
	chat, err := r.FindByTelegramChatID(ctx, chatID)
	if err != nil {
		chat = entity.NewChat(chatID, chatType)
		err := r.Create(ctx, chat)
		if err != nil {
			return nil, err
		}
		return r.FindByTelegramChatID(ctx, chatID)
	}

	return chat, nil
}

// GetActiveChats retrieves a list of active chats up to the specified limit.
func (r *chatRepository) GetActiveChats(ctx context.Context, limit int) ([]*entity.Chat, error) {
	var chats []*entity.Chat
	err := r.db.WithContext(ctx).
		Where("is_active = ?", true).
		Limit(limit).
		Find(&chats).Error
	if err != nil {
		return nil, err
	}

	return chats, nil
}
