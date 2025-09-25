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
func (m *messageRepository) Create(ctx context.Context, message *entity.Message) error {
	return m.db.WithContext(ctx).Create(message).Error
}

// Delete implements repository.MessageRepository.
func (m *messageRepository) Delete(ctx context.Context, message *entity.Message) error {
	return m.db.WithContext(ctx).Delete(message).Error
}

// FindByChatID implements repository.MessageRepository.
func (m *messageRepository) FindByChatID(ctx context.Context, chatID uuid.UUID) ([]*entity.Message, error) {
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

// FindByUserID implements repository.MessageRepository.
func (m *messageRepository) FindByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.Message, error) {
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

// FindByID implements repository.MessageRepository.
func (m *messageRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Message, error) {
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

// FindByTelegramChatID implements repository.MessageRepository.
func (m *messageRepository) FindByTelegramChatID(ctx context.Context, telegramChatID types.TelegramChatID) ([]*entity.Message, error) {
	chat, err := m.chatRepo.FindByTelegramChatID(ctx, telegramChatID)
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

// FindByTelegramUserID implements repository.MessageRepository.
func (m *messageRepository) FindByTelegramUserID(ctx context.Context, telegramUserID types.TelegramUserID) ([]*entity.Message, error) {
	user, err := m.userRepo.FindByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, err
	}

	var messages []*entity.Message
	if err := m.db.WithContext(ctx).
		Where("user_id = ? AND is_deleted = ?", user.ID, false).
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
func (m *messageRepository) Update(ctx context.Context, message *entity.Message) error {
	if _, err := m.FindByID(ctx, message.ID); err != nil {
		return err
	}

	return m.db.WithContext(ctx).Save(message).Error
}
