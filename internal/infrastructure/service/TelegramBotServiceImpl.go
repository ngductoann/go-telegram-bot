package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go-telegram-bot/internal/domain/entity"
	domainService "go-telegram-bot/internal/domain/service"
)

type telegramBot struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

// NewTelegramBot creates a new instance of TelegramBotService.
func NewTelegramBot(token string) domainService.TelegramBotService {
	return &telegramBot{
		token:   token,
		baseURL: "https://api.telegram.org/bot" + token,
		httpClient: &http.Client{
			Timeout: 40 * time.Second, // Increased to handle 30s long polling + buffer
		},
	}
}

// isTimeoutError checks if the error is a timeout-related error that shouldn't be logged as critical
func isTimeoutError(err error) bool {
	if err == nil {
		return false
	}

	errorMsg := err.Error()

	// Check for common timeout error patterns
	if strings.Contains(errorMsg, "context deadline exceeded") ||
		strings.Contains(errorMsg, "Client.Timeout exceeded") ||
		strings.Contains(errorMsg, "timeout") {
		return true
	}

	// Check for network timeout errors
	var netError net.Error
	if errors.As(err, &netError) && netError.Timeout() {
		return true
	}

	return false
}

// SendMessage sends a message to a Telegram chat.
func (b *telegramBot) SendMessage(
	ctx context.Context,
	chatID int64,
	text string,
) error {
	data := map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	url := b.baseURL + "/sendMessage"
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := b.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("non-200 response: %d - %s", resp.StatusCode, string(body))
	}

	return nil
}

// GetUpdates retrieves updates from the Telegram bot.
func (b *telegramBot) GetUpdates(ctx context.Context, offset int64) ([]entity.TelegramUpdate, error) {
	params := url.Values{}
	params.Add("offset", fmt.Sprintf("%d", offset))
	params.Add("timeout", "30")

	url := b.baseURL + "/getUpdates?" + params.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	resp, err := b.httpClient.Do(req)
	if err != nil {
		// Check if it's a timeout error - these are expected during long polling
		if isTimeoutError(err) {
			// Return empty results instead of error for timeout (silent handling)
			return []entity.TelegramUpdate{}, nil
		}
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response: %d", resp.StatusCode)
	}

	var response struct {
		OK     bool                    `json:"ok"`
		Result []entity.TelegramUpdate `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if !response.OK {
		return nil, fmt.Errorf("API response not OK")
	}

	return response.Result, nil
}
