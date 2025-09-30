package service

import (
	"context"

	"go-telegram-bot/internal/domain/types"
)

// BotUseCase defines the interface for bot business logic operations
type BotUseCase interface {
	// HandleHomeIPCommand processes the /home_ip command
	HandleHomeIPCommand(ctx context.Context, chatID types.TelegramChatID) (*types.SendMessageResponse, error)

	// HandleStartCommand processes the /start command
	HandleStartCommand(ctx context.Context, chatID types.TelegramChatID) (*types.SendMessageResponse, error)

	// HandleHelpCommand processes the /help command
	HandleHelpCommand(ctx context.Context, chatID types.TelegramChatID) (*types.SendMessageResponse, error)

	// HandleUnknownCommand processes unknown or invalid commands
	HandleUnknownCommand(ctx context.Context, chatID types.TelegramChatID) (*types.SendMessageResponse, error)

	// ProcessUpdate processes incoming Telegram updates
	ProcessUpdate(ctx context.Context, update types.TelegramUpdate) error

	// ValidateUpdate validates the structure and content of an update
	ValidateUpdate(update types.TelegramUpdate) error
}
