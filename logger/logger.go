package logger

import (
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogger *zap.SugaredLogger
	once          sync.Once
	level         zap.AtomicLevel
)

func init() {
	level = zap.NewAtomicLevelAt(toZapLevel(ErrorLevel))
}

type Level uint32

// These are the different logging levels.
const (
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = iota
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

//SetLevel sets the log level to filter logs
func SetLevel(l Level) {
	level.SetLevel(toZapLevel(l))
}

func LevelFromString(str string) Level {
	str = strings.ToUpper(str)
	switch str {
	case "ERROR":
		return ErrorLevel
	case "WARN":
		return WarnLevel
	case "INFO":
		return InfoLevel
	default:
		return DebugLevel
	}
}

func getLogger() *zap.SugaredLogger {
	if defaultLogger == nil {
		once.Do(func() {
			zapConfig := zap.NewDevelopmentConfig()
			zapConfig.Level = level
			zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			zapLogger, err := zapConfig.Build(zap.AddCallerSkip(1), zap.AddStacktrace(toZapLevel(ErrorLevel)))
			if err != nil {
				panic(err)
			}

			defaultLogger = zapLogger.Sugar()
		})
	}
	return defaultLogger
}

func toZapLevel(l Level) zapcore.Level {
	switch l {
	case ErrorLevel:
		return zapcore.ErrorLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case InfoLevel:
		return zapcore.InfoLevel
	default:
		return zapcore.DebugLevel
	}
}

// Debug writes out a debug log to global logger
func Debug(msg string, kvs ...interface{}) {
	getLogger().With(kvs...).Debugw(msg)
}

// Info writes out an info log to global logger
func Info(msg string, kvs ...interface{}) {
	getLogger().With(kvs...).Infow(msg)
}

// Warn writes out a warning log to global logger
func Warn(msg string, kvs ...interface{}) {
	getLogger().With(kvs...).Warnw(msg)
}

// Error writes out an error log to global logger
func Error(msg string, kvs ...interface{}) {
	getLogger().With(kvs...).Errorw(msg)
}

// Fatal writes out an error log to global logger
// and then exit
func Fatal(msg string, kvs ...interface{}) {
	getLogger().With(kvs...).Fatalw(msg)

}
