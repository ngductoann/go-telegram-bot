package handler

import (
	"context"
	"go-telegram-bot/internal/application/service"
)

type HelpCommandHandler struct {
	botAppService *service.BotApplicationService
}

// NewHelpCommandHandler creates a new instance of HelpCommandHandler
func NewHelpCommandHandler(
	botAppService *service.BotApplicationService,
) *HelpCommandHandler {
	return &HelpCommandHandler{
		botAppService: botAppService,
	}
}

// Handle processes the /help command
func (h *HelpCommandHandler) Handle(ctx context.Context, chatID int64) error {
	return h.botAppService.HandleHelpCommand(ctx, chatID)
}

// GetCommand returns the command string that this handler processes
func (h *HelpCommandHandler) GetCommand() string {
	return "/help"
}
