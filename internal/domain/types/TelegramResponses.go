package types

import (
	"time"
)

// BaseResponse represents the common structure for all Telegram API responses
type BaseResponse struct {
	OK          bool            `json:"ok"`
	ErrorCode   *int            `json:"error_code,omitempty"`
	Description *string         `json:"description,omitempty"`
	Parameters  *ResponseParams `json:"parameters,omitempty"`
}

// ResponseParams contains information about why a request was unsuccessful
type ResponseParams struct {
	MigrateToChatID  *int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter       *int   `json:"retry_after,omitempty"`
	TooManyRequests  *bool  `json:"too_many_requests,omitempty"`
	ChatNotFound     *bool  `json:"chat_not_found,omitempty"`
	UserNotFound     *bool  `json:"user_not_found,omitempty"`
	BotBlocked       *bool  `json:"bot_blocked,omitempty"`
	BotKicked        *bool  `json:"bot_kicked,omitempty"`
	InvalidBotToken  *bool  `json:"invalid_bot_token,omitempty"`
	NetworkError     *bool  `json:"network_error,omitempty"`
	InternalAPIError *bool  `json:"internal_api_error,omitempty"`
}

// APIResponse represents a generic Telegram API response with a result of type T
type APIResponse[T any] struct {
	BaseResponse
	Result *T `json:"result,omitempty"`
}

// Specific response types for different Telegram API methods
type (
	// SendMessageResponse represents the response from sendMessage API
	SendMessageResponse = APIResponse[TelegramMessage]

	// GetUpdatesResponse represents the response from getUpdates API
	GetUpdatesResponse = APIResponse[[]TelegramUpdate]

	// GetMeResponse represents the response from getMe API
	GetMeResponse = APIResponse[TelegramUser]

	// GetChatResponse represents the response from getChat API
	GetChatResponse = APIResponse[TelegramChat]

	// DeleteWebhookResponse represents the response from deleteWebhook API
	DeleteWebhookResponse = APIResponse[bool]

	// SetWebhookResponse represents the response from setWebhook API
	SetWebhookResponse = APIResponse[bool]

	// GetWebhookInfoResponse represents the response from getWebhookInfo API
	GetWebhookInfoResponse = APIResponse[WebhookInfo]

	// EditMessageTextResponse represents the response from editMessageText API
	EditMessageTextResponse = APIResponse[TelegramMessage]

	// DeleteMessageResponse represents the response from deleteMessage API
	DeleteMessageResponse = APIResponse[bool]

	// ForwardMessageResponse represents the response from forwardMessage API
	ForwardMessageResponse = APIResponse[TelegramMessage]

	// SendPhotoResponse represents the response from sendPhoto API
	SendPhotoResponse = APIResponse[TelegramMessage]

	// SendDocumentResponse represents the response from sendDocument API
	SendDocumentResponse = APIResponse[TelegramMessage]

	// GetFileResponse represents the response from getFile API
	GetFileResponse = APIResponse[TelegramFile]

	// BanChatMemberResponse represents the response from banChatMember API
	BanChatMemberResponse = APIResponse[bool]

	// UnbanChatMemberResponse represents the response from unbanChatMember API
	UnbanChatMemberResponse = APIResponse[bool]

	// GetChatMemberResponse represents the response from getChatMember API
	GetChatMemberResponse = APIResponse[TelegramChatMember]

	// GetChatMembersCountResponse represents the response from getChatMembersCount API
	GetChatMembersCountResponse = APIResponse[int]

	// AnswerCallbackQueryResponse represents the response from answerCallbackQuery API
	AnswerCallbackQueryResponse = APIResponse[bool]

	// AnswerInlineQueryResponse represents the response from answerInlineQuery API
	AnswerInlineQueryResponse = APIResponse[bool]
)

// WebhookInfo contains information about the current status of a webhook
type WebhookInfo struct {
	URL                          string   `json:"url"`
	HasCustomCertificate         bool     `json:"has_custom_certificate"`
	PendingUpdateCount           int      `json:"pending_update_count"`
	IPAddress                    *string  `json:"ip_address,omitempty"`
	LastErrorDate                *int64   `json:"last_error_date,omitempty"`
	LastErrorMessage             *string  `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate *int64   `json:"last_synchronization_error_date,omitempty"`
	MaxConnections               *int     `json:"max_connections,omitempty"`
	AllowedUpdates               []string `json:"allowed_updates,omitempty"`
}

// TelegramFile represents a file ready to be downloaded
type TelegramFile struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	FileSize     *int    `json:"file_size,omitempty"`
	FilePath     *string `json:"file_path,omitempty"`
}

// TelegramError represents a structured error from Telegram API
type TelegramError struct {
	Code        int             `json:"code"`
	Description string          `json:"description"`
	Parameters  *ResponseParams `json:"parameters,omitempty"`
	Method      string          `json:"method,omitempty"`
	RequestData map[string]any  `json:"request_data,omitempty"`
	Timestamp   time.Time       `json:"timestamp"`
	Retryable   bool            `json:"retryable"`
}

// Error implements the error interface
func (e *TelegramError) Error() string {
	return e.Description
}

// IsRetryable returns whether the error can be retried
func (e *TelegramError) IsRetryable() bool {
	return e.Retryable
}

// GetRetryAfter returns the retry delay in seconds if available
func (e *TelegramError) GetRetryAfter() *int {
	if e.Parameters != nil {
		return e.Parameters.RetryAfter
	}
	return nil
}

// ResponseMetadata contains metadata about the API response
type ResponseMetadata struct {
	RequestID  string        `json:"request_id,omitempty"`
	Timestamp  time.Time     `json:"timestamp"`
	Duration   time.Duration `json:"duration"`
	Method     string        `json:"method"`
	StatusCode int           `json:"status_code"`
	Attempt    int           `json:"attempt,omitempty"`
	Cached     bool          `json:"cached,omitempty"`
	RateLimit  *RateLimit    `json:"rate_limit,omitempty"`
}

// RateLimit contains rate limiting information
type RateLimit struct {
	Limit     int       `json:"limit"`
	Remaining int       `json:"remaining"`
	Reset     time.Time `json:"reset"`
}

// EnhancedResponse wraps the API response with additional metadata
type EnhancedResponse[T any] struct {
	APIResponse[T]
	Metadata *ResponseMetadata `json:"metadata,omitempty"`
}

// Response validation methods
func (r *BaseResponse) IsSuccess() bool {
	return r.OK
}

func (r *BaseResponse) HasError() bool {
	return !r.OK || r.ErrorCode != nil
}

func (r *BaseResponse) GetErrorCode() int {
	if r.ErrorCode != nil {
		return *r.ErrorCode
	}
	return 0
}

func (r *BaseResponse) GetErrorDescription() string {
	if r.Description != nil {
		return *r.Description
	}
	return ""
}

// IsRateLimited checks if the error is due to rate limiting
func (r *BaseResponse) IsRateLimited() bool {
	return r.ErrorCode != nil && *r.ErrorCode == 429
}

// IsBotBlocked checks if the bot was blocked by user
func (r *BaseResponse) IsBotBlocked() bool {
	return r.ErrorCode != nil && *r.ErrorCode == 403 &&
		r.Parameters != nil && r.Parameters.BotBlocked != nil && *r.Parameters.BotBlocked
}

// IsChatNotFound checks if the chat was not found
func (r *BaseResponse) IsChatNotFound() bool {
	return r.ErrorCode != nil && *r.ErrorCode == 400 &&
		r.Parameters != nil && r.Parameters.ChatNotFound != nil && *r.Parameters.ChatNotFound
}

// IsNetworkError checks if it's a network-related error
func (r *BaseResponse) IsNetworkError() bool {
	return r.Parameters != nil && r.Parameters.NetworkError != nil && *r.Parameters.NetworkError
}

// IsInternalError checks if it's an internal API error
func (r *BaseResponse) IsInternalError() bool {
	return r.Parameters != nil && r.Parameters.InternalAPIError != nil && *r.Parameters.InternalAPIError
}

// ShouldRetry determines if the request should be retried based on the error
func (r *BaseResponse) ShouldRetry() bool {
	if r.IsSuccess() {
		return false
	}

	// Retry on rate limit, internal errors, or network errors
	return r.IsRateLimited() || r.IsInternalError() || r.IsNetworkError()
}

// GetRetryDelay returns the recommended retry delay
func (r *BaseResponse) GetRetryDelay() time.Duration {
	if r.Parameters != nil && r.Parameters.RetryAfter != nil {
		return time.Duration(*r.Parameters.RetryAfter) * time.Second
	}

	// Default retry delays based on error type
	if r.IsRateLimited() {
		return 60 * time.Second // 1 minute for rate limits
	}
	if r.IsInternalError() {
		return 5 * time.Second // 5 seconds for internal errors
	}
	if r.IsNetworkError() {
		return 2 * time.Second // 2 seconds for network errors
	}

	return 0
}

// Helper functions for response creation
func NewSuccessResponse[T any](result *T) *APIResponse[T] {
	return &APIResponse[T]{
		BaseResponse: BaseResponse{OK: true},
		Result:       result,
	}
}

func NewErrorResponse[T any](code int, description string) *APIResponse[T] {
	return &APIResponse[T]{
		BaseResponse: BaseResponse{
			OK:          false,
			ErrorCode:   &code,
			Description: &description,
		},
	}
}

func NewTelegramError(code int, description string, method string, requestData map[string]any) *TelegramError {
	return &TelegramError{
		Code:        code,
		Description: description,
		Method:      method,
		RequestData: requestData,
		Timestamp:   time.Now(),
		Retryable:   shouldRetryError(code),
	}
}

// shouldRetryError determines if an error code indicates a retryable error
func shouldRetryError(code int) bool {
	retryableCodes := map[int]bool{
		429: true, // Too Many Requests
		500: true, // Internal Server Error
		502: true, // Bad Gateway
		503: true, // Service Unavailable
		504: true, // Gateway Timeout
	}
	return retryableCodes[code]
}

// Response builders for common patterns
type ResponseBuilder[T any] struct {
	response *APIResponse[T]
}

func NewResponseBuilder[T any]() *ResponseBuilder[T] {
	return &ResponseBuilder[T]{
		response: &APIResponse[T]{},
	}
}

func (b *ResponseBuilder[T]) Success(result *T) *APIResponse[T] {
	b.response.OK = true
	b.response.Result = result
	return b.response
}

func (b *ResponseBuilder[T]) Error(code int, description string) *APIResponse[T] {
	b.response.OK = false
	b.response.ErrorCode = &code
	b.response.Description = &description
	return b.response
}

func (b *ResponseBuilder[T]) WithParameters(params *ResponseParams) *ResponseBuilder[T] {
	b.response.Parameters = params
	return b
}

// ResponseError encapsulates details about an API error response
type ResponseError struct {
	Method      string         `json:"method"`
	RequestData map[string]any `json:"request_data"`
	Response    *BaseResponse  `json:"response"`
	HTTPStatus  int            `json:"http_status"`
	Timestamp   time.Time      `json:"timestamp"`
	Duration    time.Duration  `json:"duration"`
	Retries     int            `json:"retries"`
}

// Error implements the error interface
func (e *ResponseError) Error() string {
	if e.Response != nil && e.Response.HasError() {
		return e.Response.GetErrorDescription()
	}
	return "Unknown API error"
}

// ShouldRetry indicates if the request should be retried based on the response
func (e *ResponseError) ShouldRetry() bool {
	if e.Response != nil {
		return e.Response.ShouldRetry()
	}
	return false
}

// GetRetryDelay returns the recommended delay before retrying the request
func (e *ResponseError) GetRetryDelay() time.Duration {
	if e.Response != nil {
		return e.Response.GetRetryDelay()
	}
	return 0
}
