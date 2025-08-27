package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ngductoann/go-telegram-bot/internal/domain/entity"
	"github.com/ngductoann/go-telegram-bot/internal/domain/errors"
	"github.com/ngductoann/go-telegram-bot/internal/domain/repository"
	"gorm.io/gorm"
)

// UserProfileRepository is the interface for user profile repository
// It defines methods to interact with user profiles in the database
type UserProfileRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserProfile, error)
	Create(ctx context.Context, profile *entity.UserProfile) error
	Update(ctx context.Context, profile *entity.UserProfile) error
	Delete(ctx context.Context, userID uuid.UUID) error
	GetOrCreate(ctx context.Context, userID uuid.UUID) (*entity.UserProfile, error)
}

// userProfileRepository implements UserProfileRepository interface
type userProfileRepository struct {
	db *gorm.DB
}

// NewUserProfileRepository creates a new instance of UserProfileRepository
func NewUserProfileRepository(db *gorm.DB) repository.UserProfileRepository {
	return &userProfileRepository{
		db: db,
	}
}

// GetByUserID retrieves a user profile by user ID
func (upr *userProfileRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserProfile, error) {
	var profile entity.UserProfile
	err := upr.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&profile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}

	return &profile, nil
}

// Create creates a new user profile
func (upr *userProfileRepository) Create(ctx context.Context, profile *entity.UserProfile) error {
	now := time.Now()
	profile.CreatedAt = now
	profile.UpdatedAt = now
	return upr.db.WithContext(ctx).Create(profile).Error
}

// Update updates an existing user profile
func (upr *userProfileRepository) Update(ctx context.Context, profile *entity.UserProfile) error {
	profile.UpdatedAt = time.Now()
	return upr.db.WithContext(ctx).Save(profile).Error
}

// Delete deletes a user profile by user ID
func (upr *userProfileRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	return upr.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&entity.UserProfile{}).Error
}

// GetOrCreate retrieves a user profile by user ID or creates a new one if it doesn't exist
func (upr *userProfileRepository) GetOrCreate(ctx context.Context, userID uuid.UUID) (*entity.UserProfile, error) {
	// Try to get existing profile
	profile, err := upr.GetByUserID(ctx, userID)
	if err == nil {
		return profile, nil
	}

	// If not found, create new profile
	if err == errors.ErrUserNotFound {
		newProfile := entity.NewUserProfileWithUserID(userID)

		if err := upr.Create(ctx, newProfile); err != nil {
			return nil, err
		}

		return newProfile, nil
	}

	return nil, err
}
