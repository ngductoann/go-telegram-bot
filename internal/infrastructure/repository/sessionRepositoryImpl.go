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

// sessionRepository implements the SessionRepository interface
type sessionRepository struct {
	db *gorm.DB
}

// NewSessionRepository creates a new session repository
func NewSessionRepository(db *gorm.DB) repository.SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

// GetByUserAndChat retrieves a session by user ID and chat ID
func (r *sessionRepository) GetByUserAndChat(ctx context.Context, userID uuid.UUID, chatID types.TelegramChatID) (*entity.UserSession, error) {
	var session entity.UserSession
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND chat_id = ?", userID, chatID).
		First(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrSessionNotFound
		}
		return nil, err
	}
	return &session, nil
}

// GetByID retrieves a session by its ID
func (r *sessionRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.UserSession, error) {
	var session entity.UserSession
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrSessionNotFound
		}
		return nil, err
	}
	return &session, nil
}

// Create inserts a new session into the database
func (r *sessionRepository) Create(ctx context.Context, session *entity.UserSession) error {
	return r.db.WithContext(ctx).Create(session).Error
}

// Update modifies an existing session in the database
func (r *sessionRepository) Update(ctx context.Context, session *entity.UserSession) error {
	return r.db.WithContext(ctx).Save(session).Error
}

// Delete removes a session by its ID
func (r *sessionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.UserSession{}).Error
}

// DeleteExpired removes all expired sessions from the database
func (r *sessionRepository) DeleteExpired(ctx context.Context) error {
	return r.db.WithContext(ctx).
		Where("expires_at < ?", time.Now()).
		Delete(&entity.UserSession{}).Error
}

// ExtendSession extends the expiration time of a session by a given duration
func (r *sessionRepository) ExtendSession(ctx context.Context, id uuid.UUID, duration time.Duration) error {
	return r.db.WithContext(ctx).Model(&entity.UserSession{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"expires_at": time.Now().Add(duration),
			"updated_at": time.Now(),
		}).Error
}
