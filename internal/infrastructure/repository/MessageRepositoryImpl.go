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

type messageRepository struct {
	db       *gorm.DB
	userRepo repository.UserRepository
	chatRepo repository.ChatRepository
}

// NewMessageRepository creates a new instance of MessageRepository
func NewMessageRepository(
	db *gorm.DB,
	userRepo repository.UserRepository,
	chatRepo repository.ChatRepository,
) repository.MessageRepository {
	return &messageRepository{
		db:       db,
		userRepo: userRepo,
		chatRepo: chatRepo,
	}
}

// Create implements repository.MessageRepository.
func (m *messageRepository) Create(
	ctx context.Context, message *entity.Message,
) error {
	return m.db.WithContext(ctx).Create(message).Error
}

// Delete implements repository.MessageRepository.
func (m *messageRepository) Delete(
	ctx context.Context, telegramMessageID int64,
) error {
	return m.db.WithContext(ctx).
		Where("telegram_message_id = ? AND is_deleted = ?", telegramMessageID, false).
		Delete(&entity.Message{}).Error
}

// GetByChatID implements repository.MessageRepository.
func (m *messageRepository) GetByChatID(
	ctx context.Context, chatID uuid.UUID,
) ([]*entity.Message, error) {
	var messages []*entity.Message
	if err := m.db.WithContext(ctx).
		Where("chat_id = ? AND is_deleted = ?", chatID, false).
		Order("created_at DESC").
		Find(&messages).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrMessageNotFound
		}
	}
	return messages, nil
}

// GetByUserID implements repository.MessageRepository.
func (m *messageRepository) GetByUserID(
	ctx context.Context, userID uuid.UUID,
) ([]*entity.Message, error) {
	var messages []*entity.Message
	if err := m.db.WithContext(ctx).
		Where("user_id = ? AND is_deleted = ?", userID, false).
		Order("created_at DESC").
		Find(&messages).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrMessageNotFound
		}
	}

	return messages, nil
}

// GetByID implements repository.MessageRepository.
func (m *messageRepository) GetByID(
	ctx context.Context, id uuid.UUID,
) (*entity.Message, error) {
	var message entity.Message

	if err := m.db.WithContext(ctx).
		Where("id = ? AND is_deleted = ?", id, false).
		First(&message).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrMessageNotFound
		}
		return nil, err
	}

	return &message, nil
}

// GetByTelegramID implements repository.MessageRepository.
func (m *messageRepository) GetByTelegramID(
	ctx context.Context, telegramMessageID int64,
) (*entity.Message, error) {
	var message entity.Message

	if err := m.db.WithContext(ctx).
		Where("telegram_message_id = ? AND is_deleted = ?", telegramMessageID, false).
		First(&message).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrMessageNotFound
		}
		return nil, err
	}

	return &message, nil
}

// GetByTelegramChatID implements repository.MessageRepository.
func (m *messageRepository) GetByTelegramChatID(
	ctx context.Context, telegramChatID types.TelegramChatID,
) ([]*entity.Message, error) {
	chat, err := m.chatRepo.GetByTelegramChatID(ctx, telegramChatID)
	if err != nil {
		return nil, err
	}

	var messages []*entity.Message
	if err := m.db.WithContext(ctx).
		Where("chat_id = ? AND is_deleted = ?", chat.ID, false).
		Order("created_at DESC").
		Find(&messages).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrMessageNotFound
		}
		return nil, err
	}

	return messages, nil
}

// GetByTelegramUserID implements repository.MessageRepository.
func (m *messageRepository) GetByTelegramUserID(
	ctx context.Context, telegramUserID types.TelegramUserID,
) ([]*entity.Message, error) {
	var messages []*entity.Message
	if err := m.db.WithContext(ctx).
		Where("user_id = ? AND is_deleted = ?", telegramUserID, false).
		Order("created_at DESC").
		Find(&messages).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrMessageNotFound
		}
		return nil, err
	}
	return messages, nil
}

// Update implements repository.MessageRepository.
func (m *messageRepository) Update(
	ctx context.Context, message *entity.Message,
) error {
	return m.db.WithContext(ctx).Save(message).Error
}
