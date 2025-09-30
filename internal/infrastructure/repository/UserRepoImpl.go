package repository

import (
	"context"
	"time"

	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/errors"
	"go-telegram-bot/internal/domain/repository"
	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

// GetByTelegramUserID retrieves a user by their Telegram user ID
func (r *userRepository) GetByTelegramUserID(
	ctx context.Context, telegramUserID types.TelegramUserID,
) (*entity.User, error) {
	var user entity.User

	if err := r.db.WithContext(ctx).Where("telegram_user_id = ?", telegramUserID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// GetByUUID retrieves a user by their UUID
func (r *userRepository) GetByUUID(
	ctx context.Context, uuid uuid.UUID,
) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("id = ?", uuid).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// GetAll retrieves all users from the database
func (r *userRepository) GetAll(
	ctx context.Context,
) ([]*entity.User, error) {
	var users []*entity.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Create adds a new user to the database
func (r *userRepository) Create(
	ctx context.Context, user *entity.User,
) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// Update modifies an existing user in the database
func (r *userRepository) Update(
	ctx context.Context, user *entity.User,
) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete removes a user from the database
func (r *userRepository) Delete(
	ctx context.Context, telegramUserID types.TelegramUserID,
) error {
	return r.db.WithContext(ctx).
		Where("telegram_user_id = ?", telegramUserID).
		Delete(&entity.User{}).Error
}

// UpdareLastSeen updates the LastSeenAt field of a user
func (r *userRepository) UpdareLastSeen(
	ctx context.Context,
	telegramUserID types.TelegramUserID,
	timestamp time.Time,
) error {
	return r.db.WithContext(ctx).Model(&entity.User{}).
		Where("telegram_user_id = ?", telegramUserID).
		Updates(map[string]any{
			"last_seen_at": timestamp,
		}).Error
}

// DeactivateUser sets the IsActive field of a user to false
func (r *userRepository) DeactivateUser(
	ctx context.Context, telegramUserID types.TelegramUserID,
) error {
	return r.db.WithContext(ctx).Model(&entity.User{}).
		Where("telegram_user_id = ?", telegramUserID).
		Updates(map[string]any{
			"is_active": false,
		}).Error
}

func (r *userRepository) ActivateUser(
	ctx context.Context, telegramUserID types.TelegramUserID,
) error {
	return r.db.WithContext(ctx).Model(&entity.User{}).
		Where("telegram_user_id = ?", telegramUserID).
		Updates(map[string]any{
			"is_active": true,
		}).Error
}
