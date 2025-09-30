package usecase

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/errors"
	"go-telegram-bot/internal/domain/repository"

	"github.com/google/uuid"
)

// UserProfileUseCase handles user profile related business logic
type UserProfileUseCase struct {
	userProfileRepo repository.UserProfileRepository
	userRepo        repository.UserRepository
}

// NewUserProfileUseCase creates a new instance of UserProfileUseCase
func NewUserProfileUseCase(
	userProfileRepo repository.UserProfileRepository,
	userRepo repository.UserRepository,
) *UserProfileUseCase {
	return &UserProfileUseCase{
		userProfileRepo: userProfileRepo,
		userRepo:        userRepo,
	}
}

// GetByUserID retrieves a user profile by user ID
func (u *UserProfileUseCase) GetByUserID(
	ctx context.Context, userID uuid.UUID,
) (*entity.UserProfile, error) {
	if userID == uuid.Nil {
		return nil, errors.ErrInvalidInput
	}

	// Check if user exists
	_, err := u.userRepo.GetByUUID(ctx, userID)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	userProfile, err := u.userProfileRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}

// GetByID retrieves a user profile by its ID
func (u *UserProfileUseCase) GetByID(
	ctx context.Context, id uuid.UUID,
) (*entity.UserProfile, error) {
	if id == uuid.Nil {
		return nil, errors.ErrInvalidInput
	}

	userProfile, err := u.userProfileRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return userProfile, nil
}

// Create creates a new user profile
func (u *UserProfileUseCase) Create(
	ctx context.Context, profile *entity.UserProfile,
) error {
	_, err := u.GetByUserID(ctx, profile.UserID)
	if err == nil {
		return errors.ErrProfileIsExist
	}

	return u.userProfileRepo.Create(ctx, profile)
}

// Update updates an existing user profile
func (u *UserProfileUseCase) Update(
	ctx context.Context, profile *entity.UserProfile,
) error {
	userProfile, err := u.GetByID(ctx, profile.ID)
	if err != nil {
		return err
	}
	userProfile.DateActive = profile.DateActive

	return u.userProfileRepo.Update(ctx, userProfile)
}

// Delete deletes a user profile
func (u *UserProfileUseCase) Delete(
	ctx context.Context, profile *entity.UserProfile,
) error {
	_, err := u.GetByID(ctx, profile.ID)
	if err != nil {
		return nil
	}

	return u.userProfileRepo.Delete(ctx, profile)
}
