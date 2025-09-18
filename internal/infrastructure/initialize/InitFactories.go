package initialize

import "go-telegram-bot/internal/infrastructure/factory"

func (c *Container) InitFactories() {
	c.ApplicationFactory = factory.NewApplicationServiceFactory()
	c.PresentationFactory = factory.NewPresentationFactory()
}
