package logger

import (
	"os"
	"time"

	"github.com/ngductoann/go-telegram-bot/internal/infrastructure/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// ZapLogger implements the Logger interface using zap
type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(cfg *config.Config) (*ZapLogger, error) {
	var zapLevel zapcore.Level
	// Level: debug -> info -> warn -> error -> fatal
	switch cfg.Log.LogLevel {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "fatal":
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Log.Filepath,
		MaxSize:    cfg.Log.MaxSize,    // MB
		MaxBackups: cfg.Log.MaxBackups, // số file backup
		MaxAge:     cfg.Log.MaxAge,     // ngày
		Compress:   cfg.Log.Compress,   // gzip backup
	})

	// Multi output (stdout + file)
	consoleWriter := zapcore.AddSync(os.Stdout)

	writer := zapcore.NewMultiWriteSyncer(consoleWriter, fileWriter)

	// Core with sampling
	core := zapcore.NewSamplerWithOptions(
		zapcore.NewCore(encoder, writer, zapLevel),
		time.Second, // window
		100,         // log đầy đủ 100 dòng đầu
		1,           // sau đó mỗi loại log chỉ in 1 dòng / giây
	)

	// Build logger
	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel), // if error is ErrorLevel or higher, add stacktrace
	)

	return &ZapLogger{logger}, nil
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func (z *ZapLogger) Debug(msg string, fields ...interface{}) {
	zapFields := z.convertFields(fields...)
	z.logger.Debug(msg, zapFields...)
}

func (z *ZapLogger) Info(msg string, fields ...interface{}) {
	zapFields := z.convertFields(fields...)
	z.logger.Info(msg, zapFields...)
}

func (z *ZapLogger) Warn(msg string, fields ...interface{}) {
	zapFields := z.convertFields(fields...)
	z.logger.Warn(msg, zapFields...)
}

func (z *ZapLogger) Error(msg string, fields ...interface{}) {
	zapFields := z.convertFields(fields...)
	z.logger.Error(msg, zapFields...)
}

func (z *ZapLogger) Fatal(msg string, fields ...interface{}) {
	zapFields := z.convertFields(fields...)
	z.logger.Fatal(msg, zapFields...)
}

func (z *ZapLogger) Panic(msg string, fields ...interface{}) {
	zapFields := z.convertFields(fields...)
	z.logger.Panic(msg, zapFields...)
}

func (z *ZapLogger) CheckError(err error, msg string, fatal bool) {
	if err == nil {
		return
	}

	if fatal {
		z.Fatal(msg, "error", err)
	} else {
		z.Fatal(msg, "errro", err)
		panic(err)
	}
}

func (z *ZapLogger) CheckErrorFatal(err error, msg string) {
	z.CheckError(err, msg, true)
}

func (z *ZapLogger) CheckErrorPanic(err error, msg string) {
	z.CheckError(err, msg, false)
}

func (z *ZapLogger) convertFields(fields ...interface{}) []zap.Field {
	if len(fields)%2 != 0 {
		// if odd number of fields, add the last one as a generic field
		fields = append(fields, "unknown")
	}

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
