package config

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App      App      `mapstructure:"app"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
	Log      Log      `mapstructure:"log"`
}

type App struct {
	BotToken    string `mapstructure:"bot_token"`
	Environment string `mapstructure:"environment"`
}

type Database struct {
	// database connection
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl_mode"`

	// pool settings
	MaxIdleConns    int           `mapstructure:"maxIdleConns"`
	MaxOpenConns    int           `mapstructure:"maxOpenConns"`
	ConnMaxLifetime time.Duration `mapstructure:"connMaxLifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"connMaxIdleTime"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Log struct {
	LogLevel   string `mapstructure:"log_level"`
	Filepath   string `mapstructure:"filepath"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

// LoadConfig loads configuration from Environment variables and yaml Config
func LoadConfig() (*Config, error) {
	_ = godotenv.Load() // Ignore error if .env not found

	v := viper.New()
	v.AddConfigPath("./configs/")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// GetBotToken returns the bot token
func (c *Config) GetBotToken() string {
	return c.App.BotToken
}

// GetDatabaseURL returns Database URL
func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

// GetRedisURL returns the redis URL
func (c *Config) GetRedisURL() string {
	if c.Redis.Password != "" {
		return fmt.Sprintf(
			"redis://:%s@%s:%d",
			c.Redis.Password,
			c.Redis.Host,
			c.Redis.Port,
		)
	}
	return fmt.Sprintf(
		"redis://%s:%d",
		c.Redis.Host,
		c.Redis.Port,
	)
}

// GetLogLevel returns the log level
func (c *Config) GetLogLevel() string {
	return c.Log.LogLevel
}
