package entity

import (
	"time"

	"github.com/google/uuid"
)

// UserProfile represents a user profile in the system
type UserProfile struct {
	BaseEntityWithUUID

	UserID     uuid.UUID  `json:"user_id" gorm:"type:uuid;not null;uniqueIndex"`
	DateActive *time.Time `json:"date_active" gorm:"type:timestamp;default:current_timestamp;not null"`

	User User `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// NewUserProfile creates a new UserProfile instance
func NewUserProfile(userID uuid.UUID) *UserProfile {
	return &UserProfile{
		UserID: userID,
	}
}
