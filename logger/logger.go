package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// init the logger
var logger = logrus.New()

func init() {
	// Set the log level based on the environment variable
	logLevel := os.Getenv("LOG_LEVEL")
	switch strings.ToLower(logLevel) {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	// set log format
	logFormat := os.Getenv("LOG_FORMAT")
	if logFormat == "text" {
		logger.SetFormatter(&logrus.TextFormatter{
			ForceQuote:      true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})
	} else {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     true,
		})
	}

}

// Info logs a message with severity Info
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Debug logs a message with severity Debug
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Warn logs a message with severity Warn
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error logs a message with severity Error
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Error logs a message with severity Fatal
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
