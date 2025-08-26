package entity

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
)

// CallbackData represents the data structure for handling callback queries in Telegram bots.
type CallbackData struct {
	CommandKey types.CommandKey `json:"command_key" gorm:"type:varchar(100);not null"`
	Case       types.Case       `json:"case" gorm:"type:varchar(100)"`
	Step       types.Step       `json:"step" gorm:"type:varchar(100)"`
	Payload    string           `json:"payload,omitempty" gorm:"type:text"`
}

// ToJSON converts CallbackData to JSON string
// ToJSON creates a compact string representation for callback data
// Format: "cmd:case:step:payload" (using : as delimiter for compactness)
func (cd *CallbackData) ToJson() (string, error) {
	// Use a compact format instead of JSON to stay within Telegram's 64-byte limit
	result := string(cd.CommandKey) + ":" + string(cd.Case) + ":" + fmt.Sprintf("%d", cd.Step)
	if cd.Payload != "" {
		result += ":" + cd.Payload
	}

	// Check if it exceeds Telegram's callback data limit (64 bytes)
	if len(result) > 64 {
		return "", fmt.Errorf("callback data too long: %d bytes (max 64)", len(result))
	}

	return result, nil
}

// FromJSON creates CallbackData from compact string
func CallbackDataFromJSON(data string) (*CallbackData, error) {
	parts := strings.Split(data, ":")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid callback data format")
	}

	step, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("invalid step value: %v", err)
	}

	cd := &CallbackData{
		CommandKey: types.CommandKey(parts[0]),
		Case:       types.Case(parts[1]),
		Step:       types.Step(step),
	}

	if len(parts) > 3 {
		cd.Payload = parts[3]
	}

	return cd, nil
}
