package types

type MessageType string

const (
	MessageTypeText    MessageType = "text"
	MessageTypeCommand MessageType = "command"
	MessageTypeError   MessageType = "error"
	MessageTypeInfo    MessageType = "info"
)

// MessageType defines the type for message types
var validMessageTypes = map[MessageType]struct{}{
	MessageTypeText:    {},
	MessageTypeCommand: {},
	MessageTypeError:   {},
	MessageTypeInfo:    {},
}

// IsValid checks if the MessageType is valid
func (mt MessageType) IsValid() bool {
	_, ok := validMessageTypes[mt]
	return ok
}
