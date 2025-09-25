package initialize

import (
	"go-telegram-bot/internal/application/service"
	"go-telegram-bot/internal/domain/repository"
	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/infrastructure/config"
	"go-telegram-bot/internal/infrastructure/factory"
	"go-telegram-bot/internal/presentation"

	"gorm.io/gorm"
)

type Container struct {
	Config *config.Config
	Logger domainService.Logger
	DB     *gorm.DB

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

	// Repositories
	UserRepo        repository.UserRepository
	UserProfileRepo repository.UserProfileRepository
	ChatRepo        repository.ChatRepository
	MessageRepo     repository.MessageRepository
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

	// init database connection
	container.initDatabase()

	// init repositories
	container.InitRepositories()

	return container, nil
}
