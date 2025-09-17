package service

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
)

type BotUseCase interface {
	HandleHomeIPCommand(ctx context.Context, chatID int64) error
	ProcessUpdate(ctx context.Context, update entity.TelegramUpdate) error
}
