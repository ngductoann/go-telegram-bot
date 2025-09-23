package errors

import (
	"errors"
)

// Domain errors
var (
	// User errors
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidUserData   = errors.New("invalid user data")

	// Profile errors
	ErrProfileNotFound      = errors.New("profile not found")
	ErrBioTooLong           = errors.New("bio is too long")
	ErrInappropriateContent = errors.New("content contains inappropriate words")
	ErrInvalidAvatarURL     = errors.New("invalid avatar URL")
	ErrTooManyPreferences   = errors.New("too many preferences")
	ErrInvalidTimezone      = errors.New("invalid timezone")

	// Chat errors
	ErrChatNotFound      = errors.New("chat not found")
	ErrChatAlreadyExists = errors.New("chat already exists")
	ErrInvalidChatData   = errors.New("invalid chat data")

	// Message errors
	ErrMessageNotFound       = errors.New("message not found")
	ErrInvalidMessageData    = errors.New("invalid message data")
	ErrMessageAlreadyDeleted = errors.New("message already deleted")

	// Session errors
	ErrSessionNotFound    = errors.New("session not found")
	ErrSessionExpired     = errors.New("session expired")
	ErrInvalidSessionData = errors.New("invalid session data")

	// Flow errors
	ErrHandlerNotFound      = errors.New("handler not found")
	ErrInvalidFlow          = errors.New("invalid flow configuration")
	ErrInvalidCallbackData  = errors.New("invalid callback data")
	ErrFlowValidationFailed = errors.New("flow validation failed")

	// General errors
	ErrInvalidInput     = errors.New("invalid input")
	ErrPermissionDenied = errors.New("permission denied")
	ErrInternalError    = errors.New("internal error")
)
