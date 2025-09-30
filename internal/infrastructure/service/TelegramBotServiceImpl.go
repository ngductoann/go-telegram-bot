package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/domain/types"
	"go-telegram-bot/internal/infrastructure/config"
)

type telegramBot struct {
	config     config.ClientConfig
	httpClient *http.Client
	logger     domainService.Logger
	metrics    *ClientMetrics
}

// ClientMetrics tracks API performance metrics
type ClientMetrics struct {
	RequestCount    int64         `json:"request_count"`
	ErrorCount      int64         `json:"error_count"`
	RetryCount      int64         `json:"retry_count"`
	AverageLatency  time.Duration `json:"average_latency"`
	LastRequestTime time.Time     `json:"last_request_time"`
	RateLimitHits   int64         `json:"rate_limit_hits"`
}

// NewTelegramBot creates a new instance of TelegramBotService.
func NewTelegramBot(
	config config.ClientConfig, httpClient *http.Client, logger domainService.Logger,
) domainService.TelegramBotService {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: config.Timeout,
		}
	}

	// // Ensure BaseURL includes the bot token
	// if config.BaseURL != "" && config.Token != "" {
	// 	// If BaseURL doesn't end with the token, append it
	// 	if config.BaseURL == "https://api.telegram.org/bot" {
	// 		config.BaseURL = config.BaseURL + config.Token
	// 	}
	// } else if config.Token != "" {
	// 	// Default BaseURL with token
	// 	config.BaseURL = "https://api.telegram.org/bot" + config.Token
	// }

	return &telegramBot{
		config:     config,
		logger:     logger,
		httpClient: httpClient,
		metrics:    &ClientMetrics{},
	}
}

// SendMessageWithResponse sends a message and returns the full response for detailed handling
func (b *telegramBot) SendMessageWithResponse(
	ctx context.Context, request *types.SendMessageRequest,
) (*types.SendMessageResponse, error) {
	startTime := time.Now()

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	response, err := b.makeRequestWithRetry(ctx, "POST", "/sendMessage", requestBody, nil)
	if err != nil {
		return nil, err
	}

	var sendResponse types.SendMessageResponse
	if err := json.Unmarshal(response, &sendResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Update metrics
	b.updateMetrics(startTime, err)

	if b.config.EnableLogging && b.logger != nil {
		b.logger.Debug("SendMessage completed",
			"chat_id", request.ChatID,
			"success", sendResponse.IsSuccess(),
			"duration", time.Since(startTime),
		)
	}

	return &sendResponse, nil
}

// GetUpdatesWithResponse retrieves updates and returns the full response for detailed handling
func (b *telegramBot) GetUpdatesWithResponse(
	ctx context.Context, request *types.GetUpdatesRequest,
) (*types.GetUpdatesResponse, error) {
	startTime := time.Now()

	// Build query parameters
	params := make(map[string]any)
	if request.Offset > 0 {
		params["offset"] = request.Offset
	}
	if request.Limit != nil {
		params["limit"] = *request.Limit
	}
	if request.Timeout != nil {
		params["timeout"] = *request.Timeout
	}
	if len(request.AllowedUpdates) > 0 {
		params["allowed_updates"] = request.AllowedUpdates
	}

	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	response, err := b.makeRequestWithRetry(ctx, "POST", "/getUpdates", requestBody, nil)
	if err != nil {
		return nil, err
	}

	var updatesResponse types.GetUpdatesResponse
	if err := json.Unmarshal(response, &updatesResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check if the API returned an error even with HTTP 200
	if updatesResponse.HasError() {
		return nil, &types.ResponseError{
			Method:      "/getUpdates",
			RequestData: map[string]any{"params": params},
			Response:    &updatesResponse.BaseResponse,
			HTTPStatus:  200,
			Timestamp:   time.Now(),
		}
	}

	// Update metrics
	b.updateMetrics(startTime, err)

	return &updatesResponse, nil
}

// GetMeWithResponse returns information about the bot with detailed response handling
func (b *telegramBot) GetMeWithResponse(
	ctx context.Context,
) (*types.GetMeResponse, error) {
	startTime := time.Now()

	response, err := b.makeRequestWithRetry(ctx, "GET", "/getMe", nil, nil)
	if err != nil {
		return nil, err
	}

	var getMeResponse types.GetMeResponse
	if err := json.Unmarshal(response, &getMeResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check if the API returned an error even with HTTP 200
	if getMeResponse.HasError() {
		return nil, &types.ResponseError{
			Method:      "/getMe",
			RequestData: map[string]any{},
			Response:    &getMeResponse.BaseResponse,
			HTTPStatus:  200,
			Timestamp:   time.Now(),
		}
	}

	// Update metrics
	b.updateMetrics(startTime, err)

	return &getMeResponse, nil
}

// DeleteWebhookWithResponse removes the webhook and returns the full response for detailed handling
func (b *telegramBot) DeleteWebhookWithResponse(
	ctx context.Context,
) (*types.DeleteWebhookResponse, error) {
	startTime := time.Now()

	response, err := b.makeRequestWithRetry(ctx, "POST", "/deleteWebhook", nil, nil)
	if err != nil {
		return nil, err
	}

	var deleteResponse types.DeleteWebhookResponse
	if err := json.Unmarshal(response, &deleteResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check if the API returned an error even with HTTP 200
	if deleteResponse.HasError() {
		return nil, &types.ResponseError{
			Method:      "/deleteWebhook",
			RequestData: map[string]any{},
			Response:    &deleteResponse.BaseResponse,
			HTTPStatus:  200,
			Timestamp:   time.Now(),
		}
	}

	// Update metrics
	b.updateMetrics(startTime, err)

	return &deleteResponse, nil
}

// SendMessages sends multiple messages in a batch, respecting rate limits
func (b *telegramBot) SendMessages(
	ctx context.Context, requests []*types.SendMessageRequest,
) ([]*types.SendMessageResponse, error) {
	responses := make([]*types.SendMessageResponse, len(requests))

	for i, request := range requests {
		response, err := b.SendMessageWithResponse(ctx, request)
		if err != nil {
			return nil, fmt.Errorf("failed to send message to chat_id %v: %w", request.ChatID, err)
		}
		responses[i] = response

		// Add a small delay between requests to avoid hitting rate limits
		if i < len(requests)-1 {
			time.Sleep(b.config.RateLimitDelay)
		}
	}

	return responses, nil
}

// SendMessageWithRetry sends a message with retry logic for transient errors
func (b *telegramBot) SendMessageWithRetry(
	ctx context.Context, request *types.SendMessageRequest, maxRetries int,
) (*types.SendMessageResponse, error) {
	startTime := time.Now()

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	response, err := b.makeRequestWithRetry(ctx, "POST", "/sendMessage", requestBody, &maxRetries)
	if err != nil {
		return nil, err
	}

	var sendResponse types.SendMessageResponse
	if err := json.Unmarshal(response, &sendResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Update metrics
	b.updateMetrics(startTime, err)

	return &sendResponse, nil
}

// GetUpdatesWithRetry retrieves updates with retry logic for transient errors
func (b *telegramBot) GetUpdatesWithRetry(
	ctx context.Context, request *types.GetUpdatesRequest, maxRetries int,
) (*types.GetUpdatesResponse, error) {
	startTime := time.Now()

	// Build query parameters
	params := make(map[string]any)
	if request.Offset > 0 {
		params["offset"] = request.Offset
	}
	if request.Limit != nil {
		params["limit"] = *request.Limit
	}
	if request.Timeout != nil {
		params["timeout"] = *request.Timeout
	}
	if len(request.AllowedUpdates) > 0 {
		params["allowed_updates"] = request.AllowedUpdates
	}

	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	response, err := b.makeRequestWithRetry(ctx, "POST", "/getUpdates", requestBody, &maxRetries)
	if err != nil {
		return nil, err
	}

	var updatesResponse types.GetUpdatesResponse
	if err := json.Unmarshal(response, &updatesResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Update metrics
	b.updateMetrics(startTime, err)

	return &updatesResponse, nil
}

// Placeholder implementations for other methods - implement as needed
func (b *telegramBot) SetWebhook(ctx context.Context, request *types.SetWebhookRequest) (*types.SetWebhookResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) GetWebhookInfo(ctx context.Context) (*types.GetWebhookInfoResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) GetFile(ctx context.Context, fileID string) (*types.GetFileResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) DownloadFile(ctx context.Context, filePath string) ([]byte, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) EditMessageText(ctx context.Context, request *types.EditMessageTextRequest) (*types.EditMessageTextResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) DeleteMessage(ctx context.Context, chatID types.TelegramChatID, messageID int64) (*types.DeleteMessageResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) ForwardMessage(ctx context.Context, request *types.ForwardMessageRequest) (*types.ForwardMessageResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) SendPhoto(ctx context.Context, request *types.SendPhotoRequest) (*types.SendPhotoResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) SendDocument(ctx context.Context, request *types.SendDocumentRequest) (*types.SendDocumentResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) GetChat(ctx context.Context, chatID types.TelegramChatID) (*types.GetChatResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) BanChatMember(ctx context.Context, request *types.BanChatMemberRequest) (*types.BanChatMemberResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) UnbanChatMember(ctx context.Context, request *types.UnbanChatMemberRequest) (*types.UnbanChatMemberResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) GetChatMember(ctx context.Context, chatID types.TelegramChatID, userID types.TelegramUserID) (*types.GetChatMemberResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) GetChatMembersCount(ctx context.Context, chatID types.TelegramChatID) (*types.GetChatMembersCountResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) AnswerCallbackQuery(ctx context.Context, request *types.AnswerCallbackQueryRequest) (*types.AnswerCallbackQueryResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

func (b *telegramBot) AnswerInlineQuery(ctx context.Context, request *types.AnswerInlineQueryRequest) (*types.AnswerInlineQueryResponse, error) {
	// Implementation here
	return nil, fmt.Errorf("not implemented")
}

// makeRequestWithRetry makes an HTTP request with retry logic for transient errors
func (b *telegramBot) makeRequestWithRetry(
	ctx context.Context, method, endpoint string, body []byte, maxRetries *int,
) ([]byte, error) {
	// Check if maxRetries is nil or invalid, use default config value
	if maxRetries == nil || *maxRetries < 0 {
		maxRetries = &b.config.MaxRetries
	}

	// Track the last error to return if all retries fail
	var lastErr error

	// Attempt the request up to maxRetries + 1 times
	for attempt := 0; attempt <= *maxRetries; attempt++ {

		// If this is a retry attempt, wait before retrying
		if attempt > 0 {
			retryDelay := b.config.RetryDelay * time.Duration(attempt)

			if b.config.EnableLogging && b.logger != nil {
				b.logger.Debug("Retrying request",
					"attempt", attempt,
					"delay", retryDelay,
					"endpoint", endpoint,
				)
			}

			// Respect context cancellation while waiting
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(retryDelay):
			}

			b.metrics.RetryCount++
		}

		// Make the actual request
		response, err := b.makeRequest(ctx, method, endpoint, body)
		if err == nil {
			return response, nil
		}

		lastErr = err

		// check if error retryable
		if respErr, ok := err.(*types.ResponseError); ok {
			if !respErr.ShouldRetry() {
				break
			}

			// Handle rate limiting (HTTP 429)
			if respErr.Response != nil && respErr.Response.IsRateLimited() {
				b.metrics.RateLimitHits++

				delay := respErr.GetRetryDelay()
				if delay <= 0 {
					delay = b.config.RateLimitDelay
				}

				if b.config.EnableLogging && b.logger != nil {
					b.logger.Warn("Rate limit hit, backing off",
						"endpoint", endpoint,
						"retry_after", delay,
					)
				}

				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				case <-time.After(delay):
				}
			}
		}

	}

	b.metrics.ErrorCount++
	return nil, lastErr
}

// makeRequest constructs and sends an HTTP request to the Telegram API
func (b *telegramBot) makeRequest(
	ctx context.Context, method, endpoint string, body []byte,
) ([]byte, error) {
	b.metrics.RequestCount++
	b.metrics.LastRequestTime = time.Now()

	url := b.config.BaseURL + endpoint

	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewBuffer(body)
	}

	req, err := b.makeHttpRequest(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	resp, err := b.httpClient.Do(req)
	if err != nil {
		return nil, &types.ResponseError{
			Method:      endpoint,
			RequestData: map[string]any{"body": string(body)},
			HTTPStatus:  0,
			Timestamp:   time.Now(),
		}
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		// Try to parse the error response to get detailed error information
		var baseResponse types.BaseResponse
		if parseErr := json.Unmarshal(responseBody, &baseResponse); parseErr == nil {
			return nil, &types.ResponseError{
				Method:      endpoint,
				RequestData: map[string]any{"body": string(body)},
				Response:    &baseResponse,
				HTTPStatus:  resp.StatusCode,
				Timestamp:   time.Now(),
			}
		}

		// If parsing fails, return error without parsed response
		return nil, &types.ResponseError{
			Method:      endpoint,
			RequestData: map[string]any{"body": string(body)},
			HTTPStatus:  resp.StatusCode,
			Timestamp:   time.Now(),
		}
	}

	return responseBody, nil
}

// makeHttpRequestContext creates an HTTP request with the given context, method, URL, and body
func (b *telegramBot) makeHttpRequest(
	ctx context.Context, method, url string, bodyReader io.Reader,
) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if b.config.UserAgent != "" {
		req.Header.Set("User-Agent", b.config.UserAgent)
	}

	return req, nil
}

// updateMetrics updates the client metrics based on the request duration and error status
func (b *telegramBot) updateMetrics(
	startTime time.Time, err error,
) {
	duration := time.Since(startTime)

	// Update average latency (simple moving average)
	if b.metrics.RequestCount > 0 {
		b.metrics.AverageLatency = time.Duration(
			(int64(b.metrics.AverageLatency)*b.metrics.RequestCount + int64(duration)) /
				(b.metrics.RequestCount + 1),
		)
	} else {
		b.metrics.AverageLatency = duration
	}
}

// GetMetrics returns the current client metrics
func (b *telegramBot) GetMetrics() *ClientMetrics {
	return b.metrics
}

// ResetMetrics resets the client metrics to initial state
func (b *telegramBot) ResetMetrics() {
	b.metrics = &ClientMetrics{}
}
