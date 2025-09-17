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
