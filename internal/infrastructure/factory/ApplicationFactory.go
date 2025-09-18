package factory

import (
	"go-telegram-bot/internal/application/service"
	domainService "go-telegram-bot/internal/domain/service"
)

// ApplicationServiceFactory creates application layer services
type ApplicationServiceFactory struct{}

// NewApplicationServiceFactory creates a new application service factory
func NewApplicationServiceFactory() *ApplicationServiceFactory {
	return &ApplicationServiceFactory{}
}

// CreateBotApplicationService creates a configured BotApplicationService
func (f *ApplicationServiceFactory) CreateBotApplicationService(
	ipService domainService.IPService,
	telegramBot domainService.TelegramBotService,
	logger domainService.Logger,
) *service.BotApplicationService {
	return service.NewBotApplicationService(ipService, telegramBot, logger)
}
