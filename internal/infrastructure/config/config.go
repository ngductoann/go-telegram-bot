package config

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	// Config file settings
	ConfigPath = "./configs/"
	ConfigName = "config"
	ConfigType = "yaml"
)

type Config struct {
	App      App      `mapstructure:"app"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
	Log      Log      `mapstructure:"log"`
}

type App struct {
	BotToken    string `mapstructure:"bot_token" env:"BOT_TOKEN"`
	Environment string `mapstructure:"environment" env:"ENVIRONMENT"`
}

type Database struct {
	// Connection settings
	Host     string `mapstructure:"host" env:"DB_HOST"`
	Port     int    `mapstructure:"port" env:"DB_PORT"`
	User     string `mapstructure:"user" env:"DB_USER"`
	Password string `mapstructure:"password" env:"DB_PASSWORD"`
	Name     string `mapstructure:"name" env:"DB_NAME"`
	SSLMode  string `mapstructure:"ssl_mode" env:"DB_SSL_MODE"`

	// Pool settings
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}

type Redis struct {
	Host     string `mapstructure:"host" env:"REDIS_HOST"`
	Port     string `mapstructure:"port" env:"REDIS_PORT"`
	Password string `mapstructure:"password" env:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"db" env:"REDIS_DB"`
}

type Log struct {
	LogLevel   string `mapstructure:"log_level" env:"LOG_LEVEL"`
	Filepath   string `mapstructure:"filepath"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

// LoadConfig loads configuration from environment variables and YAML file
func LoadConfig() (*Config, error) {
	// Load .env file (optional)
	_ = godotenv.Load()

	// Create viper instance
	v := viper.New()

	// Configure viper
	v.AddConfigPath(ConfigPath)
	v.SetConfigName(ConfigName)
	v.SetConfigType(ConfigType)

	// Enable automatic environment variable lookup
	v.AutomaticEnv()

	// Bind specific environment variables to config keys
	bindEnvironmentVariables(v)

	// Read config file - the YAML file contains default values
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal to struct
	config := &Config{}
	if err := v.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate critical settings
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return config, nil
}

// bindEnvironmentVariables binds environment variables to viper keys
func bindEnvironmentVariables(v *viper.Viper) {
	// App configuration
	v.BindEnv("app.bot_token", "BOT_TOKEN")
	v.BindEnv("app.environment", "ENVIRONMENT")

	// Database configuration
	v.BindEnv("database.host", "DB_HOST")
	v.BindEnv("database.port", "DB_PORT")
	v.BindEnv("database.user", "DB_USER")
	v.BindEnv("database.password", "DB_PASSWORD")
	v.BindEnv("database.name", "DB_NAME")
	v.BindEnv("database.ssl_mode", "DB_SSL_MODE")

	// Redis configuration
	v.BindEnv("redis.host", "REDIS_HOST")
	v.BindEnv("redis.port", "REDIS_PORT")
	v.BindEnv("redis.password", "REDIS_PASSWORD")
	v.BindEnv("redis.db", "REDIS_DB")

	// Log configuration
	v.BindEnv("log.log_level", "LOG_LEVEL")
}

// validateConfig validates critical configuration values
func validateConfig(config *Config) error {
	// BOT_TOKEN is required in production
	if config.App.Environment == "production" && config.App.BotToken == "" {
		return fmt.Errorf("BOT_TOKEN is required in production environment")
	}

	if config.Database.Host == "" {
		return fmt.Errorf("database host is required")
	}

	if config.Database.Port <= 0 {
		return fmt.Errorf("database port must be positive")
	}

	return nil
}

// GetBotToken returns the bot token
func (c *Config) GetBotToken() string {
	return c.App.BotToken
}

// GetEnvironment returns the environment
func (c *Config) GetEnvironment() string {
	return c.App.Environment
}

// IsDevelopment checks if the environment is development
func (c *Config) IsDevelopment() bool {
	return c.App.Environment == "development"
}

// IsProduction checks if the environment is production
func (c *Config) IsProduction() bool {
	return c.App.Environment == "production"
}

// GetDatabaseURL returns formatted database connection URL
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

// GetRedisURL returns formatted Redis connection URL
func (c *Config) GetRedisURL() string {
	if c.Redis.Password != "" {
		return fmt.Sprintf(
			"redis://:%s@%s:%s/%d",
			c.Redis.Password,
			c.Redis.Host,
			c.Redis.Port,
			c.Redis.DB,
		)
	}
	return fmt.Sprintf(
		"redis://%s:%s/%d",
		c.Redis.Host,
		c.Redis.Port,
		c.Redis.DB,
	)
}

// GetLogLevel returns the log level
func (c *Config) GetLogLevel() string {
	return c.Log.LogLevel
}

// GetLogFilePath returns the log file path
func (c *Config) GetLogFilePath() string {
	return c.Log.Filepath
}
