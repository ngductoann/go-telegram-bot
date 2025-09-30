package service

import (
	"context"
	"fmt"
	"strings"

	usecase "go-telegram-bot/internal/application/usecase/command"
	"go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/domain/types"
)

// BotUseCaseImpl implements BotUseCase interface
type BotUseCaseImpl struct {
	ipService   service.IPService
	telegramBot service.TelegramBotService
	logger      service.Logger
}

// NewBotUseCaseImpl creates a new instance of BotUseCaseImpl
func NewBotUseCaseImpl(
	ipService service.IPService,
	telegramBot service.TelegramBotService,
	logger service.Logger,
) service.BotUseCase {
	return &BotUseCaseImpl{
		ipService:   ipService,
		telegramBot: telegramBot,
		logger:      logger,
	}
}

// HandleHomeIPCommand processes the /home_ip command
func (u *BotUseCaseImpl) HandleHomeIPCommand(
	ctx context.Context, chatID types.TelegramChatID,
) (*types.SendMessageResponse, error) {
	u.logger.Info("Handling /home_ip command", "chat_id", chatID)
	return usecase.HomeIPHandler(ctx, chatID, u.ipService, u.logger, u.telegramBot)
}

// HandleStartCommand processes the /start command
func (u *BotUseCaseImpl) HandleStartCommand(
	ctx context.Context, chatID types.TelegramChatID,
) (*types.SendMessageResponse, error) {
	u.logger.Info("Handling /start command", "chat_id", chatID)
	return usecase.StartHandler(ctx, chatID, u.telegramBot)
}

// HandleHelpCommand processes the /help command
func (u *BotUseCaseImpl) HandleHelpCommand(
	ctx context.Context, chatID types.TelegramChatID,
) (*types.SendMessageResponse, error) {
	u.logger.Info("Handling /help command", "chat_id", chatID)
	return usecase.HelpHandler(ctx, chatID, u.telegramBot)
}

// HandleUnknownCommand processes unknown or invalid commands
func (u *BotUseCaseImpl) HandleUnknownCommand(
	ctx context.Context, chatID types.TelegramChatID,
) (*types.SendMessageResponse, error) {
	u.logger.Info("Handling unknown command", "chat_id", chatID)
	return usecase.UnknownHandler(ctx, chatID, u.telegramBot)
}

// ProcessUpdate processes incoming Telegram updates
func (u *BotUseCaseImpl) ProcessUpdate(
	ctx context.Context, update types.TelegramUpdate,
) error {
	// Check if the update contains a message
	if update.Message == nil {
		return nil // ignore non-message updates
	}

	message := update.Message
	chatID := message.Chat.ID
	text := message.Text

	u.logger.Info("Received message", "chat_id", chatID, "text", text)

	command, _ := u.extractCommand(text)

	var err error = nil
	// var response *types.SendMessageResponse

	switch command {
	case "home_ip":
		_, err = u.HandleHomeIPCommand(ctx, chatID)
	case "start":
		_, err = u.HandleStartCommand(ctx, chatID)
	case "help":
		_, err = u.HandleHelpCommand(ctx, chatID)
	default:
		if command != "" {
			_, err = u.HandleUnknownCommand(ctx, chatID)
			return err
		}
		return nil
	}

	return err
}

// ValidateUpdate validates the structure and content of an update
func (u *BotUseCaseImpl) ValidateUpdate(update types.TelegramUpdate) error {
	if update.UpdateID == 0 {
		return fmt.Errorf("invalid update: missing update ID")
	}

	if update.Message != nil {
		if int64(update.Message.Chat.ID) == 0 {
			return fmt.Errorf("invalid update: missing chat ID")
		}
	}

	return nil
}

// ExtractCommand extracts command from message text
func (u *BotUseCaseImpl) extractCommand(text *string) (command string, args []string) {
	if text == nil || *text == "" {
		return "", nil
	}

	if !strings.HasPrefix(*text, "/") {
		return "", nil
	}

	parts := strings.Fields(*text)
	if len(parts) == 0 {
		return "", nil
	}

	command = parts[0]
	if len(parts) > 1 {
		args = parts[1:]
	}

	return command, args
}
