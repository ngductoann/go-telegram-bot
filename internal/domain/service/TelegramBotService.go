package service

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
)

type TelegramBotService interface {
	SendMessage(ctx context.Context, chatID int64, text string) error
	GetUpdates(ctx context.Context, offset int64) ([]entity.TelegramUpdate, error)
}
