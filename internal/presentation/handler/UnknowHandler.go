package handler

import (
	"context"
	"go-telegram-bot/internal/application/service"
)

type UnknownCommandHandler struct {
	botAppService *service.BotApplicationService
}

func NewUnknownCommandHandler(
	botAppService *service.BotApplicationService,
) *UnknownCommandHandler {
	return &UnknownCommandHandler{
		botAppService: botAppService,
	}
}

func (h *UnknownCommandHandler) Handle(ctx context.Context, chatID int64) error {
	return h.botAppService.HandleUnknownCommand(ctx, chatID)
}

func (h *UnknownCommandHandler) GetCommand() string {
	return "unknown"
}
