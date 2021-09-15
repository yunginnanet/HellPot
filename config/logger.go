package config

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var (
	logFile        *os.File
	logDir         string
	CurrentLogFile string
)

// Logger retrieves an instance of our zerolog loggger so we can hook it in our main package.
func Logger() zerolog.Logger {
	logDir = snek.GetString("logger.directory")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		println("cannot create log directory: " + logDir + "(" + err.Error() + ")")
		os.Exit(1)
	}

	tnow := "HellPot"

	if snek.GetBool("logger.use_date_filename") {
		tnow = strings.Replace(time.Now().Format(time.RFC822), " ", "_", -1)
		tnow = strings.Replace(tnow, ":", "-", -1)
	}

	CurrentLogFile = logDir + tnow + ".log"

	if logFile, err = os.OpenFile(CurrentLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		println("cannot create log file: " + err.Error())
		os.Exit(1)
	}

	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{NoColor: NoColor, Out: os.Stdout}, logFile)
	return zerolog.New(multi).With().Timestamp().Logger()
}
