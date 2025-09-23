package types

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSupergroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

var validChatType = map[ChatType]struct{}{
	ChatTypePrivate:    {},
	ChatTypeGroup:      {},
	ChatTypeSupergroup: {},
	ChatTypeChannel:    {},
}

func (ct ChatType) IsValid() bool {
	_, ok := validChatType[ct]
	return ok
}
