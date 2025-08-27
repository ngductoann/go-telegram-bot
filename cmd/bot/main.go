package main

import (
	"log"

	"github.com/ngductoann/go-telegram-bot/internal/infrastructure/container"
)

func main() {
	log.Println("Starting Telegram Bot..")
	c, err := container.NewContainer()

	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	c.Logger.Info("Container initialized successfully")
}
