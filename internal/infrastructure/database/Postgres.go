package database

import (
	"go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/infrastructure/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresConnection creates a new GORM DB connection to a PostgreSQL database.
func NewPostgresConnection(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// SetPool configures the connection pool settings for the given GORM DB instance.
func SetPool(cfg *config.Config, db *gorm.DB, logger service.Logger) {
	logger.Info(
		"Setting up Postgres connection pool: max_open_conns=%d, max_idle_conns=%d, conn_max_lifetime=%s, conn_max_idle_time=%s",
		cfg.Postgres.MaxOpenConns,
		cfg.Postgres.MaxIdleConns,
		cfg.Postgres.ConnMaxLifetime,
		cfg.Postgres.ConnMaxIdleTime,
	)
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("Failed to get sql.DB from gorm.DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.Postgres.ConnMaxIdleTime)
}
