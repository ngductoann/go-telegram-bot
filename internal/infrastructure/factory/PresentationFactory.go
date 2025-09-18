package factory

import (
	"go-telegram-bot/internal/application/service"
	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/presentation"
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
