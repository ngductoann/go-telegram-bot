package util

import "strings"

// EscapeMarkdownV2 escapes special characters for Telegram MarkdownV2 format
// Reference: https://core.telegram.org/bots/api#markdownv2-style
func EscapeMarkdownV2(text string) string {
	// Define special characters that need escaping in MarkdownV2
	// Characters: _ * [ ] ( ) ~ ` > # + - = | { } . !
	const specialChars = "_*[]()~`>#+-=|{}.!:"

	// Use strings.Builder for efficient string concatenation
	var builder strings.Builder
	builder.Grow(len(text) * 2) // Pre-allocate with estimated capacity

	for _, char := range text {
		// Check if character needs escaping
		if strings.ContainsRune(specialChars, char) {
			builder.WriteRune('\\')
		}
		builder.WriteRune(char)
	}

	return builder.String()
}
