package repository

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
)

type UserProfileRepository interface {
	// Define methods for user profile repository
	FindByUserID(ctx context.Context, userID string) (entity.UserProfile, error)
	Create(ctx context.Context, profile *entity.UserProfile) error
	Update(ctx context.Context, profile *entity.UserProfile) error
	Delete(ctx context.Context, profile *entity.UserProfile) error
}
