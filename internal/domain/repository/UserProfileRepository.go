package repository

import (
	"context"

	"go-telegram-bot/internal/domain/entity"

	"github.com/google/uuid"
)

type UserProfileRepository interface {
	// Define methods for user profile repository
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserProfile, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.UserProfile, error)
	Create(ctx context.Context, profile *entity.UserProfile) error
	Update(ctx context.Context, profile *entity.UserProfile) error
	Delete(ctx context.Context, profile *entity.UserProfile) error
}
