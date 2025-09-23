package entity

import (
	"fmt"
	"time"

	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Message represents a message entity in the system
type Message struct {
	BaseEntityWithUUID

	TelegramMessageID int64              `json:"telegram_message_id" gorm:"type:bigint;not null;uniqueIndex"`
	ChatID            uuid.UUID          `json:"chat_id" gorm:"type:bigint;not null;index"`
	UserID            uuid.UUID          `json:"user_id" gorm:"type:uuid;not null;index"`
	Content           string             `json:"content" gorm:"type:text;not null"`
	MessageType       *types.MessageType `json:"message_type" gorm:"type:varchar(50);not null;default:'text'"`
	ParseMode         *types.ParseMode   `json:"parse_mode,omitempty" gorm:"type:varchar(20);default:null"`
	Command           *types.Command     `json:"command,omitempty" gorm:"type:varchar(50);default:null"`
	ReplyToID         *uuid.UUID         `json:"reply_to_id,omitempty" gorm:"type:uuid;default:null"`

	IsDeleted bool `json:"is_deleted" gorm:"type:boolean;not null;default:false"`
	IsEdited  bool `json:"is_edited" gorm:"type:boolean;not null;default:false"`

	User    *User    `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Chat    *Chat    `json:"chat,omitempty" gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReplyTo *Message `json:"reply_to,omitempty" gorm:"foreignKey:ReplyToID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// NewMessage creates a new Message instance
func NewMessage(
	telegramMessageID int64,
	chatID, userID uuid.UUID,
	content string,
	messageType *types.MessageType,
	parseMode *types.ParseMode,
	replyToID *uuid.UUID,
) *Message {
	return &Message{
		TelegramMessageID: telegramMessageID,
		ChatID:            chatID,
		UserID:            userID,
		Content:           content,
		MessageType:       messageType,
		ParseMode:         parseMode,
		ReplyToID:         replyToID,
	}
}

// NewMarkdownMessage creates a new Message instance with MarkdownV2 parse mode
func (m *Message) NewMarkdownMessage(
	telegramMessageID int64,
	chatID, userID uuid.UUID,
	content string,
	replyToID *uuid.UUID,
) *Message {
	messageType := types.MessageTypeText
	parseMode := types.ParseModeMarkdownV2
	return &Message{
		TelegramMessageID: telegramMessageID,
		ChatID:            chatID,
		UserID:            userID,
		Content:           content,
		MessageType:       &messageType,
		ParseMode:         &parseMode,
	}
}

// IsValid marks the message as edited
func (m *Message) Edited() {
	m.IsEdited = true
}

// Delete marks the message as deleted
func (m *Message) Delete() {
	m.IsDeleted = true
}

// BeforeSave is a GORM hook that runs before saving a Message
func (m *Message) BeforeSave(tx *gorm.DB) error {
	if m.MessageType == nil || !m.MessageType.IsValid() {
		var mt string
		if m.MessageType != nil {
			mt = string(*m.MessageType)
		}
		return fmt.Errorf("invalid message type: %s", mt)
	}
	if m.ParseMode != nil {
		if !(*m.ParseMode).IsValid() {
			return fmt.Errorf("invalid parse mode: %s", *m.ParseMode)
		}
	}
	return nil
}

// IsCommand checks if the message is a command type
func (m *Message) IsCommand() bool {
	return m.MessageType != nil && *m.MessageType == types.MessageTypeCommand && m.Command != nil && m.Command.IsValid()
}

// GetSendTime returns the time the message was created
func (m *Message) GetSendTime() time.Time {
	return m.CreatedAt
}
