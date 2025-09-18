package presentation

import (
	"context"
	"time"

	"go-telegram-bot/internal/application/service"
	"go-telegram-bot/internal/domain/entity"
	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/presentation/handler"
	"go-telegram-bot/internal/presentation/middleware"
	"go-telegram-bot/internal/presentation/router"
)

// TelegramHandler handles Telegram bot updates and commands
type TelegramHandler struct {
	bot               domainService.TelegramBotService
	commandRouter     *router.CommandRouter
	loggingMiddleware *middleware.LoggingMiddleware
	errorMiddleware   *middleware.ErrorHandlingMiddleware
	logger            domainService.Logger
}

// NewTelegramHandler creates a new instance of TelegramHandler
func NewTelegramHandler(
	botAppService *service.BotApplicationService,
	bot domainService.TelegramBotService,
	logger domainService.Logger,
) *TelegramHandler {
	// Create handlers - all now depend on botAppService
	startHandler := handler.NewStartCommandHandler(botAppService)
	helpHandler := handler.NewHelpCommandHandler(botAppService)
	homeIPHandler := handler.NewHomeIPCommandHandler(botAppService)
	unknownHandler := handler.NewUnknownCommandHandler(botAppService)

	// Create router and register handlers
	commandRouter := router.NewCommandRouter(unknownHandler)
	commandRouter.RegisterHandler(startHandler)
	commandRouter.RegisterHandler(helpHandler)
	commandRouter.RegisterHandler(homeIPHandler)

	// Create middleware
	loggingMiddleware := middleware.NewLoggingMiddleware(logger)
	errorMiddleware := middleware.NewErrorHandlingMiddleware(bot, logger)

	return &TelegramHandler{
		bot:               bot,
		commandRouter:     commandRouter,
		loggingMiddleware: loggingMiddleware,
		errorMiddleware:   errorMiddleware,
		logger:            logger,
	}
}

// ProcessUpdate processes incoming updates from Telegram
func (h *TelegramHandler) ProcessUpdate(
	ctx context.Context,
	update entity.TelegramUpdate,
) error {
	// Apply middleware chain: logging -> error handling -> command routing
	return h.loggingMiddleware.Process(ctx, update, func(ctx context.Context, update entity.TelegramUpdate) error {
		return h.errorMiddleware.Process(ctx, update, func(ctx context.Context, update entity.TelegramUpdate) error {
			return h.commandRouter.Route(ctx, update)
		})
	})
}

// StartPolling starts polling for updates from Telegram
func (h *TelegramHandler) StartPolling(ctx context.Context) error {
	h.logger.Info("🤖 Starting to poll for Telegram updates...")

	// Delete any existing webhook to avoid 409 conflicts
	h.logger.Info("🔧 Clearing any existing webhook...")
	if err := h.bot.DeleteWebhook(ctx); err != nil {
		h.logger.Warn("⚠️  Failed to delete webhook (this might be normal if no webhook was set)", "error", err)
	} else {
		h.logger.Info("✅ Webhook cleared successfully")
	}

	var offset int64 = 0

	for {
		select {
		case <-ctx.Done():
			h.logger.Info("🛑 Bot is stopping...")
			return ctx.Err()
		default:
			// Lấy updates từ Telegram
			updates, err := h.bot.GetUpdates(ctx, offset)
			if err != nil {
				// Only log non-timeout errors as errors
				h.logger.Error("⚠️  Lỗi khi lấy updates (sẽ retry)", "error", err)
				time.Sleep(5 * time.Second) // Đợi 5 giây trước khi thử lại
				continue
			}

			// Xử lý từng update
			for _, update := range updates {
				// Cập nhật offset để không nhận lại update cũ
				if update.UpdateID >= offset {
					offset = update.UpdateID + 1
				}

				// Xử lý update trong goroutine riêng để không block
				go func(update entity.TelegramUpdate) {
					updateCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
					defer cancel()

					if err := h.ProcessUpdate(updateCtx, update); err != nil {
						h.logger.Error("❌ Lỗi khi xử lý update", "update_id", update.UpdateID, "error", err)
					}
				}(update)
			}

			// Nghỉ ngắn để không spam API
			if len(updates) == 0 {
				time.Sleep(1 * time.Second)
			}
		}
	}
}
