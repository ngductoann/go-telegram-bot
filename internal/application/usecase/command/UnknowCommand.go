package usecase

import (
	"context"
	"fmt"

	"go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/domain/types"
	"go-telegram-bot/internal/shared/util"
)

// UnknownHandler handles unknown commands
func UnknownHandler(
	ctx context.Context,
	chatID types.TelegramChatID,
	bot service.TelegramBotService,
) (*types.SendMessageResponse, error) {
	message := fmt.Sprintf("❓ *%s*\n\n"+
		"🤔 %s\n\n"+
		"📝 *%s:*\n"+
		"• `%s` \\- Bắt đầu sử dụng bot\n"+
		"• `%s` \\- Xem thông tin IP\n"+
		"• `%s` \\- Hiển thị hướng dẫn\n\n"+
		"💡 %s",
		util.EscapeMarkdownV2("Lệnh không hợp lệ"),
		util.EscapeMarkdownV2("Tôi không hiểu lệnh này. Vui lòng sử dụng các lệnh sau:"),
		util.EscapeMarkdownV2("Các lệnh có sẵn"),
		util.EscapeMarkdownV2("/start"),
		util.EscapeMarkdownV2("/home_ip"),
		util.EscapeMarkdownV2("/help"),
		util.EscapeMarkdownV2("Hoặc gõ /help để xem hướng dẫn chi tiết."),
	)

	parseMode := types.ParseModeMarkdownV2
	response, err := bot.SendMessageWithResponse(ctx, &types.SendMessageRequest{
		ChatID:    chatID,
		Text:      message,
		ParseMode: &parseMode,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send unknown command message: %w", err)
	}

	return response, nil
}
