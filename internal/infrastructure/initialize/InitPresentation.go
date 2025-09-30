package initialize

func (c *Container) InitPresentationLayer() {
	c.BotApplicationService = c.PresentationFactory.CreatePresentationService(
		c.BotUseCase,
		c.Logger,
	)
	c.TelegramHandler = c.PresentationFactory.CreateTelegramHandler(
		c.BotApplicationService,
		c.TelegramBot,
		c.Logger,
	)
}
