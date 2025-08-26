package service

import (
	"github.com/go-telegram/bot/models"
	"github.com/ngductoann/go-telegram-bot/internal/domain/entity"
)

// UpdateLocal represents the local update structure
type UpdateLocal struct {
	Update       *models.Update
	User         *entity.User
	Chat         *entity.Chat
	Message      *entity.Message
	CallbackData *entity.CallbackData
	UserSession  *entity.UserSession
}

// Logger Defines the interface for logging
type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
	Panic(msg string, fields ...interface{})

	// Provide structured logging with error and message
	CheckError(err error, msg string, fatal bool)
	CheckErrorPanic(err error, msg string)
	CheckErrorFatal(err error, msg string)
}
