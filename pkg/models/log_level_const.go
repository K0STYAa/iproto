package models

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	FATAL   = "FATAL"
	PANIC   = "PANIC"
)

func LogStart() {
	// Get the log level from the environment variable
	logLevelStr := os.Getenv("LOG_LEVEL")

	// Set the log level based on the environment variable
	var logLevel logrus.Level
	switch logLevelStr {
	case DEBUG:
		logLevel = logrus.DebugLevel
	case INFO:
		logLevel = logrus.InfoLevel
	case WARNING:
		logLevel = logrus.WarnLevel
	case ERROR:
		logLevel = logrus.ErrorLevel
	case FATAL:
		logLevel = logrus.FatalLevel
	case PANIC:
		logLevel = logrus.PanicLevel
	default:
		logrus.Fatalf("Invalid log level: %v", logLevelStr)
	}

	// Configure the logger with the selected log level
	logrus.SetLevel(logLevel)

	// Set Custom Formatter
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
}
