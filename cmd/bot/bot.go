package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go-telegram-bot/internal/infrastructure/initialize"
)

func main() {
	log.Println("Bot is starting...")

	// Initialize container with all dependencies
	container, err := initialize.NewContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}
	container.Logger.Info("Successfully loaded config and initialized dependencies")

	// Initialize delivery layer - now using factory from container
	telegramHandler := container.TelegramHandler
	container.Logger.Info("Delivery layer initialized")

	// Create context with cancel for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// run bot in goroutine
	errChan := make(chan error, 1)
	go func() {
		container.Logger.Info("🤖 Bot is start polling from telegram...")
		if err := telegramHandler.StartPolling(ctx); err != nil && err != context.Canceled {
			errChan <- err
		}
	}()

	container.Logger.Info("Bot is running. Press Ctrl+C to stop.")

	// wait signal or error
	select {
	case <-sigChan:
		// received an interrupt signal, shut down gracefully
		container.Logger.Info("Received shutdown signal, exiting...")
	case err := <-errChan:
		// received an error from the bot
		container.Logger.Error("Bot encountered an error", "error", err)
	}

	// Perform graceful shutdown
	cancel()
	log.Println("Bot stopped gracefully.")
}
