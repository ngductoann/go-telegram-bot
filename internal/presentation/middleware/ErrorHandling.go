package middleware

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
	domainService "go-telegram-bot/internal/domain/service"
)

// ErrorHandlingMiddleware handles errors and sends user-friendly messages
type ErrorHandlingMiddleware struct {
	telegramBot domainService.TelegramBotService
	logger      domainService.Logger
}

// NewErrorHandlingMiddleware creates a new error handling middleware
func NewErrorHandlingMiddleware(telegramBot domainService.TelegramBotService, logger domainService.Logger) *ErrorHandlingMiddleware {
	return &ErrorHandlingMiddleware{
		telegramBot: telegramBot,
		logger:      logger,
	}
}

// Process handles errors from the next handler and sends user-friendly messages
func (m *ErrorHandlingMiddleware) Process(ctx context.Context, update entity.TelegramUpdate, next func(context.Context, entity.TelegramUpdate) error) error {
	err := next(ctx, update)

	if err != nil {
		// Log the error
		m.logger.Error("🔥 Error in handler chain:", "error", err)

		// Send user-friendly error message if it's a message update
		if update.Message != nil && update.Message.Chat != nil {
			errorMsg := "⚠️ Đã xảy ra lỗi khi xử lý yêu cầu của bạn. Vui lòng thử lại sau."

			if sendErr := m.telegramBot.SendMessage(ctx, update.Message.Chat.ID, errorMsg); sendErr != nil {
				m.logger.Error("❌ Failed to send error message:", "error", sendErr)
			}
		}
	}

	return err
}
