package service

import (
	"context"

	"go-telegram-bot/internal/domain/types"
)

// TelegramBotService defines the interface for Telegram bot operations
type TelegramBotService interface {
	// Enhanced methods with better response handling
	SendMessageWithResponse(ctx context.Context, request *types.SendMessageRequest) (*types.SendMessageResponse, error)
	GetUpdatesWithResponse(ctx context.Context, request *types.GetUpdatesRequest) (*types.GetUpdatesResponse, error)
	GetMeWithResponse(ctx context.Context) (*types.GetMeResponse, error)
	DeleteWebhookWithResponse(ctx context.Context) (*types.DeleteWebhookResponse, error)

	// Batch operations
	SendMessages(ctx context.Context, requests []*types.SendMessageRequest) ([]*types.SendMessageResponse, error)

	// Advanced operations with retry logic
	SendMessageWithRetry(ctx context.Context, request *types.SendMessageRequest, maxRetries int) (*types.SendMessageResponse, error)
	GetUpdatesWithRetry(ctx context.Context, request *types.GetUpdatesRequest, maxRetries int) (*types.GetUpdatesResponse, error)

	// Webhook management
	SetWebhook(ctx context.Context, request *types.SetWebhookRequest) (*types.SetWebhookResponse, error)
	GetWebhookInfo(ctx context.Context) (*types.GetWebhookInfoResponse, error)

	// File operations
	GetFile(ctx context.Context, fileID string) (*types.GetFileResponse, error)
	DownloadFile(ctx context.Context, filePath string) ([]byte, error)

	// Message management
	EditMessageText(ctx context.Context, request *types.EditMessageTextRequest) (*types.EditMessageTextResponse, error)
	DeleteMessage(ctx context.Context, chatID types.TelegramChatID, messageID int64) (*types.DeleteMessageResponse, error)
	ForwardMessage(ctx context.Context, request *types.ForwardMessageRequest) (*types.ForwardMessageResponse, error)

	// Media sending
	SendPhoto(ctx context.Context, request *types.SendPhotoRequest) (*types.SendPhotoResponse, error)
	SendDocument(ctx context.Context, request *types.SendDocumentRequest) (*types.SendDocumentResponse, error)

	// Chat management
	GetChat(ctx context.Context, chatID types.TelegramChatID) (*types.GetChatResponse, error)
	BanChatMember(ctx context.Context, request *types.BanChatMemberRequest) (*types.BanChatMemberResponse, error)
	UnbanChatMember(ctx context.Context, request *types.UnbanChatMemberRequest) (*types.UnbanChatMemberResponse, error)
	GetChatMember(ctx context.Context, chatID types.TelegramChatID, userID types.TelegramUserID) (*types.GetChatMemberResponse, error)
	GetChatMembersCount(ctx context.Context, chatID types.TelegramChatID) (*types.GetChatMembersCountResponse, error)

	// Callback and inline query handling
	AnswerCallbackQuery(ctx context.Context, request *types.AnswerCallbackQueryRequest) (*types.AnswerCallbackQueryResponse, error)
	AnswerInlineQuery(ctx context.Context, request *types.AnswerInlineQueryRequest) (*types.AnswerInlineQueryResponse, error)
}
