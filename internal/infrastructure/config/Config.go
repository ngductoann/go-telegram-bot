package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	ConfigPath = "./configs/"
	ConfigType = "yaml"
)

type Config struct {
	App      App          `mapstructure:"app"`
	Logger   Logger       `mapstructure:"logger"`
	Postgres Postgres     `mapstructure:"postgres"`
	Client   ClientConfig `mapstructure:"client"`
}

type App struct {
	Environment string `mapstructure:"environment" env:"ENVIRONMENT"`
}

type Logger struct {
	LogLevel   string `mapstructure:"log_level" env:"LOG_LEVEL"`
	FilePath   string `mapstructure:"file_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type Postgres struct {
	// Connection settings
	Host     string `mapstructure:"host" env:"POSTGRES_HOST"`
	Port     int    `mapstructure:"port" env:"POSTGRES_PORT"`
	User     string `mapstructure:"user" env:"POSTGRES_USER"`
	Password string `mapstructure:"password" env:"POSTGRES_PASSWORD"`
	Name     string `mapstructure:"name" env:"POSTGRES_DB"`
	SSLMode  string `mapstructure:"ssl_mode" env:"POSTGRES_SSLMODE"`

	// Pool settings
	MaxOpenConns    int           `mapstructure:"max_open_conns" env:"POSTGRES_MAX_OPEN_CONNS"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" env:"POSTGRES_MAX_IDLE_CONNS"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" env:"POSTGRES_CONN_MAX_LIFETIME"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time" env:"POSTGRES_CONN_MAX_IDLE_TIME"`
}

// ClientConfig holds configuration for the EnhancedTelegramBotService client
type ClientConfig struct {
	Token          string        `json:"token" env:"TELEGRAM_BOT_TOKEN"`
	BaseURL        string        `json:"base_url,omitempty" env:"TELEGRAM_API_BASE_URL"`
	Timeout        time.Duration `json:"timeout" env:"TELEGRAM_API_TIMEOUT"`
	MaxRetries     int           `json:"max_retries" env:"TELEGRAM_API_MAX_RETRIES"`
	RetryDelay     time.Duration `json:"retry_delay" env:"TELEGRAM_API_RETRY_DELAY"`
	RateLimitDelay time.Duration `json:"rate_limit_delay" env:"TELEGRAM_API_RATE_LIMIT_DELAY"`
	EnableMetrics  bool          `json:"enable_metrics" env:"TELEGRAM_API_ENABLE_METRICS"`
	EnableLogging  bool          `json:"enable_logging" env:"TELEGRAM_API_ENABLE_LOGGING"`
	UserAgent      string        `json:"user_agent,omitempty" env:"TELEGRAM_API_USER_AGENT"`
}

func getFileConfig(env string) string {
	switch env {
	case "production":
		return "config_prod"
	case "development":
		return "config_dev"
	case "testing":
		return "config_testing"
	default:
		return "config"
	}
}

// LoadConfig loads the configuration from file and environment variables
func LoadConfig() (*Config, error) {
	_ = godotenv.Load() // Load .env file if exists, ignore error if not found
	env := os.Getenv("ENVIRONMENT")

	// Create a new Viper instance
	v := viper.New()
	v.SetConfigName(getFileConfig(env))
	v.SetConfigType(ConfigType)
	v.AddConfigPath(ConfigPath)

	// Enable automatic environment variable lookup
	v.AutomaticEnv()
	bindEnvironmentVariables(v)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	config := &Config{} // Initialize an empty Config struct
	if err := v.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	config.Client.BaseURL = config.Client.BaseURL + config.Client.Token

	// Print the loaded configuration
	fmt.Println("=== Loaded Configuration ===")
	fmt.Printf("%+v\n", config)
	fmt.Println("============================")

	return config, nil
}

// bindEnvironmentVariables binds specific environment variables to configuration keys
func bindEnvironmentVariables(v *viper.Viper) {
	// Bind app environment variables to config keys
	v.BindEnv("app.environment", "ENVIRONMENT")

	// Database configuration
	v.BindEnv("postgres.host", "POSTGRES_HOST")
	v.BindEnv("postgres.port", "POSTGRES_PORT")
	v.BindEnv("postgres.user", "POSTGRES_USER")
	v.BindEnv("postgres.password", "POSTGRES_PASSWORD")
	v.BindEnv("postgres.name", "POSTGRES_NAME")
	v.BindEnv("postgres.ssl_mode", "POSTGRES_SSL_MODE")
	v.BindEnv("postgres.max_open_conns", "POSTGRES_MAX_OPEN_CONNS")
	v.BindEnv("postgres.max_idle_conns", "POSTGRES_MAX_IDLE_CONNS")
	v.BindEnv("postgres.conn_max_lifetime", "POSTGRES_CONN_MAX_LIFETIME")
	v.BindEnv("postgres.conn_max_idle_time", "POSTGRES_CONN_MAX_IDLE_TIME")

	// Client configuration
	v.BindEnv("client.token", "TELEGRAM_BOT_TOKEN")
	v.BindEnv("client.base_url", "TELEGRAM_API_BASE_URL")
	v.BindEnv("client.timeout", "TELEGRAM_API_TIMEOUT")
	v.BindEnv("client.max_retries", "TELEGRAM_API_MAX_RETRIES")
	v.BindEnv("client.retry_delay", "TELEGRAM_API_RETRY_DELAY")
	v.BindEnv("client.rate_limit_delay", "TELEGRAM_API_RATE_LIMIT_DELAY")
	v.BindEnv("client.enable_metrics", "TELEGRAM_API_ENABLE_METRICS")
	v.BindEnv("client.enable_logging", "TELEGRAM_API_ENABLE_LOGGING")
	v.BindEnv("client.user_agent", "TELEGRAM_API_USER_AGENT")
}

// GetDatabaseURL constructs the database connection URL from the configuration.
func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Name,
		c.Postgres.SSLMode,
	)
}
