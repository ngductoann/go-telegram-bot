package logger

import (
	"context"
	"fmt"
	domainService "go-telegram-bot/internal/domain/service"
	"go-telegram-bot/internal/infrastructure/config"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// ZapLogger implements domainService.Logger interface using Uber's Zap library.
type ZapLogger struct {
	logger *zap.Logger
}

// NewZapLogger creates a new instance of ZapLogger based on the provided configuration.
func NewZapLogger(cfg *config.Config) (*ZapLogger, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	// Set log level
	if cfg.Logger.LogLevel == "" {
		cfg.Logger.LogLevel = "info"
	}
	zapLevel, err := getLogLevel(cfg)
	if err != nil {
		return nil, err
	}

	// Configure the encoder and output
	encoder := getEncoderLogger()

	// Configure lumberjack for log rotation
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Logger.FilePath,
		MaxSize:    cfg.Logger.MaxSize, // megabytes
		MaxBackups: cfg.Logger.MaxBackups,
		MaxAge:     cfg.Logger.MaxAge, // days
		Compress:   cfg.Logger.Compress,
	})

	// MultiWriteSyncer to write to both file and console
	consoleWriter := zapcore.AddSync(zapcore.Lock(os.Stdout)) // Console output

	write := zapcore.NewMultiWriteSyncer(consoleWriter, fileWriter) // Write to both console and file

	// Create a sampler to limit log volume for high-frequency logs
	core := zapcore.NewSamplerWithOptions(
		zapcore.NewCore(encoder, write, zapLevel),
		time.Second,
		100,
		1,
	)

	// Build the logger
	logger := zap.New(
		core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return &ZapLogger{logger}, nil
}

// getLogLevel maps string log levels from config to zapcore.Level.
func getLogLevel(cfg *config.Config) (*zapcore.Level, error) {
	var zapLevel zapcore.Level

	switch cfg.Logger.LogLevel {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warn":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	case "dpanic":
		zapLevel = zap.DPanicLevel
	case "panic":
		zapLevel = zap.PanicLevel
	case "fatal":
		zapLevel = zap.FatalLevel
	default:
		return nil, fmt.Errorf("invalid log level: %s", cfg.Logger.LogLevel)
	}
	return &zapLevel, nil
}

// getEncoderLogger configures the log encoder format.
func getEncoderLogger() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (z *ZapLogger) Debug(msg string, fields ...any) {
	zapFields := z.convertFields(fields...)
	z.logger.Debug(msg, zapFields...)
}

func (z *ZapLogger) Info(msg string, fields ...any) {
	zapFields := z.convertFields(fields...)
	z.logger.Info(msg, zapFields...)
}

func (z *ZapLogger) Warn(msg string, fields ...any) {
	zapFields := z.convertFields(fields...)
	z.logger.Warn(msg, zapFields...)
}

func (z *ZapLogger) Error(msg string, fields ...any) {
	zapFields := z.convertFields(fields...)
	z.logger.Error(msg, zapFields...)
}

func (z *ZapLogger) Fatal(msg string, fields ...any) {
	zapFields := z.convertFields(fields...)
	z.logger.Fatal(msg, zapFields...)
}

func (z *ZapLogger) Panic(msg string, fields ...any) {
	zapFields := z.convertFields(fields...)
	z.logger.Panic(msg, zapFields...)
}

func (z *ZapLogger) convertFields(fields ...any) []zap.Field {
	if len(fields)%2 != 0 {
		// if odd number of fields, add the last one as a generic field
		fields = append(fields, "unknown")
	}

	// Convert to zap.Fields assuming fields are in key-value pairs
	zapFields := make([]zap.Field, 0, len(fields)/2)
	for i := 0; i < len(fields); i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue
		}
		value := fields[i+1]
		zapFields = append(zapFields, zap.Any(key, value))
	}

	return zapFields
}

// WithContext returns a logger with context (for now, returns the same logger as zap doesn't directly support context)
func (z *ZapLogger) WithContext(ctx context.Context) domainService.Logger {
	// For now, we return the same logger. In a more advanced implementation,
	// we could extract trace/span information from context
	return z
}

// WithField returns a logger with a single field
func (z *ZapLogger) WithField(key string, value any) domainService.Logger {
	newLogger := z.logger.With(zap.Any(key, value))
	return &ZapLogger{logger: newLogger}
}

// WithFields returns a logger with multiple fields
func (z *ZapLogger) WithFields(fields map[string]any) domainService.Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}
	newLogger := z.logger.With(zapFields...)
	return &ZapLogger{logger: newLogger}
}
