package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

type (
	TelegramUserID int64 // Telegram user ID
	TelegramChatID int64 // Telegram chat ID
)

// CommandKey represents a command key for Telegram bot commands
type CommandKey string

// Value implements the driver.Valuer interface for TelegramUserID
func (tuid TelegramUserID) Value() (driver.Value, error) {
	return int64(tuid), nil
}

// Scan implements the sql.Scanner interface for TelegramUserID
func (tuid *TelegramUserID) Scan(value interface{}) error {
	if value == nil {
		*tuid = 0
		return nil
	}

	switch v := value.(type) {
	case int64:
		*tuid = TelegramUserID(v)
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		*tuid = TelegramUserID(i)
	default:
		return fmt.Errorf("cannot scan %T into TelegramUserID", value)
	}

	return nil
}

// Value implements the driver.Valuer interface for TelegramChatID
func (tcid *TelegramChatID) Value() (driver.Value, error) {
	return int64(*tcid), nil
}

// Scan implements the sql.Scanner interface for TelegramChatID
func (tcid *TelegramChatID) Scan(value interface{}) error {
	if value == nil {
		*tcid = 0
		return nil
	}

	switch v := value.(type) {
	case int64:
		*tcid = TelegramChatID(v)
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		*tcid = TelegramChatID(i)
	default:
		return fmt.Errorf("cannot scan %T into TelegramChatID", value)
	}

	return nil
}

// Validate checks if the TelegramUserID is valid (not nil and greater than 0)
func (tuid *TelegramUserID) Validate() bool {
	return tuid != nil && *tuid > 0
}

// Validate checks if the TelegramChatID is valid (not equal to 0)
// Chat to telegram user id greater than 0
// Chat to telegram group id less than 0
func (tcid *TelegramChatID) Validate() bool {
	return *tcid != 0
}
