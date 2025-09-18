package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	ConfigPath = "./configs/"
	ConfigType = "yaml"
)

type Config struct {
	App    App    `mapstructure:"app"`
	Logger Logger `mapstructure:"logger"`
}

type App struct {
	TelegramBotToken string `mapstructure:"telegram_bot_token" env:"TELEGRAM_BOT_TOKEN"`
	Environment      string `mapstructure:"environment" env:"ENVIRONMENT"`
}

type Logger struct {
	LogLevel   string `mapstructure:"log_level" env:"LOG_LEVEL"`
	FilePath   string `mapstructure:"file_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load() // Load .env file if exists, ignore error if not found
	env := os.Getenv("ENVIRONMENT")
	var configName string
	switch env {
	case "production":
		configName = "config_prod"
	case "development":
		configName = "config_dev"
	case "testing":
		configName = "config_testing"
	default:
		configName = "config"
	}

	// Create a new Viper instance
	v := viper.New()
	v.SetConfigName(configName)
	v.SetConfigType(ConfigType)
	v.AddConfigPath(ConfigPath)
	bindEnvironmentVariables(v)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	config := &Config{} // Initialize an empty Config struct
	if err := v.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

// bindEnvironmentVariables binds specific environment variables to configuration keys
func bindEnvironmentVariables(v *viper.Viper) {
	// Bind app environment variables to config keys
	v.BindEnv("app.telegram_bot_token", "TELEGRAM_BOT_TOKEN")
	v.BindEnv("app.environment", "ENVIRONMENT")
}

// validateConfig checks for required configuration values
func validateConfig(config *Config) error {
	if config.App.Environment == "production" && config.App.TelegramBotToken == "" {
		return errors.New("TELEGRAM_BOT_TOKEN must be set in production environment")
	}
	return nil
}
