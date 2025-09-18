package handler

import (
	"context"
	"go-telegram-bot/internal/domain/entity"
)

// CommandHandler defines the interface for handling specific commands
type CommandHandler interface {
	Handle(ctx context.Context, chatID int64) error
	GetCommand() string
}

type UpdateHandler interface {
	ProcessUpdate(ctx context.Context, update entity.TelegramUpdate) error
}

type TelegramHandler interface {
	UpdateHandler
	StartPolling(ctx context.Context) error
}
