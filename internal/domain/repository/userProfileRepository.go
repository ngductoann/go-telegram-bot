package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngductoann/go-telegram-bot/internal/domain/entity"
)

// UserProfileRepository defines the interface for user profile repository
type UserProfileRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserProfile, error)
	Create(ctx context.Context, profile *entity.UserProfile) error
	Update(ctx context.Context, profile *entity.UserProfile) error
	Delete(ctx context.Context, userID uuid.UUID) error
	GetOrCreate(ctx context.Context, userID uuid.UUID) (*entity.UserProfile, error)
}
