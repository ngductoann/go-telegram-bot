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

type messageRepository struct {
	db *gorm.DB
}

// NewMessageRepository creates a new instance of messageRepository.
func NewMessageRepository(db *gorm.DB) repository.MessageRepository {
	return &messageRepository{db: db}
}

// GetByID retrieves a message by its UUID.
func (mr *messageRepository) GetByID(ctx context.Context, id uuid.UUID) (
	*entity.Message, error,
) {
	var message entity.Message
	err := mr.db.WithContext(ctx).
		Where("id = ?", id).
		First(&message).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrMessageNotFound
		}
		return nil, err
	}

	return &message, err
}

// GetByChatID retrieves messages by chat ID (Telegram Chat ID) with a limit.
func (mr *messageRepository) GetByChatID(ctx context.Context, chatID types.TelegramChatID, limit int) (
	[]*entity.Message, error,
) {
	err := mr.db.WithContext(ctx).
		Where("chat_id = ? and is_deleted = ?", chatID, false).
		Order("created_at desc").
		Limit(limit).
		Find(&[]*entity.Message{}).Error
	if err != nil {
		return nil, err
	}

	return []*entity.Message{}, err
}

// Create inserts a new message into the database.
func (mr *messageRepository) Create(ctx context.Context, message *entity.Message) error {
	return mr.db.WithContext(ctx).Create(message).Error
}

// Update modifies an existing message in the database.
func (mr *messageRepository) Update(ctx context.Context, message *entity.Message) error {
	return mr.db.WithContext(ctx).Save(message).Error
}

// Delete removes a message from the database by its UUID.
func (mr *messageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return mr.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Message{}).Error
}

// Search looks for messages containing the query string in their content, limited by the specified number.
func (mr *messageRepository) Search(ctx context.Context, query string, limit int) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := mr.db.WithContext(ctx).
		Where("content ILIKE ? AND is_deleted = ?", "%"+query+"%", false).
		Limit(limit).
		Find(&messages).Error
	if err != nil {
		return nil, err
	}

	return messages, err
}
