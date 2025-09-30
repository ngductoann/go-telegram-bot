# Makefile for go-telegram-bot

BINARY_NAME=go-telegram-bot
BUILD_FOLDER=build
CMD_PATH=cmd/bot

.PHONY: build run clean fmt vet

build:
	go build -o ./${BUILD_FOLDER}/$(BINARY_NAME) $(CMD_PATH)/bot.go

run: build
	./${BUILD_FOLDER}/$(BINARY_NAME)

clean:
	rm -f ./${BUILD_FOLDER}/$(BINARY_NAME)

fmt:
	go fmt ./...

vet:
	go vet ./...

deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

# Database migration using golang-migrate

# Database commands
db-migrate:
	@echo "Running database migrations..."
	go run ./cmd/migrate/ -action=migrate

db-reset:
	@echo "Warning: This will drop all tables and recreate them!"
	@echo "Are you sure? Press Ctrl+C to cancel, Enter to continue..."
	@read confirm
	@echo "Resetting database (drop + migrate + indexes + seed)..."
	go run ./cmd/migrate/ -action=reset

db-drop:
	@echo "Warning: This will drop all tables!"
	@echo "Are you sure? Press Ctrl+C to cancel, Enter to continue..."
	@read confirm
	@echo "Dropping all tables..."
	go run ./cmd/migrate/ -action=drop
