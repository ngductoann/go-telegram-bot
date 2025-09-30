package usecase

import (
	"context"
	"fmt"
	"time"

	"go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/domain/types"
	"go-telegram-bot/internal/shared/util"
)

// HomeIPHandler handles the /home_ip command
func HomeIPHandler(
	ctx context.Context,
	chatID types.TelegramChatID,
	ipService service.IPService,
	logger service.Logger,
	bot service.TelegramBotService,
) (*types.SendMessageResponse, error) {
	// Retrieve IP information from the IP service
	ipInfo, err := ipService.GetIPInfo(ctx)
	if err != nil {
		logger.Error("Failed to get IP info", "error", err)

		errorMsg := "❌ Failed to retrieve IP information. Please try again later."
		response, err := bot.SendMessageWithResponse(ctx, &types.SendMessageRequest{
			ChatID: chatID,
			Text:   errorMsg,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to send error message: %w", err)
		}

		return response, fmt.Errorf("Failed to send error message to Telegram", "error", err)
	}

	// Create response message with proper MarkdownV2 formatting
	message := fmt.Sprintf(
		"🏠 *%s*\n\n"+
			"🔗 *Local IP:* `%s`\n"+
			"🌍 *WAN IP:* `%s`\n\n"+
			"⏰ *%s*",
		util.EscapeMarkdownV2("Thông tin IP hiện tại"),
		util.EscapeMarkdownV2(ipInfo.LocalIP),
		util.EscapeMarkdownV2(ipInfo.PublicIP),
		util.EscapeMarkdownV2(
			fmt.Sprintf(
				"Cập nhật: %s",
				time.Now().Format("02/01/2006 15:04:05"),
			),
		),
	)

	parseMode := types.ParseModeMarkdownV2
	// telegramMessageID, err := bot.SendMessage(ctx, chatID, message, &parseMode)
	response, err := bot.SendMessageWithResponse(ctx, &types.SendMessageRequest{
		ChatID:    chatID,
		Text:      message,
		ParseMode: &parseMode,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send IP info message: %w", err)
	}

	return response, nil
}
