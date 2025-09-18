package initialize

import "go-telegram-bot/internal/shared/logger"

func (c *Container) InitLogger() error {
	logger, err := logger.NewZapLogger(c.Config)
	if err != nil {
		return err
	}
	c.Logger = logger
	return nil
}
