package main

import (
	"flag"
	"fmt"
	"log"

	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/infrastructure/config"
	"go-telegram-bot/internal/infrastructure/database"

	"gorm.io/gorm"
)

func main() {
	action := flag.String("action", "migrate", "Migrate action: migrate, drop, reset")
	flag.Parse()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	db, err := database.NewPostgresConnection(cfg.GetDatabaseURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	switch *action {
	case "migrate":
		if err := runMigration(db); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		fmt.Println("Migrations completed successfully")
	case "drop":
		if err := dropTables(db); err != nil {
			log.Fatalf("Failed to drop tables: %v", err)
		}
		fmt.Println("All Tables dropped successfully")
	case "reset": // Drop and recreate tables
		// Drop tables first
		if err := dropTables(db); err != nil {
			log.Fatalf("Failed to drop tables: %v", err)
		}
		fmt.Println("All Tables dropped successfully")

		// Run migrations to recreate tables
		if err := runMigration(db); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		fmt.Println("Database reset successfully")
	default:
		log.Fatalf("Unknown action: %s", *action)
	}
}

// runMigration performs the auto-migration for all entities
func runMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Chat{},
		&entity.Message{},
		&entity.UserProfile{},
	)
}

// dropTables drops all tables for the entities
func dropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&entity.User{},
		&entity.Chat{},
		&entity.Message{},
		&entity.UserProfile{},
	)
}
