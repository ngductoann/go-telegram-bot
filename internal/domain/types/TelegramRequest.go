package types

// Request structures for enhanced methods
type SendMessageRequest struct {
	ChatID                TelegramChatID `json:"chat_id"`
	Text                  string         `json:"text"`
	ParseMode             *ParseMode     `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool          `json:"disable_web_page_preview,omitempty"`
	DisableNotification   *bool          `json:"disable_notification,omitempty"`
	ReplyToMessageID      *int64         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           any            `json:"reply_markup,omitempty"`
}

type GetUpdatesRequest struct {
	Offset         int64    `json:"offset,omitempty"`
	Limit          *int     `json:"limit,omitempty"`
	Timeout        *int     `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type SetWebhookRequest struct {
	URL                string   `json:"url"`
	Certificate        []byte   `json:"certificate,omitempty"`
	IPAddress          *string  `json:"ip_address,omitempty"`
	MaxConnections     *int     `json:"max_connections,omitempty"`
	AllowedUpdates     []string `json:"allowed_updates,omitempty"`
	DropPendingUpdates *bool    `json:"drop_pending_updates,omitempty"`
	SecretToken        *string  `json:"secret_token,omitempty"`
}

type EditMessageTextRequest struct {
	ChatID                *TelegramChatID `json:"chat_id,omitempty"`
	MessageID             *int64          `json:"message_id,omitempty"`
	InlineMessageID       *string         `json:"inline_message_id,omitempty"`
	Text                  string          `json:"text"`
	ParseMode             *ParseMode      `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool           `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           any             `json:"reply_markup,omitempty"`
}

type ForwardMessageRequest struct {
	ChatID              TelegramChatID `json:"chat_id"`
	FromChatID          TelegramChatID `json:"from_chat_id"`
	MessageID           int64          `json:"message_id"`
	DisableNotification *bool          `json:"disable_notification,omitempty"`
}

type SendPhotoRequest struct {
	ChatID              TelegramChatID `json:"chat_id"`
	Photo               any            `json:"photo"` // InputFile or string
	Caption             *string        `json:"caption,omitempty"`
	ParseMode           *ParseMode     `json:"parse_mode,omitempty"`
	DisableNotification *bool          `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int64         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         any            `json:"reply_markup,omitempty"`
}

type SendDocumentRequest struct {
	ChatID                      TelegramChatID `json:"chat_id"`
	Document                    any            `json:"document"` // InputFile or string
	Thumbnail                   any            `json:"thumbnail,omitempty"`
	Caption                     *string        `json:"caption,omitempty"`
	ParseMode                   *ParseMode     `json:"parse_mode,omitempty"`
	DisableContentTypeDetection *bool          `json:"disable_content_type_detection,omitempty"`
	DisableNotification         *bool          `json:"disable_notification,omitempty"`
	ReplyToMessageID            *int64         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup                 any            `json:"reply_markup,omitempty"`
}

type BanChatMemberRequest struct {
	ChatID         TelegramChatID `json:"chat_id"`
	UserID         TelegramUserID `json:"user_id"`
	UntilDate      *int64         `json:"until_date,omitempty"`
	RevokeMessages *bool          `json:"revoke_messages,omitempty"`
}

type UnbanChatMemberRequest struct {
	ChatID       TelegramChatID `json:"chat_id"`
	UserID       TelegramUserID `json:"user_id"`
	OnlyIfBanned *bool          `json:"only_if_banned,omitempty"`
}

type AnswerCallbackQueryRequest struct {
	CallbackQueryID string  `json:"callback_query_id"`
	Text            *string `json:"text,omitempty"`
	ShowAlert       *bool   `json:"show_alert,omitempty"`
	URL             *string `json:"url,omitempty"`
	CacheTime       *int    `json:"cache_time,omitempty"`
}

type AnswerInlineQueryRequest struct {
	InlineQueryID     string              `json:"inline_query_id"`
	Results           []InlineQueryResult `json:"results"`
	CacheTime         *int                `json:"cache_time,omitempty"`
	IsPersonal        *bool               `json:"is_personal,omitempty"`
	NextOffset        *string             `json:"next_offset,omitempty"`
	SwitchPmText      *string             `json:"switch_pm_text,omitempty"`
	SwitchPmParameter *string             `json:"switch_pm_parameter,omitempty"`
}

// InlineQueryResult represents one result of an inline query
type InlineQueryResult interface {
	GetType() string
}
