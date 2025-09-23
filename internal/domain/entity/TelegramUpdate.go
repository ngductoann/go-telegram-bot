package entity

// TelegramUpdate represents a Telegram bot update message
type TelegramUpdate struct {
	UpdateID int64    `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}
