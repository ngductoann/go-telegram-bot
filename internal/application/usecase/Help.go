package usecase

import (
	"context"
	"fmt"
	"time"

	"go-telegram-bot/internal/domain/entity"
	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/shared/util"
)

type HelpUseCase struct {
	telegramBot domainService.TelegramBotService
	logger      domainService.Logger
}

func NewHelpUseCase(telegramBot domainService.TelegramBotService, logger domainService.Logger) *HelpUseCase {
	return &HelpUseCase{
		telegramBot: telegramBot,
		logger:      logger,
	}
}

func (h *HelpUseCase) Execute(ctx context.Context, chatID int64) error {
	message := fmt.Sprintf(
		"📖 *Hướng dẫn sử dụng Bot IP*\n\n"+
			"🔹 `%s` \\- Lấy thông tin IP hiện tại:\n"+
			"   • Local IP: Địa chỉ IP trong mạng nội bộ\n"+
			"   • WAN IP: Địa chỉ IP công khai trên Internet\n"+
			"🔹 `%s` \\- Hiển thị lời chào và các lệnh cơ bản\n"+
			"🔹 `%s` \\- Hiển thị hướng dẫn này\n\n"+
			"ℹ️ *Lưu ý:* Bot sẽ tự động cập nhật thông tin IP mỗi khi bạn gọi lệnh\\.",
		util.EscapeMarkdownV2("/home_ip"),
		util.EscapeMarkdownV2("/start"),
		util.EscapeMarkdownV2("/help"),
	)

	botMsg := &entity.BotMessage{
		ChatID:      chatID,
		Text:        message,
		SentAt:      time.Now().Unix(),
		Command:     "/help",
		MessageType: entity.MessageTypeInfo,
		ParseMode:   "MarkdownV2",
	}

	if err := h.telegramBot.SendBotMessage(ctx, botMsg); err != nil {
		return fmt.Errorf("failed to send help message: %w", err)
	}
	return nil
}
