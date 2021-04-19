package logger

import (
	"github.com/sirupsen/logrus"
)

var level logrus.Level = logrus.WarnLevel

type nameHook struct {
	Name string
}

// Levels returns the levels handled by the hook
func (n nameHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.PanicLevel,
		logrus.FatalLevel,
	}
}

// Fire runs the hook for the entry
func (n nameHook) Fire(e *logrus.Entry) error {
	e.Data["name"] = n.Name
	return nil
}

// SetLogLevel sets the log l	evel to the given level
func SetLogLevel(newLevel logrus.Level) {
	level = newLevel
}

// GetLogger creates and returns a new logger instance with a name
func GetLogger(name string) *logrus.Logger {
	// Create a new instance of the logger. You can have any number of instances.
	var logger = logrus.New()

	// The TextFormatter is default, you don't actually have to do this.
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.AddHook(nameHook{
		Name: name,
	})
	return logger
}
