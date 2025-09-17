package entity

// BotMessage represents a message sent by a bot in a chat application.
type BotMessage struct {
	ChatID  int64  `json:"chat_id"`
	Text    string `json:"text"`
	SentAt  int64  `json:"sent_at"`           // Unix timestamp
	Command string `json:"command,omitempty"` // Optional command associated with the message
}
