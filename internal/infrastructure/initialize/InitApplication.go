package initialize

func (c *Container) InitApplicationServices() {
	c.BotApplicationService = c.ApplicationFactory.CreateBotApplicationService(c.IpService, c.TelegramBot, c.Logger)
}
