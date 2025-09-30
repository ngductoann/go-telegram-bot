package usecase

import (
	"context"

	"go-telegram-bot/internal/application/dto"
	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/errors"
	"go-telegram-bot/internal/domain/repository"
	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
)

// ChatUseCase handles chat-related business logic.
type ChatUseCase struct {
	chatRepo repository.ChatRepository
}

// NewChatUseCase creates a new instance of ChatUseCase with the provided ChatRepository.
func NewChatUseCase(chatRepo repository.ChatRepository) *ChatUseCase {
	return &ChatUseCase{chatRepo: chatRepo}
}

// GetByUUID retrieves a chat by its UUID.
func (u *ChatUseCase) GetByUUID(
	ctx context.Context, id uuid.UUID,
) (*dto.ChatResponse, error) {
	chat, err := u.chatRepo.GetByUUID(ctx, id)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		return nil, errors.ErrChatNotFound
	}

	return mapChatToDTO(chat), nil
}

// GetByTelegramChatID retrieves a chat by its Telegram chat ID.
func (u *ChatUseCase) GetByTelegramChatID(
	ctx context.Context, chatID types.TelegramChatID,
) (*dto.ChatResponse, error) {
	chat, err := u.chatRepo.GetByTelegramChatID(ctx, chatID)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		return nil, errors.ErrChatNotFound
	}

	return mapChatToDTO(chat), nil
}

// CreateChat creates a new chat if it doesn't already exist.
func (u *ChatUseCase) CreateChat(
	ctx context.Context, req *dto.CreateChatRequest,
) (*dto.ChatResponse, error) {
	// check if chat already exists
	existingChat, err := u.chatRepo.GetByTelegramChatID(ctx, req.TelegramID)
	if err != nil && existingChat != nil {
		return nil, errors.ErrChatAlreadyExists
	}

	// Create new chat entity
	chat := entity.NewChat(
		req.TelegramID, req.Type,
	)
	chat.Title = req.Title
	chat.Username = req.Username
	chat.Description = req.Description

	// Save to repository
	if err := u.chatRepo.Create(ctx, chat); err != nil {
		return nil, err
	}

	return mapChatToDTO(chat), nil
}

// UpdateChat updates the details of an existing chat.
func (u *ChatUseCase) UpdateChat(
	ctx context.Context, chatID types.TelegramChatID, req *dto.UpdateChatRequest,
) (*dto.ChatResponse, error) {
	// check if chat exists
	chat, err := u.chatRepo.GetByTelegramChatID(ctx, chatID)
	if err != nil {
		return nil, err
	}
	if chat == nil {
		return nil, errors.ErrChatNotFound
	}

	// Update chat details
	chat.Title = req.Title
	chat.Username = req.Username
	chat.Description = req.Description
	if err := u.chatRepo.Update(ctx, chat); err != nil {
		return nil, err
	}

	return mapChatToDTO(chat), nil
}

// DeleteChat deletes a chat by its Telegram chat ID.
func (u *ChatUseCase) DeleteChat(
	ctx context.Context, chatID types.TelegramChatID,
) error {
	// check if chat exists
	_, err := u.chatRepo.GetByTelegramChatID(ctx, chatID)
	if err != nil {
		return err
	}
	return u.chatRepo.Delete(ctx, chatID)
}

// GetActiveChats retrieves a list of active chats up to the specified limit.
func (u *ChatUseCase) GetActiveChats(
	ctx context.Context, limit int,
) ([]*dto.ChatResponse, error) {
	chats, err := u.chatRepo.GetActiveChats(ctx, limit)
	if err != nil {
		return nil, err
	}

	chatDTOs := make([]*dto.ChatResponse, len(chats))
	for i, chat := range chats {
		chatDTOs[i] = mapChatToDTO(chat)
	}
	return chatDTOs, nil
}

// UpdateChat updates the details of an existing chat.
func mapChatToDTO(chat *entity.Chat) *dto.ChatResponse {
	return &dto.ChatResponse{
		ID:          chat.ID,
		TelegramID:  chat.TelegramChatID,
		Title:       chat.Title,
		Username:    chat.Username,
		Description: chat.Description,
		IsActive:    chat.IsActive,
		CreatedAt:   chat.CreatedAt,
		UpdatedAt:   chat.UpdatedAt,
	}
}
