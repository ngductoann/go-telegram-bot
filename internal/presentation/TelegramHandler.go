package presentation

import (
	"context"
	"time"

	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/domain/types"
	"go-telegram-bot/internal/presentation/middleware"
	"go-telegram-bot/internal/presentation/service"
)

// TelegramHandler handles Telegram bot updates and commands
type TelegramHandler struct {
	bot               domainService.TelegramBotService
	botAppService     *service.BotApplicationService
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
	// Create middleware
	loggingMiddleware := middleware.NewLoggingMiddleware(logger)
	errorMiddleware := middleware.NewErrorHandlingMiddleware(bot, logger)

	return &TelegramHandler{
		bot:               bot,
		botAppService:     botAppService,
		loggingMiddleware: loggingMiddleware,
		errorMiddleware:   errorMiddleware,
		logger:            logger,
	}
}

// ProcessUpdate processes incoming updates from Telegram
func (h *TelegramHandler) ProcessUpdate(
	ctx context.Context,
	update types.TelegramUpdate,
) error {
	// Apply middleware chain: logging -> error handling -> command routing
	return h.loggingMiddleware.Process(
		ctx, update, func(ctx context.Context, update types.TelegramUpdate) error {
			return h.errorMiddleware.Process(
				ctx, update, func(ctx context.Context, update types.TelegramUpdate) error {
					return h.botAppService.ProcessUpdate(ctx, update)
				})
		})
}

// StartPolling starts polling for updates from Telegram
func (h *TelegramHandler) StartPolling(ctx context.Context) error {
	h.logger.Info("🤖 Starting to poll for Telegram updates...")

	// Delete any existing webhook to avoid 409 conflicts
	h.logger.Info("🔧 Clearing any existing webhook...")
	_, err := h.bot.DeleteWebhookWithResponse(ctx)
	if err != nil {
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
			request := types.GetUpdatesRequest{
				Offset: offset,
			}
			resp, err := h.bot.GetUpdatesWithResponse(ctx, &request)
			if err != nil {
				h.logger.Error("❌ Failed to get updates", "error", err)
				time.Sleep(2 * time.Second)
				continue
			}

			if resp.Result == nil {
				time.Sleep(5 * time.Second)
				continue
			}

			updates := *resp.Result

			// process each update
			for _, update := range updates {
				if update.UpdateID >= offset {
					offset = update.UpdateID + 1
				}

				// process update in a separate goroutine
				go func(update types.TelegramUpdate) {
					updateCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
					defer cancel()

					if err := h.ProcessUpdate(updateCtx, update); err != nil {
						h.logger.Error("❌ Failed to process update", "error", err, "update_id", update.UpdateID)
					}
				}(update)
			}

			if len(updates) == 0 {
				time.Sleep(1 * time.Second)
			}
		}
	}
}
