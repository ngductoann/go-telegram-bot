package entity

import (
	"time"
)

// MessageType represents the type of bot message
type MessageType string

const (
	MessageTypeText    MessageType = "text"
	MessageTypeCommand MessageType = "command"
	MessageTypeError   MessageType = "error"
	MessageTypeInfo    MessageType = "info"
)

// BotMessage represents a message sent by a bot in a chat application.
type BotMessage struct {
	ChatID      int64       `json:"chat_id"`
	Text        string      `json:"text"`
	SentAt      int64       `json:"sent_at"`           // Unix timestamp
	Command     string      `json:"command,omitempty"` // Optional command associated with the message
	MessageType MessageType `json:"message_type"`
	ParseMode   string      `json:"parse_mode,omitempty"` // Markdown, HTML, etc.
}

// NewBotMessage creates a new bot message with current timestamp
func NewBotMessage(chatID int64, text string, messageType MessageType) *BotMessage {
	return &BotMessage{
		ChatID:      chatID,
		Text:        text,
		SentAt:      time.Now().Unix(),
		MessageType: messageType,
	}
}

// NewCommandMessage creates a new command-type bot message
func NewCommandMessage(chatID int64, text string, command string) *BotMessage {
	msg := NewBotMessage(chatID, text, MessageTypeCommand)
	msg.Command = command
	return msg
}

// NewMarkdownMessage creates a new bot message with Markdown formatting
func NewMarkdownMessage(chatID int64, text string, messageType MessageType) *BotMessage {
	msg := NewBotMessage(chatID, text, messageType)
	msg.ParseMode = "Markdown"
	return msg
}

// IsCommand checks if the message is a command type
func (m *BotMessage) IsCommand() bool {
	return m.MessageType == MessageTypeCommand && m.Command != ""
}

// IsValid validates the bot message fields
func (m *BotMessage) IsValid() bool {
	return m.ChatID != 0 && m.Text != "" && m.SentAt > 0
}

// GetSentTime returns the sent time as a time.Time object
func (m *BotMessage) GetSentTime() time.Time {
	return time.Unix(m.SentAt, 0)
}
