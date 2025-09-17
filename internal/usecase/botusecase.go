package usecase

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go-telegram-bot/internal/domain/entity"
	domainService "go-telegram-bot/internal/domain/service"
)

type botUseCase struct {
	ipService   domainService.IPService
	telegramBot domainService.TelegramBotService
}

// NewBotUseCase creates a new instance of BotUseCase
func NewBotUseCase(ipService domainService.IPService, telegramBot domainService.TelegramBotService) domainService.BotUseCase {
	return &botUseCase{
		ipService:   ipService,
		telegramBot: telegramBot,
	}
}

// HandleHomeIPCommand processes the /home_ip command
func (u *botUseCase) HandleHomeIPCommand(ctx context.Context, chatID int64) error {
	log.Printf("Handling /home_ip command for chat ID: %d", chatID)

	// Retrieve IP information
	ipInfo, err := u.ipService.GetIPInfo(ctx)
	if err != nil {
		log.Printf("Error retrieving IP info: %v", err)

		// Send error message to user
		errorMsg := "❌Failed to retrieve IP information. Please try again later."
		if sendErr := u.telegramBot.SendMessage(ctx, chatID, errorMsg); sendErr != nil {
			return fmt.Errorf("failed to send error message: %w", sendErr)
		}
		return fmt.Errorf("failed to get IP info: %w", err)
	}

	// Create response message
	// Tạo tin nhắn phản hồi
	message := fmt.Sprintf(`🏠 **Thông tin IP của máy:**

🔗 **Local IP:** %s
🌐 **WAN IP:** %s

*Cập nhật lúc: hiện tại*`, ipInfo.LocalIP, ipInfo.PublicIP)

	// Send message to user
	if err := u.telegramBot.SendMessage(ctx, chatID, message); err != nil {
		log.Printf("Error sending message: %v", err)
		return fmt.Errorf("failed to send message: %w", err)
	}

	log.Printf("Successfully handled /home_ip command for chat ID: %d", chatID)
	return nil
}

// ProcessUpdate processes incoming updates from Telegram
func (u *botUseCase) ProcessUpdate(ctx context.Context, update entity.TelegramUpdate) error {
	// Check if the update contains a message
	if update.Message == nil {
		return nil // Ignore non-message updates
	}

	message := update.Message
	chatID := message.Chat.ID
	text := message.Text

	log.Printf("Received message from chat ID %d: %s", chatID, text)

	switch text {
	case "/start":
		welcomeMsg := `👋 Chào mừng bạn đến với Bot IP!

🤖 **Các lệnh có sẵn:**
• /home_ip - Xem thông tin IP local và WAN của máy
• /help - Hiển thị hướng dẫn sử dụng

💡 **Sử dụng:** Chỉ cần gửi lệnh /home_ip để xem thông tin IP hiện tại.`

		return u.telegramBot.SendMessage(ctx, chatID, welcomeMsg)
	case "/help":
		helpMsg := `📖 **Hướng dẫn sử dụng Bot IP:**

🔹 **/home_ip** - Lấy thông tin IP hiện tại:
   • Local IP: Địa chỉ IP trong mạng nội bộ
   • WAN IP: Địa chỉ IP công khai trên Internet

🔹 **/start** - Hiển thị lời chào và các lệnh cơ bản
🔹 **/help** - Hiển thị hướng dẫn này

ℹ️ **Lưu ý:** Bot sẽ tự động cập nhật thông tin IP mỗi khi bạn gọi lệnh.`

		return u.telegramBot.SendMessage(ctx, chatID, helpMsg)
	case "/home_ip":
		return u.HandleHomeIPCommand(ctx, chatID)
	default:
		if strings.HasPrefix(text, "/") {
			unknownCmdMsg := `❓ Lệnh không xác định.

Vui lòng sử dụng /help để xem danh sách lệnh hợp lệ.`
			return u.telegramBot.SendMessage(ctx, chatID, unknownCmdMsg)
		}

		return nil // Ignore non-command messages
	}
}
