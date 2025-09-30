package usecase

import (
	"context"

	"go-telegram-bot/internal/application/dto"
	"go-telegram-bot/internal/domain/entity"
	"go-telegram-bot/internal/domain/repository"
	"go-telegram-bot/internal/domain/types"

	"github.com/google/uuid"
)

type MessageUseCase struct {
	userRepo     repository.UserRepository
	chatrepo     repository.ChatRepository
	meessageRepo repository.MessageRepository
}

// NewMessageUseCase creates a new instance of MessageUseCase
func NewMessageUseCase(
	userRepo repository.UserRepository,
	chatrepo repository.ChatRepository,
	meessageRepo repository.MessageRepository,
) *MessageUseCase {
	return &MessageUseCase{
		userRepo:     userRepo,
		chatrepo:     chatrepo,
		meessageRepo: meessageRepo,
	}
}

// GetByID retrieves a message by its UUID
func (uc *MessageUseCase) GetByID(
	ctx context.Context, id uuid.UUID,
) (*dto.MessageResponse, error) {
	message, err := uc.meessageRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapMessageToDTO(message), err
}

// GetByTelegramID retrieves a message by its Telegram message ID
func (uc *MessageUseCase) GetByTelegramID(
	ctx context.Context, telegramMessageID int64,
) (*dto.MessageResponse, error) {
	message, err := uc.meessageRepo.GetByTelegramID(ctx, telegramMessageID)
	if err != nil {
		return nil, err
	}
	return mapMessageToDTO(message), err
}

// GetByTelegramUserID retrieves messages by Telegram user ID
func (uc *MessageUseCase) GetByTelegramUserID(
	ctx context.Context, telegramUserID types.TelegramUserID,
) ([]*dto.MessageResponse, error) {
	// check user exists
	_, err := uc.userRepo.GetByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, err
	}

	// fetch messages
	messages, err := uc.meessageRepo.GetByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, err
	}

	return mapMessagesToDTO(messages), nil
}

// GetByTelegramChatID retrieves messages by Telegram chat ID
func (uc *MessageUseCase) GetByTelegramChatID(
	ctx context.Context, telegramChatID types.TelegramChatID,
) ([]*dto.MessageResponse, error) {
	_, err := uc.chatrepo.GetByTelegramChatID(ctx, telegramChatID)
	if err != nil {
		return nil, err
	}
	messages, err := uc.meessageRepo.GetByTelegramChatID(ctx, telegramChatID)
	if err != nil {
		return nil, err
	}

	return mapMessagesToDTO(messages), nil
}

// GetByChatID retrieves messages by internal chat UUID
func (uc *MessageUseCase) GetByChatID(
	ctx context.Context, chatID uuid.UUID,
) ([]*dto.MessageResponse, error) {
	_, err := uc.chatrepo.GetByUUID(ctx, chatID)
	if err != nil {
		return nil, err
	}
	messages, err := uc.meessageRepo.GetByChatID(ctx, chatID)
	if err != nil {
		return nil, err
	}
	return mapMessagesToDTO(messages), err
}

// GetByUserID retrieves messages by internal user UUID
func (uc *MessageUseCase) GetByUserID(
	ctx context.Context, userID uuid.UUID,
) ([]*dto.MessageResponse, error) {
	_, err := uc.userRepo.GetByUUID(ctx, userID)
	if err != nil {
		return nil, err
	}

	messages, err := uc.meessageRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return mapMessagesToDTO(messages), nil
}

// Create creates a new message
func (uc *MessageUseCase) Create(
	ctx context.Context, req *dto.CreateMessageRequest,
) (*dto.MessageResponse, error) {
	// check user exists
	user, err := uc.userRepo.GetByTelegramUserID(ctx, req.TelegramUserID)
	if err != nil {
		return nil, err
	}
	// check chat exists
	chat, err := uc.chatrepo.GetByTelegramChatID(ctx, req.TelegramChatID)
	if err != nil {
		return nil, err
	}

	message := entity.NewMessage(
		req.TelegramID,
		chat.ID,
		user.ID,
		req.Content,
		req.MessageType,
		req.ParseMode,
		nil,
	)
	if req.RepliedToID != nil {
		message, err := uc.meessageRepo.GetByTelegramID(ctx, *req.RepliedToID)
		if err != nil {
			return nil, err
		}
		message.ReplyToID = &message.ID
	}

	uc.meessageRepo.Create(ctx, message)
	return mapMessageToDTO(message), nil
}

// Update updates an existing message
func (uc *MessageUseCase) Update(
	ctx context.Context, req *dto.UpdateMessageRequest,
) (*dto.MessageResponse, error) {
	message, err := uc.meessageRepo.GetByTelegramID(ctx, req.TelegramID)
	if err != nil {
		return nil, err
	}
	message.Content = req.Content

	return mapMessageToDTO(message), nil
}

// Delete marks a message as deleted by its Telegram message ID
func (uc *MessageUseCase) Delete(
	ctx context.Context, telegramMessageID int64,
) error {
	return uc.meessageRepo.Delete(ctx, telegramMessageID)
}

// Update updates an existing message
func mapMessageToDTO(message *entity.Message) *dto.MessageResponse {
	return &dto.MessageResponse{
		ID:          message.ID,
		TelegramID:  message.TelegramMessageID,
		ChatID:      message.ChatID,
		UserID:      message.UserID,
		Content:     message.Content,
		MessageType: message.MessageType,
		ReplyToID:   message.ReplyToID,
		IsEdited:    message.IsEdited,
		IsDeleted:   message.IsDeleted,
		CreatedAt:   message.CreatedAt,
		UpdatedAt:   message.UpdatedAt,
	}
}

// mapMessagesToDTO maps a slice of Message entities to a slice of MessageResponse DTOs
func mapMessagesToDTO(messages []*entity.Message) []*dto.MessageResponse {
	messageDTOs := make([]*dto.MessageResponse, len(messages))
	for i, message := range messages {
		messageDTOs[i] = mapMessageToDTO(message)
	}
	return messageDTOs
}
