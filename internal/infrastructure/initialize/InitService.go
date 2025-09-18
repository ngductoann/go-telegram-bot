package initialize

import "go-telegram-bot/internal/infrastructure/service"

func (c *Container) InitService() {
	c.IpService = service.NewIPService()
	c.TelegramBot = service.NewTelegramBot(c.Config.App.TelegramBotToken)
}
