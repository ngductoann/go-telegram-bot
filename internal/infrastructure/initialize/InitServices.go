package initialize

import "go-telegram-bot/internal/infrastructure/service"

func (c *Container) InitServices() {
	c.IPService = service.NewIPService()
	c.TelegramBot = service.NewTelegramBot(
		c.Config.Client, nil, c.Logger,
	)
}
