package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ngductoann/go-telegram-bot/internal/domain/entity"
	"github.com/ngductoann/go-telegram-bot/internal/domain/errors"
	"github.com/ngductoann/go-telegram-bot/internal/domain/repository"
	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
	"gorm.io/gorm"
)

// userRepository implements the userRepository interface
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

// GetByTelegramUserID
func (ur *userRepository) GetByTelegramUserID(ctx context.Context, userID types.TelegramUserID) (*entity.User, error) {
	var user entity.User
	err := ur.db.WithContext(ctx).
		Where("telegram_user_id = ?", userID).
		First(&user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// GetByID retrieves a user by their UUID
func (ur *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := ur.db.WithContext(ctx).
		Where("id", id).
		First(&user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// Create creates a new user
func (ur *userRepository) Create(ctx context.Context, user *entity.User) error {
	return ur.db.WithContext(ctx).Create(user).Error
}

// Update updates an existing user
func (ur *userRepository) Update(ctx context.Context, user *entity.User) error {
	return ur.db.WithContext(ctx).Save(user).Error
}

// Delete deletes a user by their Telegram user ID
func (ur *userRepository) Delete(ctx context.Context, userID types.TelegramUserID) error {
	return ur.db.WithContext(ctx).
		Where("telegram_user_id = ?", userID).
		Delete(&entity.User{}).Error
}

// UpdateLastSeen updates the last seen timestamp of a user
func (ur *userRepository) UpdateLastSeen(ctx context.Context, userID types.TelegramUserID, timestamp time.Time) error {
	return ur.db.WithContext(ctx).Model(&entity.User{}).
		Where("telegram_user_id = ?", userID).
		Updates(map[string]interface{}{
			"last_seen_at": timestamp,
			"updated_at":   time.Now(),
		}).Error
}

// GetActiveUsers retrieves a list of active users up to the specified limit
func (ur *userRepository) GetActiveUsers(ctx context.Context, limit int) ([]*entity.User, error) {
	var users []*entity.User
	err := ur.db.WithContext(ctx).
		Where("is_active = ?", true).
		Limit(limit).
		Find(&users).Error

	return users, err
}
