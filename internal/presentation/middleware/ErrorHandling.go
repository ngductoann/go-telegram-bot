package middleware

import (
	"context"

	"go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/domain/types"
)

// ErrorHandlingMiddleware handles errors and sends user-friendly messages
type ErrorHandlingMiddleware struct {
	telegramBot service.TelegramBotService
	logger      service.Logger
}

// NewErrorHandlingMiddleware creates a new error handling middleware
func NewErrorHandlingMiddleware(
	telegramBot service.TelegramBotService, logger service.Logger,
) *ErrorHandlingMiddleware {
	return &ErrorHandlingMiddleware{
		telegramBot: telegramBot,
		logger:      logger,
	}
}

// Process handles errors from the next handler and sends user-friendly messages
func (m *ErrorHandlingMiddleware) Process(
	ctx context.Context,
	update types.TelegramUpdate,
	next func(context.Context, types.TelegramUpdate) error,
) error {
	err := next(ctx, update)
	if err != nil {
		// Log the error
		m.logger.Error("🔥 Error in handler chain:", "error", err)

		// Send user-friendly error message if it's a message update
		if update.Message != nil && update.Message.Chat != nil {
			errorMsg := "⚠️ Đã xảy ra lỗi khi xử lý yêu cầu của bạn. Vui lòng thử lại sau."

			_, err := m.telegramBot.SendMessageWithResponse(ctx, &types.SendMessageRequest{
				ChatID: update.Message.Chat.ID,
				Text:   errorMsg,
			})
			if err != nil {
				m.logger.Error("❌ Failed to send error message:", "error", err)
			}
		}
	}

	return err
}
