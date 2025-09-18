package middleware

import (
	"context"
	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/service"
	"time"
)

type LoggingMiddleware struct {
	logger service.Logger
}

func NewLoggingMiddleware(logger service.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger: logger,
	}
}

func (m *LoggingMiddleware) Process(
	ctx context.Context,
	update entity.TelegramUpdate,
	next func(ctx context.Context, update entity.TelegramUpdate) error,
) error {
	start := time.Now()

	// Log the incoming update
	if update.Message != nil {
		m.logger.Info("Received message", "chat_id", update.Message.Chat.ID, "text", update.Message.Text)
	} else {
		m.logger.Info("Received update", "update_id", update.UpdateID)
	}

	// Call the next handler in the chain
	err := next(ctx, update)

	// Log the result of processing
	duration := time.Since(start)

	if err != nil {
		m.logger.Error("Error processing update", "error", err, "duration", duration)
	} else {
		m.logger.Info("Processed update successfully", "duration", duration)
	}

	return err
}
