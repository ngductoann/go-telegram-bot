# Variables
CMD_DIR=./cmd/bot
BUILD_DIR=./bin
BINARY_NAME=telegram-bot

# Run the application
run:
	@echo "Running the bot..."
	@go run $(CMD_DIR)/main.go

# Build the application
build:
	@echo "Building the bot..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)/main.go
	@echo "Build complete. Executable is 'bot'."

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleanup complete."

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tiny


# Database commands
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
