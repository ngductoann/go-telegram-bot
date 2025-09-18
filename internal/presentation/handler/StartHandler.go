package handler

import (
	"context"
	"go-telegram-bot/internal/application/service"
)

// StartCommandHandler handles the /start command
type StartCommandHandler struct {
	botAppService *service.BotApplicationService
}

// NewStartCommandHandler creates a new instance of StartCommandHandler
func NewStartCommandHandler(
	botAppService *service.BotApplicationService,
) *StartCommandHandler {
	return &StartCommandHandler{
		botAppService: botAppService,
	}
}

// Handle processes the /start command
func (h *StartCommandHandler) Handle(ctx context.Context, chatID int64) error {
	return h.botAppService.HandleStartCommand(ctx, chatID)
}

// GetCommand returns the command string that this handler processes
func (h *StartCommandHandler) GetCommand() string {
	return "/start"
}
