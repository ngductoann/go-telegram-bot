package service

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
)

// TelegramBotService defines the interface for Telegram bot operations
type TelegramBotService interface {
	// SendMessage sends a text message to a specific chat
	SendMessage(ctx context.Context, chatID int64, text string) error

	// SendMessageWithParseMode sends a message with specific parse mode (Markdown, HTML)
	SendMessageWithParseMode(ctx context.Context, chatID int64, text string, parseMode string) error

	// SendBotMessage sends a structured bot message
	SendBotMessage(ctx context.Context, message *entity.BotMessage) error

	// GetUpdates retrieves pending updates from Telegram
	GetUpdates(ctx context.Context, offset int64) ([]entity.TelegramUpdate, error)

	// GetUpdatesWithLimit retrieves updates with a specific limit
	GetUpdatesWithLimit(ctx context.Context, offset int64, limit int) ([]entity.TelegramUpdate, error)

	// GetMe returns information about the bot
	GetMe(ctx context.Context) (*entity.User, error)

	// DeleteWebhook removes the webhook integration
	DeleteWebhook(ctx context.Context) error
}
