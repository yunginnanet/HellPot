package config

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	logFile *os.File
	logDir string
	logger zerolog.Logger
)

// StartLogger instantiates an instance of our zerolog loggger so we can hook it in our main package.
// While this does return a logger, it should not be used for additional retrievals of the logger. Use GetLogger()
func StartLogger() zerolog.Logger {
	logDir = snek.GetString("logger.directory")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		println("cannot create log directory: " + logDir + "(" + err.Error() + ")")
		os.Exit(1)
	}
	if logFile, err = os.OpenFile(logDir+time.Now().Format(time.RFC3339)+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		println("cannot create log file: " + err.Error())
		os.Exit(1)
	}
	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stderr}, logFile)
	logger = zerolog.New(multi).With().Timestamp().Logger()
	return logger
}

// GetLogger retrieves our global logger object
func GetLogger() zerolog.Logger {
	// future logic here
	return logger
}
