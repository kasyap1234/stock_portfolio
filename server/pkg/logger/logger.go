package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

func NewLogger(serviceName string, environment string) *Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Str("service", serviceName).Str("environment", environment).Logger()
	// Set log level based on environment
	switch environment {
	case "production":
		logger = logger.Level(zerolog.InfoLevel)
	case "development":
		logger = logger.Level(zerolog.DebugLevel)
	case "test":
		logger = logger.Level(zerolog.ErrorLevel)
	default:
		logger = logger.Level(zerolog.InfoLevel)
	}

	return &Logger{logger: &logger}
}
