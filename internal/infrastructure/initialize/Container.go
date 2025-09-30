package initialize

import (
	"go-telegram-bot/internal/domain/repository"
	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/infrastructure/config"
	"go-telegram-bot/internal/infrastructure/factory"
	"go-telegram-bot/internal/presentation"
	"go-telegram-bot/internal/presentation/service"

	"gorm.io/gorm"
)

type Container struct {
	Config *config.Config
	Logger domainService.Logger
	DB     *gorm.DB

	// Services
	IPService   domainService.IPService
	TelegramBot domainService.TelegramBotService
	BotUseCase  domainService.BotUseCase

	// Repositories
	UserRepo        repository.UserRepository
	UserProfileRepo repository.UserProfileRepository
	ChatRepo        repository.ChatRepository
	MessageRepo     repository.MessageRepository

	// Factories
	ApplicationFactory  *factory.ApplicationServiceFactory
	PresentationFactory *factory.PresentationFactory

	// Application Services

	// Presentation Layer
	TelegramHandler       *presentation.TelegramHandler
	BotApplicationService *service.BotApplicationService
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

	// init database connection
	container.initDatabase()

	// init repositories
	container.InitRepositories()

	// init services
	container.InitServices()

	// init factories
	container.InitFactories()

	// init application services
	container.InitApplicationServices()

	// init presentation layer
	container.InitPresentationLayer()

	return container, nil
}
