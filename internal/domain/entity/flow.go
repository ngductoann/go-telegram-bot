package entity

import "github.com/ngductoann/go-telegram-bot/internal/domain/types"

// FlowButton represents a button in a Telegram bot's flow, which can either trigger a callback or open a URL.
type FlowButton struct {
	Text         string        `json:"text" gorm:"type:varchar(255);not null"`
	CallbackData *CallbackData `json:"callback_data,omitempty" gorm:"-"`
	URL          *string       `json:"url,omitempty" gorm:"-"`
}

// FlowState represents the dynamic state of a user session.
type FlowState struct {
	Command types.CommandKey       `json:"command,omitempty"`
	Case    types.Case             `json:"case,omitempty"`
	Step    types.Step             `json:"step,omitempty"`
	Extras  map[string]interface{} `json:"extras,omitempty"` // flexible payload
}
