package initialize

func (c *Container) InitPresentationLayer() {
	c.TelegramHandler = c.PresentationFactory.CreateTelegramHandler(c.BotApplicationService, c.TelegramBot, c.Logger)
}
