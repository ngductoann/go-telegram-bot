package repository

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/errors"
	"go-telegram-bot/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// userProfileRepoImpl is the implementation of UserProfileRepository
type userProfileRepoImpl struct {
	db       *gorm.DB
	userRepo repository.UserRepository
}

// NewUserProfileRepo creates a new instance of UserProfileRepository
func NewUserProfileRepo(db *gorm.DB, userRepo repository.UserRepository) repository.UserProfileRepository {
	return &userProfileRepoImpl{db: db, userRepo: userRepo}
}

// FindByID retrieves a user profile by its ID
func (r *userProfileRepoImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.UserProfile, error) {
	var profile entity.UserProfile
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&profile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrProfileNotFound
		}
		return nil, err
	}

	return &profile, nil
}

// FindByUserID retrieves a user profile by the associated user ID
func (r *userProfileRepoImpl) FindByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*entity.UserProfile, error) {
	_, err := r.userRepo.FindByUUID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var profile entity.UserProfile
	err = r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&profile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrProfileNotFound
		}
		return nil, err
	}

	return &profile, nil
}

// Update updates an existing user profile
func (r *userProfileRepoImpl) Update(ctx context.Context, profile *entity.UserProfile) error {
	var profileInDB entity.UserProfile

	err := r.db.WithContext(ctx).
		Where("id = ?", profile.ID).
		First(&profileInDB).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrProfileNotFound
		}
		return err
	}

	return r.db.WithContext(ctx).Save(profile).Error
}

// Create creates a new user profile
func (r *userProfileRepoImpl) Create(ctx context.Context, profile *entity.UserProfile) error {
	// Check profile constraints, e.g., unique user_id
	err := r.db.WithContext(ctx).
		Where("user_id = ?", profile.UserID).First(&entity.UserProfile{}).Error
	if err == nil {
		return errors.ErrProfileIsExist
	}

	return r.db.WithContext(ctx).Create(profile).Error
}

// Delete deletes a user profile
func (r *userProfileRepoImpl) Delete(ctx context.Context, profile *entity.UserProfile) error {
	_, error := r.FindByID(ctx, profile.ID)
	if error != nil {
		return error
	}

	return r.db.WithContext(ctx).Delete(profile).Error
}
