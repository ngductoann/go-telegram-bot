package factory

import (
	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/presentation"
	"go-telegram-bot/internal/presentation/service"
)

// PresentationFactory creates presentation layer components
type PresentationFactory struct{}

// NewPresentationFactory creates a new presentation factory
func NewPresentationFactory() *PresentationFactory {
	return &PresentationFactory{}
}

// CreateTelegramHandler creates a configured TelegramHandler
func (f *PresentationFactory) CreateTelegramHandler(
	botAppService *service.BotApplicationService,
	telegramBot domainService.TelegramBotService,
	logger domainService.Logger,
) *presentation.TelegramHandler {
	return presentation.NewTelegramHandler(botAppService, telegramBot, logger)
}

// CreatePresentationService creates a BotApplicationService
func (f *PresentationFactory) CreatePresentationService(
	botUseCase domainService.BotUseCase,
	logger domainService.Logger,
) *service.BotApplicationService {
	return service.NewBotApplicationService(botUseCase, logger)
}
