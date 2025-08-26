package entity

import (
	"time"

	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
)

// ChatType represents the type of chat in Telegram.
type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSupergroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

// Chat represents a Telegram chat entity.
type Chat struct {
	// Primary key UUID + timestamps created/updated/deleted
	BaseEntityUUID

	// ID Chat in telegram (unique, not null)
	TelegramChatID types.TelegramChatID `json:"telegram_chat_id" gorm:"uniqueIndex;not null"`

	// Type chat: private, group, supergroup, channel
	Type ChatType `json:"type" gorm:"type:varchar(20);not null"`

	// Information can null → pointer
	FirstName   *string `json:"first_name" gorm:"type:varchar(100)"`
	LastName    *string `json:"last_name" gorm:"type:varchar(100)"`
	Title       *string `json:"title" gorm:"type:varchar(255)"`
	Username    *string `json:"username" gorm:"type:varchar(100);index"`
	Description *string `json:"description" gorm:"type:text"`

	// Status chat active/inactive
	IsActive bool `json:"is_active" gorm:"default:true"`
}

// NewChat creates a new Chat instance with the provided Telegram chat ID and type.
func NewChat(telegramChatID types.TelegramChatID, chatType ChatType) *Chat {
	return &Chat{
		TelegramChatID: telegramChatID,
		Type:           chatType,
		IsActive:       true,
	}
}

func (c *Chat) UpdateInfo(title *string, username *string, description *string) {
	if title != nil && *title != "" {
		c.Title = title
	}
	if username != nil && *username != "" {
		c.Username = username
	}
	if description != nil && *description != "" {
		c.Description = description
	}
}

// Deactivate sets the chat's IsActive status to false.
func (c *Chat) Deactivate() {
	c.IsActive = false
	c.UpdatedAt = time.Now()
}

// Activate sets the chat's IsActive status to true.
func (c *Chat) Activate() {
	c.IsActive = true
	c.UpdatedAt = time.Now()
}
