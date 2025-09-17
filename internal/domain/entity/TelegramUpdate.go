package entity

type TelegramUpdate struct {
	UpdateID int64 `json:"update_id"`
	Message  *struct {
		MessageID int64 `json:"message_id"`
		From      *struct {
			ID        int64  `json:"id"`
			Username  string `json:"username"`
			FirstName string `json:"first_name"`
		} `json:"from"`
		Chat *struct {
			ID   int64  `json:"id"`
			Type string `json:"type"`
		} `json:chat"`
		Date int64  `json:"date"`
		Text string `json:"text"`
	} `json:"message,omitempty"`
}
