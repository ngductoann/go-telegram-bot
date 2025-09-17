package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN is required")
	}

	return &Config{
		TelegramBotToken: botToken,
	}, nil
}

func MustLoadConfig() *Config {
	config, err := LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	return config
}
