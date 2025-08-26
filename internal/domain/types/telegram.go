package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// TelegramUserID represents a unique identifier for a Telegram user.
type TelegramUserID int64

// TelegramChatID represents a unique identifier for a Telegram chat.
type TelegramChatID int64

// CommandKey represents a command key used in Telegram bot commands.
type CommandKey string

// Case represents a case type used in various contexts.
type Case string

// Step represents a step in a process or workflow.
type Step int

func (tuid TelegramUserID) value() (driver.Value, error) {
	return int64(tuid), nil
}

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
			return err
		}
		*tuid = TelegramUserID(i)
	default:
		return fmt.Errorf("cannot scan %T into TelegramUserID", value)
	}

	return nil
}

// Value implements the driver.Valuer interface for TelegramChatID
func (tcid TelegramChatID) Value() (driver.Value, error) {
	return int64(tcid), nil
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

// Validate checks if the TelegramUserID is valid (non-zero)
func (tuid *TelegramUserID) Validate() bool {
	return *tuid > 0
}

func (tcid *TelegramChatID) Validate() bool {
	return *tcid != 0
}

// Common constants
const (
	DefaultTimeout = 300 // 5 minutes default timeout for flow steps
)

// String returns string representation of CommandKey
func (ck CommandKey) String() string {
	return string(ck)
}

// String returns string representation of Case
func (c Case) String() string {
	return string(c)
}

// Int returns int representation of Step
func (s Step) Int() int {
	return int(s)
}
