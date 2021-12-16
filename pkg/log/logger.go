package log

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger = logrus.New()

// InitLogger creates the logger with level
func InitLogger(level logrus.Level) {
	logger.SetLevel(level)
}

// GetLogger returns the current logger
func GetLogger() *logrus.Logger {
	return logger
}
