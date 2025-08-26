package database

import (
	domainService "github.com/ngductoann/go-telegram-bot/internal/domain/service"
	"github.com/ngductoann/go-telegram-bot/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresConnection creates a new PostgreSQL database connection
func NewPostgresConnection(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SetPool(cfg *config.Config, database *gorm.DB, logger domainService.Logger) {
	sqlDB, err := database.DB()
	logger.CheckErrorPanic(err, "Failed to get database instance")

	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.Database.ConnMaxIdleTime)
}
