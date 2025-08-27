package container

import (
	"fmt"

	domainService "github.com/ngductoann/go-telegram-bot/internal/domain/service"
	"github.com/ngductoann/go-telegram-bot/internal/infrastructure/config"
	"github.com/ngductoann/go-telegram-bot/internal/infrastructure/database"
	"github.com/ngductoann/go-telegram-bot/internal/shared/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Container struct {
	// Config
	Config *config.Config

	// Logger
	Logger domainService.Logger

	// Database
	DB *gorm.DB

	// Redis
	RedisClient *redis.Client
}

func NewContainer() (*Container, error) {
	container := &Container{}

	// Load configuration
	if err := container.initConfig(); err != nil {
		return nil, err
	}

	// Initialize logger
	if err := container.initLogger(); err != nil {
		return nil, err
	}

	// Initialize Postgres database
	if err := container.initDatabase(); err != nil {
		return nil, err
	}

	// Initialize Redis (optional)
	if err := container.initRedis(); err != nil {
		container.Logger.Warn("Redis connection failed, continuing without Redis: " + err.Error())
	}

	return container, nil
}

// initConfig initializes configuration
func (c *Container) initConfig() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	c.Config = cfg
	return nil
}

// initLogger initializes logger
func (c *Container) initLogger() error {
	log, err := logger.NewZapLogger(c.Config)
	if err != nil {
		return err
	}

	c.Logger = log
	return nil
}

// initDatabase initializes database connection
func (c *Container) initDatabase() error {
	db, err := database.NewPostgresConnection(c.Config.GetDatabaseURL())
	if err != nil {
		return err
	}
	database.SetPool(c.Config, db, c.Logger)
	c.DB = db

	return nil
}

// initRedis initializes Redis connection
func (c *Container) initRedis() error {
	redisClient, err := database.NewRedisConnection(
		c.Config.GetRedisURL(),
		c.Config.Redis.Host,
		fmt.Sprintf("%v", c.Config.Redis.Port),
		c.Config.Redis.Password,
		c.Config.Redis.DB,
		c.Logger,
	)
	if err != nil {
		c.Logger.Warn("Faled to connect to Redis " + err.Error())
		return err
	}

	c.RedisClient = redisClient
	return nil
}
