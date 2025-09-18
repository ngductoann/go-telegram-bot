package service

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
)

// BotUseCase defines the interface for bot business logic operations
type BotUseCase interface {
	// HandleHomeIPCommand processes the /home_ip command
	HandleHomeIPCommand(ctx context.Context, chatID int64) error

	// HandleStartCommand processes the /start command
	HandleStartCommand(ctx context.Context, chatID int64) error

	// HandleHelpCommand processes the /help command
	HandleHelpCommand(ctx context.Context, chatID int64) error

	// HandleUnknownCommand processes unknown or invalid commands
	HandleUnknownCommand(ctx context.Context, chatID int64) error

	// ProcessUpdate processes incoming Telegram updates
	ProcessUpdate(ctx context.Context, update entity.TelegramUpdate) error

	// ValidateUpdate validates the structure and content of an update
	ValidateUpdate(update entity.TelegramUpdate) error

	// ExtractCommand extracts command from message text
	ExtractCommand(text string) (command string, args []string)
}
