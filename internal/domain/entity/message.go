package entity

import (
	"time"

	"github.com/google/uuid"
)

// Message represents a Telegram message stored in the system.
type Message struct {
	BaseEntityUUID

	// Message info
	TelegramMsgID int64     `json:"telegram_msg_id" gorm:"not null;index"` // ID message in Telegram
	ChatID        uuid.UUID `json:"chat_id" gorm:"type:uuid;not null;index"`
	UserID        uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	Content       string    `json:"content,omitempty" gorm:"type:text"`
	MessageType   string    `json:"message_type" gorm:"type:varchar(50);not null"`

	ReplyToID *uuid.UUID `json:"reply_to_id,omitempty" gorm:"type:uuid;index"`

	// Flags
	IsEdited  bool `json:"is_edited" gorm:"default:false;not null"`
	IsDeleted bool `json:"is_deleted" gorm:"default:false;not null"`

	// Relationships (link with UUID)
	User    User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Chat    Chat     `json:"chat,omitempty" gorm:"foreignKey:ChatID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReplyTo *Message `json:"reply_to,omitempty" gorm:"foreignKey:ReplyToID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// NewMessage creates a new Message instance with the provided details.
func NewMessage(
	TelegramMsgID int64,
	telegramChatID uuid.UUID,
	telegramUserID uuid.UUID,
	content string,
	messageType string,
) *Message {
	return &Message{
		TelegramMsgID: TelegramMsgID,
		ChatID:        telegramChatID,
		UserID:        telegramUserID,
		Content:       content,
		MessageType:   messageType,
	}
}

// EditMessage updates the content of the message and marks it as edited.
func (m *Message) EditMessage(newContent string) {
	m.Content = newContent
	m.IsEdited = true
	m.UpdatedAt = time.Now()
}

// DeleteMessage marks the message as deleted.
func (m *Message) DeleteMessage() {
	m.IsDeleted = true
	m.UpdatedAt = time.Now()
}
