package usecase

import (
	"context"
	"fmt"
	"time"

	"go-telegram-bot/internal/domain/entity"
	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/shared/util"
)

type StartUseCase struct {
	telegramBot domainService.TelegramBotService
	logger      domainService.Logger
}

func NewStartUseCase(telegramBot domainService.TelegramBotService, logger domainService.Logger) *StartUseCase {
	return &StartUseCase{
		telegramBot: telegramBot,
		logger:      logger,
	}
}

func (uc *StartUseCase) Execute(ctx context.Context, chatID int64) error {
	message := fmt.Sprintf("*%s*\n\n🤖 *%s*\n"+
		"• `%s` \\- Xem thông tin IP local và WAN của máy\n"+
		"• `%s` \\- Hiển thị hướng dẫn sử dụng\n",
		util.EscapeMarkdownV2("👋 Chào mừng bạn đến với Bot IP!"),
		util.EscapeMarkdownV2("Các lệnh có sẵn:"),
		util.EscapeMarkdownV2("/home_ip"),
		util.EscapeMarkdownV2("/help"),
	)

	botMsg := &entity.BotMessage{
		ChatID:      chatID,
		Text:        message,
		MessageType: entity.MessageTypeInfo,
		Command:     "/start",
		ParseMode:   "MarkdownV2",
		SentAt:      time.Now().Unix(),
	}

	if err := uc.telegramBot.SendBotMessage(ctx, botMsg); err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}

	return nil
}
