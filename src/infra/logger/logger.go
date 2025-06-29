package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

var (
	log *Logger
)

func Init(environment string) {
	// Configure logger based on environment
	var config zap.Config
	if environment == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapLogger, err := config.Build()
	if err != nil {
		zap.NewExample().Sugar().Fatalf("Failed to initialize logger: %v", err)
		os.Exit(1)
	}

	log = &Logger{zapLogger.Sugar()}
}

func Get() *Logger {
	if log == nil {
		zapLogger, _ := zap.NewDevelopment()
		log = &Logger{zapLogger.Sugar()}
	}
	return log
}

func Info(args ...interface{}) {
	Get().Info(args...)
}

// Infof logs at INFO level with formatting
func Infof(template string, args ...interface{}) {
	Get().Infof(template, args...)
}

func Debug(args ...interface{}) {
	Get().Debug(args...)
}

// Debugf logs at DEBUG level with formatting
func Debugf(template string, args ...interface{}) {
	Get().Debugf(template, args...)
}

func Warn(args ...interface{}) {
	Get().Warn(args...)
}

// Warnf logs at WARN level with formatting
func Warnf(template string, args ...interface{}) {
	Get().Warnf(template, args...)
}

func Error(args ...interface{}) {
	Get().Error(args...)
}

// Errorf logs at ERROR level with formatting
func Errorf(template string, args ...interface{}) {
	Get().Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	Get().Fatal(args...)
}

// Fatalf logs at FATAL level with formatting and then calls os.Exit(1)
func Fatalf(template string, args ...interface{}) {
	Get().Fatalf(template, args...)
}

// With returns a logger with the given key-value pairs
func With(args ...interface{}) *Logger {
	return &Logger{Get().With(args...)}
}
