package log

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// InitLogger creates the logger with level
func TestInitLogger(t *testing.T) {
	InitLogger(logrus.DebugLevel)
	assert.Equal(t, logrus.DebugLevel, logger.Level)
}

// GetLogger returns the current logger
func TestGetLogger(t *testing.T) {
	logger = logrus.New()
	assert.NotNil(t, GetLogger())
	assert.Equal(t, logrus.InfoLevel, GetLogger().Level)
}
