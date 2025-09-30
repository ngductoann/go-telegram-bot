package usecase

import (
	"context"
	"fmt"

	"go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/domain/types"
	"go-telegram-bot/internal/shared/util"
)

// StartHandler handles the /start command
func StartHandler(
	ctx context.Context,
	chatID types.TelegramChatID,
	bot service.TelegramBotService,
) (*types.SendMessageResponse, error) {
	message := fmt.Sprintf("*%s*\n\n🤖 *%s*\n"+
		"• `%s` \\- Xem thông tin IP local và WAN của máy\n"+
		"• `%s` \\- Hiển thị hướng dẫn sử dụng\n",
		util.EscapeMarkdownV2("👋 Chào mừng bạn đến với Bot IP!"),
		util.EscapeMarkdownV2("Các lệnh có sẵn:"),
		util.EscapeMarkdownV2("/home_ip"),
		util.EscapeMarkdownV2("/help"),
	)

	parseMode := types.ParseModeMarkdownV2
	response, err := bot.SendMessageWithResponse(ctx, &types.SendMessageRequest{
		ChatID:    chatID,
		Text:      message,
		ParseMode: &parseMode,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send start message: %w", err)
	}

	return response, nil
}
