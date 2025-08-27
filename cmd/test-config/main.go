package main

import (
	"fmt"
	"log"

	"github.com/ngductoann/go-telegram-bot/internal/infrastructure/config"
)

func main() {
	fmt.Println("=== Clean Configuration Loading Test ===")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("✅ Configuration loaded successfully!\n\n")

	// Test App Configuration
	fmt.Printf("🤖 App Configuration:\n")
	fmt.Printf("  Bot Token: %s\n", cfg.GetBotToken())
	fmt.Printf("  Environment: %s\n", cfg.GetEnvironment())
	fmt.Printf("  Is Development: %t\n", cfg.IsDevelopment())
	fmt.Printf("  Is Production: %t\n", cfg.IsProduction())

	// Test Database Configuration
	fmt.Printf("\n🗄️  Database Configuration:\n")
	fmt.Printf("  Host: %s\n", cfg.Database.Host)
	fmt.Printf("  Port: %d\n", cfg.Database.Port)
	fmt.Printf("  User: %s\n", cfg.Database.User)
	fmt.Printf("  Name: %s\n", cfg.Database.Name)
	fmt.Printf("  SSL Mode: %s\n", cfg.Database.SSLMode)
	fmt.Printf("  Max Idle Conns: %d\n", cfg.Database.MaxIdleConns)
	fmt.Printf("  Max Open Conns: %d\n", cfg.Database.MaxOpenConns)
	fmt.Printf("  Connection URL: %s\n", cfg.GetDatabaseURL())

	// Test Redis Configuration
	fmt.Printf("\n📦 Redis Configuration:\n")
	fmt.Printf("  Host: %s\n", cfg.Redis.Host)
	fmt.Printf("  Port: %s\n", cfg.Redis.Port)
	fmt.Printf("  DB: %d\n", cfg.Redis.DB)
	fmt.Printf("  Redis URL: %s\n", cfg.GetRedisURL())

	// Test Log Configuration
	fmt.Printf("\n📝 Log Configuration:\n")
	fmt.Printf("  Log Level: %s\n", cfg.GetLogLevel())
	fmt.Printf("  Log File Path: %s\n", cfg.GetLogFilePath())
	fmt.Printf("  Max Size: %d MB\n", cfg.Log.MaxSize)
	fmt.Printf("  Max Backups: %d\n", cfg.Log.MaxBackups)
	fmt.Printf("  Max Age: %d days\n", cfg.Log.MaxAge)
	fmt.Printf("  Compress: %t\n", cfg.Log.Compress)

	fmt.Printf("\n🎉 Configuration test completed successfully!\n")
}
