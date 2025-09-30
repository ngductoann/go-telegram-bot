package service

import (
	"context"
	"database/sql"
	"fmt"

	"go-telegram-bot/internal/domain/service"

	"gorm.io/gorm"
)

// TransactionManager handles database transactions.
type TransactionManager struct {
	db     *gorm.DB
	logger service.Logger
}

// NewTransactionManager creates a new instance of TransactionManager.
func NewTransactionManager(db *gorm.DB, logger service.Logger) *TransactionManager {
	return &TransactionManager{
		db:     db,
		logger: logger,
	}
}

// WithTransaction executes the provided function within a database transaction.
func (tm *TransactionManager) WithTransaction(
	ctx context.Context,
	fn func(tx *gorm.DB) error,
) error {
	return tm.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			if err := fn(tx); err != nil {
				tm.logger.Error("Transaction failed - rolled back", "error", err)
				return err
			}
			tm.logger.Info("Transaction committed successfully")
			return nil
		})
}

// BeginTransaction starts a new database transaction and returns the transaction object.
func (tm *TransactionManager) BeginTransaction(ctx context.Context) (*gorm.DB, error) {
	tx := tm.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		tm.logger.Error("Failed to begin transaction", "error", tx.Error)
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	tm.logger.Info("Transaction started")
	return tx, nil
}

// CommitTransaction commits the provided transaction.
func (tm *TransactionManager) CommitTransaction(tx *gorm.DB) error {
	if err := tx.Commit().Error; err != nil {
		tm.logger.Error("Failed to commit transaction", "error", err)
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

// RollbackTransaction rolls back the provided transaction.
func (tm *TransactionManager) RollbackTransaction(tx *gorm.DB) error {
	if err := tx.Rollback().Error; err != nil {
		tm.logger.Error("Failed to rollback transaction", "error", err)
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}
	tm.logger.Info("Transaction rolled back")
	return nil
}

// GetDBContext returns a gorm.DB instance with the provided context.
func (tm *TransactionManager) GetDBContext(ctx context.Context) *gorm.DB {
	return tm.db.WithContext(ctx)
}

// HealthCheck verifies the database connection is alive.
func (tm *TransactionManager) HealthCheck(ctx context.Context) error {
	sqlDB, err := tm.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		tm.logger.Error("Database ping failed", "error", err)
		return fmt.Errorf("database ping failed: %w", err)
	}

	tm.logger.Info("Database connection is healthy")
	return nil
}

// GetStats returns database connection statistics
func (tm *TransactionManager) GetStats(ctx context.Context) (sql.DBStats, error) {
	sqlDB, err := tm.db.DB()
	if err != nil {
		return sql.DBStats{}, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	stats := sqlDB.Stats()
	tm.logger.Debug("Database stats retrieved",
		"open_connections", stats.OpenConnections,
		"in_use", stats.InUse,
		"idle", stats.Idle)

	return stats, nil
}
