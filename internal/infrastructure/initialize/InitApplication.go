package initialize

import "go-telegram-bot/internal/application/service"

func (c *Container) InitApplicationServices() {
	// Create BotUseCase implementation
	c.BotUseCase = service.NewBotUseCaseImpl(
		c.IPService,
		c.TelegramBot,
		c.Logger,
	)
}
