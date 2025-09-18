package service

import "context"

// Logger defines the interface for logging operations
type Logger interface {
	// Debug logs debug level messages
	Debug(msg string, fields ...interface{})

	// Info logs informational messages
	Info(msg string, fields ...interface{})

	// Warn logs warning messages
	Warn(msg string, fields ...interface{})

	// Error logs error messages
	Error(msg string, fields ...interface{})

	// Fatal logs fatal error messages and exits
	Fatal(msg string, fields ...interface{})

	// Panic logs panic messages and panics
	Panic(msg string, fields ...interface{})

	// WithContext returns a logger with context
	WithContext(ctx context.Context) Logger

	// WithField returns a logger with a single field
	WithField(key string, value interface{}) Logger

	// WithFields returns a logger with multiple fields
	WithFields(fields map[string]interface{}) Logger
}
