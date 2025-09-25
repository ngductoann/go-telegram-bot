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

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

// FindByTelegramUserID retrieves a user by their Telegram user ID
func (r *userRepository) FindByTelegramUserID(
	ctx context.Context,
	telegramUserID types.TelegramUserID,
) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Where("telegram_user_id = ?", telegramUserID).
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// FindByUUID retrieves a user by their UUID
func (r *userRepository) FindByUUID(ctx context.Context, uuid uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Where("id = ?", uuid).
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// FindAll retrieves all users from the database
func (r *userRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Create adds a new user to the database
func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	existsUser, err := r.FindByTelegramUserID(ctx, user.TelegramUserID)
	if existsUser != nil || err == nil {
		return errors.ErrUserAlreadyExists
	}
	return r.db.WithContext(ctx).Create(user).Error
}

// Update modifies an existing user in the database
func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	_, err := r.FindByUUID(ctx, user.ID)
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete removes a user from the database
func (r *userRepository) Delete(ctx context.Context, user *entity.User) error {
	_, err := r.FindByUUID(ctx, user.ID)
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Delete(user).Error
}

// GetOrCreate retrieves a user by their Telegram user ID or creates a new one if not found
func (r *userRepository) GetOrCreate(
	ctx context.Context,
	telegramUserID types.TelegramUserID,
	firstName string,
) (*entity.User, error) {
	user, err := r.FindByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		user := entity.NewUser(telegramUserID, firstName)
		if err := r.Create(ctx, user); err != nil {
			return nil, err
		}
		return user, nil
	}

	return user, nil
}
