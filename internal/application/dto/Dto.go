package dto

import (
	"go-telegram-bot/internal/domain/types"
	"time"

	"github.com/google/uuid"
)

// CreateUserRequest represents the payload to create a new user
type CreateUserRequest struct {
	TelegramID types.TelegramUserID `json:"telegram_id" validate:"required" example:"123456789"`
	Username   *string              `json:"username" example:"johndoe"`
	FirstName  string               `json:"first_name" validate:"required" example:"John"`
	LastName   *string              `json:"last_name" example:"Doe"`
	IsBot      bool                 `json:"is_bot" example:"false"`
}

// UserResponse represents the user data returned in responses
type UserResponse struct {
	ID         uuid.UUID            `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	TelegramID types.TelegramUserID `json:"telegram_id" example:"123456789"`
	Username   *string              `json:"username" example:"johndoe"`
	FirstName  string               `json:"first_name" example:"John"`
	LastName   *string              `json:"last_name" example:"Doe"`
	IsActive   bool                 `json:"is_active" example:"true"`
	IsBot      bool                 `json:"is_bot" example:"false"`
	LastSeenAt *time.Time           `json:"last_seen_at,omitempty" example:"2023-10-05T14:48:00Z"`
	CreatedAt  time.Time            `json:"created_at" example:"2023-10-01T12:00:00Z"`
	UpdatedAt  time.Time            `json:"updated_at" example:"2023-10-05T14:48:00Z"`
}

type UpdateUserRequest struct {
	Username  *string `json:"username,omitempty" example:"johndoe"`
	FirstName string  `json:"first_name,omitempty" example:"John"`
	LastName  *string `json:"last_name,omitempty" example:"Doe"`
}

// CreateChatRequest represents the payload to create a new chat
type CreateChatRequest struct {
	TelegramID  types.TelegramChatID `json:"telegram_id" validate:"required" example:"-1001234567890"`
	Type        types.ChatType       `json:"type" validate:"required,oneof=private group supergroup channel" example:"private"`
	Title       *string              `json:"title,omitempty" example:"My Group"`
	Username    *string              `json:"username,omitempty" example:"mygroup"`
	Description *string              `json:"description,omitempty" example:"This is a group chat"`
}

type UpdateChatRequest struct {
	Title       *string `json:"title,omitempty" example:"My Group"`
	Username    *string `json:"username,omitempty" example:"mygroup"`
	Description *string `json:"description,omitempty" example:"This is a group chat"`
}

// ChatResponse represents the chat data returned in responses
type ChatResponse struct {
	ID          uuid.UUID            `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	TelegramID  types.TelegramChatID `json:"telegram_id" example:"-1001234567890"`
	Title       *string              `json:"title,omitempty" example:"My Group"`
	Username    *string              `json:"username,omitempty" example:"mygroup"`
	Description *string              `json:"description,omitempty" example:"This is a group chat"`
	IsActive    bool                 `json:"is_active" example:"true"`
	CreatedAt   time.Time            `json:"created_at" example:"2023-10-01T12:00:00Z"`
	UpdatedAt   time.Time            `json:"updated_at" example:"2023-10-05T14:48:00Z"`
}

// CreateMessageRequest represents the payload to create a new message
type CreateMessageRequest struct {
	TelegramID     int64                `json:"telegram_id" validate:"required" example:"123456789"`
	TelegramChatID types.TelegramChatID `json:"telegram_chat_id" validate:"required" example:"-1001234567890"`
	TelegramUserID types.TelegramUserID `json:"telegram_user_id" validate:"required" example:"123456789"`
	Content        string               `json:"content" validate:"required" example:"Hello, world!"`
	MessageType    *types.MessageType   `json:"message_type" validate:"required,oneof=text image video audio document" example:"text"`
	RepliedToID    *int64               `json:"replied_to_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440000"`
	ParseMode      *types.ParseMode     `json:"parse_mode,omitempty" validate:"omitempty,oneof=Markdown MarkdownV2 HTML" example:"MarkdownV2"`
}

type UpdateMessageRequest struct {
	TelegramID int64  `json:"telegram_id" validate:"required" example:"123456789"`
	Content    string `json:"content" validate:"required" example:"Hello, world!"`
}

// MessageResponse represents the message data returned in responses
type MessageResponse struct {
	ID          uuid.UUID          `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	TelegramID  int64              `json:"telegram_id" example:"123456789"`
	UserID      uuid.UUID          `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ChatID      uuid.UUID          `json:"chat_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Content     string             `json:"content" example:"Hello, world!"`
	MessageType *types.MessageType `json:"message_type" example:"text"`
	ReplyToID   *uuid.UUID         `json:"replied_to_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440000"`
	ParseMode   *types.ParseMode   `json:"parse_mode,omitempty" example:"MarkdownV2"`
	IsEdited    bool               `json:"is_edited" example:"false"`
	IsDeleted   bool               `json:"is_deleted" example:"false"`
	CreatedAt   time.Time          `json:"created_at" example:"2023-10-01T12:00:00Z"`
	UpdatedAt   time.Time          `json:"updated_at" example:"2023-10-05T14:48:00Z"`
}
