package initialize

import "go-telegram-bot/internal/infrastructure/config"

func (c *Container) LoadConfig() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	c.Config = cfg
	return nil
}
