package types

type ParseMode string

const (
	ParseModeMarkdown   ParseMode = "Markdown"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	ParseModeHTML       ParseMode = "HTML"
	ParseModeNone       ParseMode = "" // nil/empty
)

var validParseModes = map[ParseMode]struct{}{
	ParseModeMarkdown:   {},
	ParseModeMarkdownV2: {},
	ParseModeHTML:       {},
	ParseModeNone:       {},
}

func (pm ParseMode) IsValid() bool {
	_, ok := validParseModes[pm]
	return ok
}
