package delivery

import (
	"context"
	"log"
	"time"

	"go-telegram-bot/internal/domain/entity"
	domainService "go-telegram-bot/internal/domain/service"
)

// telegramHandler handles Telegram bot updates and commands
type telegramHandler struct {
	botUseCase domainService.BotUseCase
	bot        domainService.TelegramBotService
}

// NewTelegramHandler creates a new instance of telegramHandler
func NewTelegramHandler(botUseCase domainService.BotUseCase, bot domainService.TelegramBotService) *telegramHandler {
	return &telegramHandler{
		botUseCase: botUseCase,
		bot:        bot,
	}
}

// StartPolling starts polling for updates from Telegram
func (h *telegramHandler) StartPolling(ctx context.Context) error {
	log.Println("🤖 Starting to poll for Telegram updates...")

	var offset int64 = 0

	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 Bot đang dừng...")
			return ctx.Err()
		default:
			// Lấy updates từ Telegram
			updates, err := h.bot.GetUpdates(ctx, offset)
			if err != nil {
				// Only log non-timeout errors as errors
				log.Printf("⚠️  Lỗi khi lấy updates (sẽ retry): %v", err)
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

					if err := h.botUseCase.ProcessUpdate(updateCtx, update); err != nil {
						log.Printf("❌ Lỗi khi xử lý update %d: %v", update.UpdateID, err)
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
