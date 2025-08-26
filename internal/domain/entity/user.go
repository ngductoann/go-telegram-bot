package entity

import (
	"time"

	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
)

// User represents a user entity in the system, typically corresponding to a Telegram user.
type User struct {
	BaseEntityUUID

	// Telegram user ID (unique, not null)
	TelegramUserID types.TelegramUserID `json:"telegram_user_id" gorm:"uniqueIndex;not null"`

	// information user
	Username     *string `json:"username,omitempty" gorm:"type:varchar(100);index"`
	FirstName    string  `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName     *string `json:"last_name,omitempty" gorm:"type:varchar(100)"`
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(10);default:'en';not null"`

	// Status user active/inactive, bot/human
	IsActive   bool       `json:"is_active" gorm:"default:true;not null"`
	IsBot      bool       `json:"is_bot" gorm:"default:false;not null"`
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`
}

// NewUser creates a new User instance with the provided Telegram user ID, first name, and bot status.
// IsActive defaults to true, while other fields can be set later.
func NewUser(userTelegram types.TelegramUserID, firstName string, isBot bool) *User {
	return &User{
		TelegramUserID: userTelegram,
		FirstName:      firstName,
		IsActive:       true,  // default active
		IsBot:          isBot, // bot status
	}
}

// UpdateLastSeen updates the user's last seen timestamp to the current time.
func (u *User) UpdateLastSeen() {
	now := time.Now()
	u.LastSeenAt = &now
	u.UpdatedAt = now
}

// Deactivate sets the user's IsActive status to false and updates the UpdatedAt timestamp to the current time.
func (u *User) Deactivate() {
	u.IsActive = false
	u.UpdatedAt = time.Now()
}

// Activate sets the user's IsActive status to true and updates the UpdatedAt timestamp to the current time.
func (u *User) Activate() {
	u.IsActive = true
	u.UpdatedAt = time.Now()
}

// UpdateUser updates the user's information based on the provided User instance.
// It only updates fields that are non-empty or non-nil in the provided User instance.
func (u *User) UpdateUser(user *User) {
	if user.FirstName != "" {
		u.FirstName = user.FirstName
	}
	if user.Username != nil && *user.Username != "" {
		u.Username = user.Username
	}
	if user.LastName != nil && *user.LastName != "" {
		u.LastName = user.LastName
	}
	if user.LanguageCode != nil && *user.LanguageCode != "" {
		u.LanguageCode = user.LanguageCode
	}
	u.IsBot = user.IsBot
}
