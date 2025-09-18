package initialize

import (
	"go-telegram-bot/internal/application/service"
	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/infrastructure/config"
	"go-telegram-bot/internal/infrastructure/factory"
	"go-telegram-bot/internal/presentation"
)

type Container struct {
	Config *config.Config
	Logger domainService.Logger

	// Services
	IpService   domainService.IPService
	TelegramBot domainService.TelegramBotService

	// Factories
	ApplicationFactory  *factory.ApplicationServiceFactory
	PresentationFactory *factory.PresentationFactory

	// Application Services
	BotApplicationService *service.BotApplicationService

	// Presentation Layer
	TelegramHandler *presentation.TelegramHandler
}

func NewContainer() (*Container, error) {
	container := &Container{}

	// load config
	if err := container.LoadConfig(); err != nil {
		return nil, err
	}

	// init logger
	if err := container.InitLogger(); err != nil {
		return nil, err
	}

	// init services
	container.InitService()

	// init factories
	container.InitFactories()

	// init application services
	container.InitApplicationServices()

	// init presentation layer
	container.InitPresentationLayer()

	return container, nil
}
