package usecase

import (
	"context"
	"time"

	"go-telegram-bot/internal/application/dto"
	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/errors"
	"go-telegram-bot/internal/domain/repository"
	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
)

// UserUseCase struct
type UserUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase function creates a new UserUserCase
func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

// GetByUUID function retrieves a user by their UUID
func (u *UserUseCase) GetByUUID(
	ctx context.Context,
	userID uuid.UUID,
) (*dto.UserResponse, error) {
	user, err := u.userRepo.GetByUUID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.ErrUserNotFound
	}
	return mapUserToDTO(user), nil
}

// GetAll function retrieves all users
func (u *UserUseCase) GetAll(
	ctx context.Context,
) ([]*dto.UserResponse, error) {
	users, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, errors.ErrUserNotFound
	}
	userDTOs := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		userDTOs[i] = mapUserToDTO(user)
	}
	return userDTOs, nil
}

// GetByTelegramUserID function retrieves a user by their Telegram user ID
func (u *UserUseCase) GetByTelegramUserID(
	ctx context.Context,
	telegramUserID types.TelegramUserID,
) (*dto.UserResponse, error) {
	user, err := u.userRepo.GetByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.ErrUserNotFound
	}

	return mapUserToDTO(user), nil
}

// CreateUser function creates a new user
func (u *UserUseCase) CreateUser(
	ctx context.Context,
	req *dto.CreateUserRequest,
) (*dto.UserResponse, error) {
	// Check if user already exists
	existingUser, err := u.userRepo.GetByTelegramUserID(ctx, req.TelegramID)
	if err != nil && existingUser == nil {
		return nil, errors.ErrUserAlreadyExists
	}

	// Create new user
	user := entity.NewUser(req.TelegramID, req.FirstName)
	user.LastName = req.LastName
	user.Username = req.Username
	user.IsBot = req.IsBot

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return mapUserToDTO(user), nil
}

// UpdateUser function updates an existing user's information
func (u *UserUseCase) UpdateUser(
	ctx context.Context,
	telegramUserID types.TelegramUserID,
	req *dto.UpdateUserRequest,
) (*dto.UserResponse, error) {
	user, err := u.userRepo.GetByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.ErrUserNotFound
	}
	user.Username = req.Username
	user.FirstName = req.FirstName
	user.LastName = req.LastName

	if err := u.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return mapUserToDTO(user), nil
}

// DeleteUser function deletes a user by their Telegram user ID
func (u *UserUseCase) DeleteUser(
	ctx context.Context, telegramUserID types.TelegramUserID,
) error {
	return u.userRepo.Delete(ctx, telegramUserID)
}

// UpdateLastSeen function updates the last seen timestamp of a user
func (u *UserUseCase) UpdateLastSeen(
	ctx context.Context, telegramUserID types.TelegramUserID,
) error {
	return u.userRepo.UpdareLastSeen(ctx, telegramUserID, time.Now())
}

// DeactivateUser function deactivates a user by their Telegram user ID
func (u *UserUseCase) DeactivateUser(
	ctx context.Context, telegramUserID types.TelegramUserID,
) error {
	return u.userRepo.DeactivateUser(ctx, telegramUserID)
}

// ActivateUser function activates a user by their Telegram user ID
func (u *UserUseCase) ActivateUser(
	ctx context.Context, telegramUserID types.TelegramUserID,
) error {
	return u.userRepo.ActivateUser(ctx, telegramUserID)
}

// GetOrCreateUser function retrieves an existing user or creates a new one if not found
func (u *UserUseCase) GetOrCreateUser(
	ctx context.Context, req *dto.CreateUserRequest,
) (*dto.UserResponse, error) {
	// Try to get existing user
	user, err := u.userRepo.GetByTelegramUserID(ctx, req.TelegramID)
	if err != nil && user != nil {
		user.UpdateLastSeen()
		if err := u.userRepo.Update(ctx, user); err != nil {
			return nil, err
		}

		return mapUserToDTO(user), nil
	}

	// Create new user
	return u.CreateUser(ctx, req)
}

// mapUserToDTO function maps a User entity to a UserResponse DTO
func mapUserToDTO(user *entity.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:         user.ID,
		TelegramID: user.TelegramUserID,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		IsActive:   user.IsActive,
		IsBot:      user.IsBot,
		LastSeenAt: user.LastSeenAt,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}
