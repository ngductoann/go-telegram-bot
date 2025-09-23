package service

import (
	"context"
	"go-telegram-bot/internal/application/usecase"
	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/service"
	"strings"
)

// BotApplicationService orchestrates the use cases for the bot application
type BotApplicationService struct {
	homeIPUseCase         *usecase.HomeIPUseCase
	startUseCase          *usecase.StartUseCase
	helpUseCase           *usecase.HelpUseCase
	unknownCommandUseCase *usecase.UnknownCommandUseCase
	telegramBotService    service.TelegramBotService
	logger                service.Logger
}

// NewBotApplicationService creates a new instance of BotApplicationService
func NewBotApplicationService(ipService service.IPService, telegramBot service.TelegramBotService, logger service.Logger) *BotApplicationService {
	return &BotApplicationService{
		homeIPUseCase:         usecase.NewHomeIPUseCase(ipService, telegramBot, logger),
		startUseCase:          usecase.NewStartUseCase(telegramBot, logger),
		helpUseCase:           usecase.NewHelpUseCase(telegramBot, logger),
		unknownCommandUseCase: usecase.NewUnknownCommandUseCase(telegramBot, logger),
		telegramBotService:    telegramBot,
		logger:                logger,
	}
}

// HandleHomeIPCommand processes the /home_ip command
func (s *BotApplicationService) HandleHomeIPCommand(
	ctx context.Context,
	chatID int64,
) error {
	return s.homeIPUseCase.Execute(ctx, chatID)
}

// HandleStartCommand processes the /start command
func (s *BotApplicationService) HandleStartCommand(
	ctx context.Context,
	chatID int64,
) error {
	return s.startUseCase.Execute(ctx, chatID)
}

// HandleHelpCommand processes the /help command
func (s *BotApplicationService) HandleHelpCommand(
	ctx context.Context,
	chatID int64,
) error {
	return s.helpUseCase.Execute(ctx, chatID)
}

// HandleUnknownCommand processes unknown commands
func (s *BotApplicationService) HandleUnknownCommand(
	ctx context.Context,
	chatID int64,
) error {
	return s.unknownCommandUseCase.Execute(ctx, chatID)
}

func (s *BotApplicationService) ProcessUpdate(ctx context.Context, update entity.TelegramUpdate) error {
	// check if the update contains a message
	if update.Message == nil {
		return nil // Ignore non-message updates
	}

	message := update.Message
	chatID := message.Chat.ID
	text := message.Content

	s.logger.Info("Received message from chat ID %d: %s", chatID, text)

	switch text {
	case "/home_ip":
		return s.HandleHomeIPCommand(ctx, chatID)
	case "/start":
		return s.HandleStartCommand(ctx, chatID)
	case "/help":
		return s.HandleHelpCommand(ctx, chatID)
	default:
		if strings.HasPrefix(text, "/") {
			return s.HandleUnknownCommand(ctx, chatID)
		}
		return nil
	}
}
