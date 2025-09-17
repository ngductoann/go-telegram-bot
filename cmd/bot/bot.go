package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go-telegram-bot/config"
	"go-telegram-bot/internal/delivery"
	infrasService "go-telegram-bot/internal/infrastructure/service"
	infrasUsecase "go-telegram-bot/internal/usecase"
)

func main() {
	log.Println("Bot is starting...")
	cfg := config.MustLoadConfig()
	log.Println("Successfully loaded config:")

	// Initialize dependencies with clean architecture principles

	// Initialize Infrastructure layer (e.g., database, external services)
	ipService := infrasService.NewIPService()
	telegramBot := infrasService.NewTelegramBot(cfg.TelegramBotToken)
	log.Println("Infrastructure service layer initialized.")

	// Use Case layer
	botUseCase := infrasUsecase.NewBotUseCase(ipService, telegramBot)
	log.Println("Use case layer initialized.")

	// Delivery layer
	telegramHandler := delivery.NewTelegramHandler(botUseCase, telegramBot)
	log.Println("Delivery layer initialized.")

	// Create context with cancel for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// run bot in goroutine
	errChan := make(chan error, 1)
	go func() {
		log.Println("🤖 Bot is start polling from telegram...")
		if err := telegramHandler.StartPolling(ctx); err != nil && err != context.Canceled {
			errChan <- err
		}
	}()

	log.Println("Bot started successfully.")
	log.Println("Bot is running. Press Ctrl+C to stop.")

	// wait signal or error
	select {
	case <-sigChan:
		// received an interrupt signal, shut down gracefully
		log.Println("Received shutdown signal, exiting...")
	case err := <-errChan:
		// received an error from the bot
		log.Printf("Bot encountered an error: %v", err)
	}

	// Perform graceful shutdown
	cancel()
	log.Println("Bot stopped gracefully.")
}
