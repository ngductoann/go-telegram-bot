package usecase

import (
	"context"
	"fmt"
	"go-telegram-bot/internal/domain/entity"
	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/shared/util"
	"time"
)

// HomeIPUseCase handles the business logic for home IP related operations
type HomeIPUseCase struct {
	ipService   domainService.IPService
	telegramBot domainService.TelegramBotService
	logger      domainService.Logger
}

// NewHomeIPUseCase creates a new instance of HomeIPUseCase
func NewHomeIPUseCase(
	ipService domainService.IPService,
	telegramBot domainService.TelegramBotService,
	logger domainService.Logger,
) *HomeIPUseCase {
	return &HomeIPUseCase{
		ipService:   ipService,
		telegramBot: telegramBot,
		logger:      logger,
	}
}

func (uc *HomeIPUseCase) Execute(ctx context.Context, chatID int64) error {
	uc.logger.Info("Handling /home_up command for chat ID: %d", chatID)

	// Retrieve IP information from the IP service
	ipInfo, err := uc.ipService.GetIPInfo(ctx)
	if err != nil {
		uc.logger.Error("Failed to get IP info: %v", err)

		errorMsg := "❌ Failed to retrieve IP information. Please try again later."
		botMsg := &entity.BotMessage{
			ChatID:      chatID,
			Text:        errorMsg,
			MessageType: entity.MessageTypeError,
			ParseMode:   "MarkdownV2",
			SentAt:      time.Now().Unix(),
		}
		if sendErr := uc.telegramBot.SendBotMessage(ctx, botMsg); sendErr != nil {
			uc.logger.Error("Failed to send error message to Telegram: %v", sendErr)
		}
		return fmt.Errorf("failed to get IP info: %w", err)
	}

	// Create response message with proper MarkdownV2 formatting
	// Escape only the dynamic content, not the formatting characters
	message := fmt.Sprintf(
		"🏠 *Thông tin IP của máy*\n\n"+
			"🔗 *Local IP:* `%s`\n"+
			"🌐 *WAN IP:* `%s`\n\n"+
			"_Thời gian cập nhật:_ %s",
		util.EscapeMarkdownV2(ipInfo.LocalIP),                            // Escape only the IP address
		util.EscapeMarkdownV2(ipInfo.PublicIP),                           // Escape only the IP address
		util.EscapeMarkdownV2(time.Now().Format("02 Jan 2006 15:04:05")), // Escape only the timestamp
	)

	// Send the message via Telegram bot
	botMsg := &entity.BotMessage{
		ChatID:      chatID,
		Text:        message,
		MessageType: entity.MessageTypeInfo,
		ParseMode:   "MarkdownV2",
		SentAt:      time.Now().Unix(),
	}

	if err := uc.telegramBot.SendBotMessage(ctx, botMsg); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	uc.logger.Info("Successfully handled /home_ip command for chat ID", "chat_id", chatID)
	return nil
}
