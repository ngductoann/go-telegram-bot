package usecase

import (
	"context"
	"fmt"
	"time"

	"go-telegram-bot/internal/domain/entity"
	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/shared/util"
)

type UnknownCommandUseCase struct {
	telegramBot domainService.TelegramBotService
	logger      domainService.Logger
}

// NewUnknownCommandUseCase creates a new instance of UnknownCommandUseCase
func NewUnknownCommandUseCase(
	telegramBot domainService.TelegramBotService,
	logger domainService.Logger,
) *UnknownCommandUseCase {
	return &UnknownCommandUseCase{
		telegramBot: telegramBot,
		logger:      logger,
	}
}

// Execute returns the unknown command message
func (u *UnknownCommandUseCase) Execute(ctx context.Context, chatID int64) error {
	// 	message := `❓ Lệnh không xác định.
	//
	// Vui lòng sử dụng /help để xem danh sách lệnh hợp lệ.`

	message := fmt.Sprintf(
		"❓ Lệnh không xác định\\.\n\n"+
			"Vui lòng sử dụng `%s` để xem danh sách lệnh hợp lệ\\.",
		util.EscapeMarkdownV2("/help"),
	)

	botMsg := &entity.BotMessage{
		ChatID:      chatID,
		Text:        message,
		SentAt:      time.Now().Unix(),
		MessageType: entity.MessageTypeCommand,
		ParseMode:   "MarkdownV2",
	}

	if err := u.telegramBot.SendBotMessage(ctx, botMsg); err != nil {
		return fmt.Errorf("failed to send unknown command message: %w", err)
	}

	return nil
}
