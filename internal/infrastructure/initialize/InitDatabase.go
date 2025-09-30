package initialize

import (
	"go-telegram-bot/internal/infrastructure/database"
)

// initDatabase initializes the database connection and sets up the connection pool.
func (c *Container) initDatabase() error {
	databaseURL := c.Config.GetDatabaseURL()
	db, err := database.NewPostgresConnection(databaseURL)
	if err != nil {
		return err
	}
	database.SetPool(c.Config, db, c.Logger)
	c.DB = db
	return nil
}
