package entity

import "go-telegram-bot/internal/domain/types"

type Chat struct {
	BaseEntityWithUUID
	TelegramChatID types.TelegramChatID `json:"telegram_chat_id" gorm:"type:bigint;not null;uniqueIndex"`
	ChatType       types.ChatType       `json:"chat_type" gorm:"type:varchar(32);not null"`
	Title          *string              `json:"title,omitempty" gorm:"type:varchar(255);default:null"`
	Username       *string              `json:"username,omitempty" gorm:"type:varchar(255);default:null"`
	Description    *string              `json:"description,omitempty" gorm:"type:text;default:null"`
	IsActive       bool                 `json:"is_active" gorm:"type:boolean;not null;default:true"`
}

func NewChat(telegramChatID types.TelegramChatID, chatType types.ChatType) *Chat {
	return &Chat{
		TelegramChatID: telegramChatID,
		ChatType:       chatType,
	}
}

func (c *Chat) Deactivate() {
	c.IsActive = false
}

func (c *Chat) Activate() {
	c.IsActive = true
}

func (c *Chat) UpdateDetails(title, username, description *string) {
	if title != nil {
		c.Title = title
	}
	if username != nil {
		c.Username = username
	}
	if description != nil {
		c.Description = description
	}
}
