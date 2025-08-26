package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/ngductoann/go-telegram-bot/internal/domain/types"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// UserSession represents a session for a user in the system, tracking their current state and activity.
type UserSession struct {
	BaseEntityUUID

	// Foreign keys
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	ChatID uuid.UUID `json:"chat_id" gorm:"type:uuid;not null;index"`

	// State tracking
	FlowStateJSON datatypes.JSON `json:"-" gorm:"column:flow_state"` // raw JSON for DB
	FlowState     *FlowState     `json:"flow_state,omitempty" gorm:"-"`

	CurrentCommand types.CommandKey `json:"current_command,omitempty" gorm:"type:varchar(100)"`
	CurrentCase    types.Case       `json:"current_case,omitempty" gorm:"type:varchar(100)"`
	CurrentStep    types.Step       `json:"current_step,omitempty" gorm:"type:varchar(100)"`
	ExpiresAt      time.Time        `json:"expires_at" gorm:"not null"`

	// Relationships
	User User `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Chat Chat `json:"chat,omitempty" gorm:"foreignKey:ChatID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// NewUserSession creates a new UserSession instance with the provided user ID, chat ID, and duration until expiration.
func NewUserSession(userID uuid.UUID, chatID uuid.UUID, duration time.Duration) *UserSession {
	// Tạo FlowState mặc định
	fs := &FlowState{
		Extras: make(map[string]interface{}),
	}

	// Marshal sang JSON để đồng bộ với DB
	data, _ := json.Marshal(fs)

	return &UserSession{
		UserID:        userID,
		ChatID:        chatID,
		FlowState:     fs,
		FlowStateJSON: datatypes.JSON(data),
		ExpiresAt:     time.Now().Add(duration),
	}
}

func (us *UserSession) BeforeSave(tx *gorm.DB) (err error) {
	if us.FlowState != nil {
		data, err := json.Marshal(us.FlowState)
		if err != nil {
			return err
		}
		us.FlowStateJSON = datatypes.JSON(data)
	}
	return nil
}

func (us *UserSession) AfterFind(tx *gorm.DB) (err error) {
	if len(us.FlowStateJSON) > 0 {
		var fs FlowState
		if err := json.Unmarshal(us.FlowStateJSON, &fs); err != nil {
			return err
		}
		us.FlowState = &fs
	}
	return nil
}

// IsExpired checks if the session has expired
func (us *UserSession) IsExpired() bool {
	return time.Now().After(us.ExpiresAt)
}

// ExtendSession extends the session expiration time
func (us *UserSession) ExtendSession(duration time.Duration) {
	us.ExpiresAt = time.Now().Add(duration)
	us.UpdatedAt = time.Now()
}

// UpdateFlowState sets command/case/step and keeps FlowState in sync.
func (us *UserSession) UpdateFlowState(command types.CommandKey, case_ types.Case, step types.Step) {
	if us.FlowState == nil {
		us.FlowState = &FlowState{Extras: make(map[string]interface{})}
	}
	us.CurrentCommand = command
	us.CurrentCase = case_
	us.CurrentStep = step

	us.FlowState.Command = command
	us.FlowState.Case = case_
	us.FlowState.Step = step

	us.UpdatedAt = time.Now()
}

// SetFlowExtra sets a custom key-value in FlowState.Extras.
func (us *UserSession) SetFlowExtra(key string, value interface{}) {
	if us.FlowState == nil {
		us.FlowState = &FlowState{Extras: make(map[string]interface{})}
	}
	if us.FlowState.Extras == nil {
		us.FlowState.Extras = make(map[string]interface{})
	}
	us.FlowState.Extras[key] = value
	us.UpdatedAt = time.Now()
}

// GetFlowExtra retrieves a value from FlowState.Extras.
func (us *UserSession) GetFlowExtra(key string) (interface{}, bool) {
	if us.FlowState == nil || us.FlowState.Extras == nil {
		return nil, false
	}
	val, ok := us.FlowState.Extras[key]
	return val, ok
}

// DeleteFlowExtra removes a key from FlowState.Extras.
func (us *UserSession) DeleteFlowExtra(key string) {
	if us.FlowState != nil && us.FlowState.Extras != nil {
		delete(us.FlowState.Extras, key)
		us.UpdatedAt = time.Now()
	}
}
