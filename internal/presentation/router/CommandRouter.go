package router

import (
	"context"
	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/presentation/handler"
	"strings"
)

// CommandRouter routes commands to their respective handlers
type CommandRouter struct {
	handlers       map[string]handler.CommandHandler
	unknownHandler handler.CommandHandler
}

// NewCommandRouter creates a new CommandRouter
func NewCommandRouter(
	unknownHandler handler.CommandHandler,
) *CommandRouter {

	return &CommandRouter{
		handlers:       make(map[string]handler.CommandHandler),
		unknownHandler: unknownHandler,
	}
}

// RegisterHandler registers a command handler
func (r *CommandRouter) RegisterHandler(cmdHandler handler.CommandHandler) {
	r.handlers[cmdHandler.GetCommand()] = cmdHandler
}

func (r *CommandRouter) Route(
	ctx context.Context,
	update entity.TelegramUpdate,
) error {
	// Check if the update contains a message with a command
	if update.Message == nil {
		return nil // Ignore updates without messages
	}

	message := update.Message
	chatID := message.Chat.ID
	text := message.Text

	// Handler for unknown commands
	if strings.HasPrefix(text, "/") {
		if handler, exists := r.handlers[text]; exists {
			return handler.Handle(ctx, chatID)
		} else {
			// Use the unknown command handler
			return r.unknownHandler.Handle(ctx, chatID)
		}
	}

	// Ignore non-command messages
	return nil
}
