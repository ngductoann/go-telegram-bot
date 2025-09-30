package usecase

import (
	"context"
	"fmt"

	"go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/domain/types"
	"go-telegram-bot/internal/shared/util"
)

// HelpHandler handles the /help command
func HelpHandler(
	ctx context.Context,
	chatID types.TelegramChatID,
	bot service.TelegramBotService,
) (*types.SendMessageResponse, error) {
	message := fmt.Sprintf("📖 *%s*\n\n"+
		"🤖 *%s*\n\n"+
		"📝 *%s:*\n"+
		"• `%s` \\- Bắt đầu sử dụng bot\n"+
		"• `%s` \\- Xem thông tin IP local và WAN của máy\n"+
		"• `%s` \\- Hiển thị hướng dẫn này\n\n"+
		"💡 *%s:*\n"+
		"Bot sẽ hiển thị thông tin IP địa phương và WAN của máy bạn cùng với thông tin địa lý\\.\n\n"+
		"🔧 *%s*",
		util.EscapeMarkdownV2("Hướng dẫn sử dụng Bot IP"),
		util.EscapeMarkdownV2("Đây là bot hỗ trợ kiểm tra thông tin IP của bạn."),
		util.EscapeMarkdownV2("Các lệnh có sẵn"),
		util.EscapeMarkdownV2("/start"),
		util.EscapeMarkdownV2("/home_ip"),
		util.EscapeMarkdownV2("/help"),
		util.EscapeMarkdownV2("Mô tả"),
		util.EscapeMarkdownV2("Cần hỗ trợ? Liên hệ quản trị viên."),
	)

	parseMode := types.ParseModeMarkdownV2
	response, err := bot.SendMessageWithResponse(ctx, &types.SendMessageRequest{
		ChatID:    chatID,
		Text:      message,
		ParseMode: &parseMode,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send help message: %w", err)
	}

	return response, nil
}
