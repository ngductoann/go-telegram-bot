package service

import (
	"context"

	"go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/domain/types"
)

// BotApplicationService is the presentation layer adapter for bot operations.
type BotApplicationService struct {
	botUseCase service.BotUseCase
	logger     service.Logger
}

// NewBotApplicationService creates a new instance of BotApplicationService.
func NewBotApplicationService(
	botUseCase service.BotUseCase,
	logger service.Logger,
) *BotApplicationService {
	return &BotApplicationService{
		botUseCase: botUseCase,
		logger:     logger,
	}
}

// HandleHomeIPCommand processes the /home_ip command.
func (s *BotApplicationService) HandleHomeIPCommand(
	ctx context.Context,
	chatID types.TelegramChatID,
) (*types.SendMessageResponse, error) {
	return s.botUseCase.HandleHomeIPCommand(ctx, chatID)
}

// HandleStartCommand processes the /start command.
func (s *BotApplicationService) HandleStartCommand(
	ctx context.Context,
	chatID types.TelegramChatID,
) (*types.SendMessageResponse, error) {
	return s.botUseCase.HandleStartCommand(ctx, chatID)
}

// HandleHelpCommand processes the /help command.
func (s *BotApplicationService) HandleHelpCommand(
	ctx context.Context,
	chatID types.TelegramChatID,
) (*types.SendMessageResponse, error) {
	return s.botUseCase.HandleHelpCommand(ctx, chatID)
}

// HandleUnknownCommand processes unknown commands.
func (s *BotApplicationService) HandleUnknownCommand(
	ctx context.Context,
	chatID types.TelegramChatID,
) (*types.SendMessageResponse, error) {
	return s.botUseCase.HandleUnknownCommand(ctx, chatID)
}

// ProcessUpdate processes an incoming Telegram update.
func (s *BotApplicationService) ProcessUpdate(ctx context.Context, update types.TelegramUpdate) error {
	return s.botUseCase.ProcessUpdate(ctx, update)
}

// ValidateUpdate validates the incoming Telegram update.
func (s *BotApplicationService) ValidateUpdate(update types.TelegramUpdate) error {
	return s.botUseCase.ValidateUpdate(update)
}
