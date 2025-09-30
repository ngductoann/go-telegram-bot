package service

import "context"

// Logger defines the interface for logging operations
type Logger interface {
	// Debug logs debug level messages
	Debug(msg string, fields ...any)

	// Info logs informational messages
	Info(msg string, fields ...any)

	// Warn logs warning messages
	Warn(msg string, fields ...any)

	// Error logs error messages
	Error(msg string, fields ...any)

	// Fatal logs fatal error messages and exits
	Fatal(msg string, fields ...any)

	// Panic logs panic messages and panics
	Panic(msg string, fields ...any)

	// WithContext returns a logger with context
	WithContext(ctx context.Context) Logger

	// WithField returns a logger with a single field
	WithField(key string, value any) Logger

	// WithFields returns a logger with multiple fields
	WithFields(fields map[string]any) Logger
}
