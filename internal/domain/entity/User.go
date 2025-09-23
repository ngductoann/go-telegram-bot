package entity

import (
	"time"

	"go-telegram-bot/internal/domain/types"
)

// User represents a Telegram user in the system
type User struct {
	BaseEntityWithUUID

	TelegramUserID types.TelegramUserID `json:"telegram_user_id" gorm:"type:bigint;uniqueIndex;not null"`

	Username  *string `json:"username" gorm:"type:varchar(255);not null"`
	FirstName string  `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName  *string `json:"last_name,omitempty" gorm:"type:varchar(255)"`

	IsActive   bool       `json:"is_active" gorm:"type:boolean;default:true;not null"`
	IsBot      bool       `json:"is_bot" gorm:"type:boolean;default:false;not null"`
	LastSeenAt *time.Time `json:"last_seen_at" gorm:"type:timestamp;default:current_timestamp;not null"`
}

// NewUser creates a new User instance
func NewUser(
	telegramUserID types.TelegramUserID,
	firstName string,
) *User {
	return &User{
		TelegramUserID: telegramUserID,
		FirstName:      firstName,
	}
}

// UpdateLastSeen updates the LastSeenAt field to the current time
func (u *User) UpdateLastSeen() {
	now := time.Now()
	u.LastSeenAt = &now
}

// SetIsBot sets the IsBot field to true
func (u *User) SetIsBot() {
	u.IsBot = true
}

// DeactivateUser sets the IsActive field to false
func (u *User) DeactivateUser() {
	u.IsActive = false
}

// ActivateUser sets the IsActive field to true
func (u *User) ActivateUser() {
	u.IsActive = true
}
