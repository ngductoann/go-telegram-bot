package middleware

import (
	"context"
	"time"

	"go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/domain/types"
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
	update types.TelegramUpdate,
	next func(ctx context.Context, update types.TelegramUpdate) error,
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
