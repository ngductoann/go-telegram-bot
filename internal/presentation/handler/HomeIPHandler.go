package handler

import (
	"context"
	"go-telegram-bot/internal/application/service"
)

// HomeIPCommandHandler handles the /home_ip command
type HomeIPCommandHandler struct {
	botAppService *service.BotApplicationService
}

// Handle processes the /home_ip command
func NewHomeIPCommandHandler(
	botAppService *service.BotApplicationService,
) *HomeIPCommandHandler {
	return &HomeIPCommandHandler{
		botAppService: botAppService,
	}
}

// Handle processes the /home_ip command
func (h *HomeIPCommandHandler) Handle(ctx context.Context, chatID int64) error {
	return h.botAppService.HandleHomeIPCommand(ctx, chatID)
}

// GetCommand returns the command string that this handler processes
func (h *HomeIPCommandHandler) GetCommand() string {
	return "/home_ip"
}
