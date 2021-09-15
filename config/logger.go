package config

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	logFile *os.File
	logDir  string
)
// Logger retrieves an instance of our zerolog loggger so we can hook it in our main package.
func Logger() zerolog.Logger {
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
	return zerolog.New(multi).With().Timestamp().Logger()
}
