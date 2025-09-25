package initialize

import (
	"fmt"

	"go-telegram-bot/internal/infrastructure/config"
	"go-telegram-bot/internal/infrastructure/database"
)

// GetDatabaseURL constructs the database connection URL from the configuration.
func GetDatabaseURL(c config.Config) string {
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

// initDatabase initializes the database connection and sets up the connection pool.
func (c *Container) initDatabase() error {
	databaseURL := GetDatabaseURL(*c.Config)
	db, err := database.NewPostgresConnection(databaseURL)
	if err != nil {
		return err
	}
	database.SetPool(c.Config, db, c.Logger)
	c.DB = db
	return nil
}
