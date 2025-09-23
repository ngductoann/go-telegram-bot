package types

type Command string

const (
	CommandStart     Command = "/start"
	CommandHelp      Command = "/help"
	CommandGetHomeIP Command = "/home_ip"
)

var validCommands = map[Command]struct{}{
	CommandStart:     {},
	CommandHelp:      {},
	CommandGetHomeIP: {},
}

func (c Command) IsValid() bool {
	_, ok := validCommands[c]
	return ok
}
