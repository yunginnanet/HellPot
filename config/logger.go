package config

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var (
	// CurrentLogFile is used for accessing the location of the currently used log file across packages.
	CurrentLogFile string
	logFile        *os.File
	logDir         string
	logger         zerolog.Logger
)

// StartLogger instantiates an instance of our zerolog loggger so we can hook it in our main package.
// While this does return a logger, it should not be used for additional retrievals of the logger. Use GetLogger()
func StartLogger() zerolog.Logger {
	logDir = snek.GetString("logger.directory")
	if !strings.HasSuffix(logDir, "/") {
		logDir += "/"
	}
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		println("cannot create log directory: " + logDir + "(" + err.Error() + ")")
		os.Exit(1)
	}

	logFileName := "HellPot"

	if snek.GetBool("logger.use_date_filename") {
		tn := strings.ReplaceAll(time.Now().Format(time.RFC822), " ", "_")
		tn = strings.ReplaceAll(logFileName, ":", "-")
		logFileName = logFileName + "_" + tn
	}

	CurrentLogFile = logDir + logFileName + ".log"

	/* #nosec */
	if logFile, err = os.OpenFile(CurrentLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666); err != nil {
		println("cannot create log file: " + err.Error())
		os.Exit(1)
	}

	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{NoColor: NoColor, Out: os.Stdout}, logFile)
	logger = zerolog.New(multi).With().Timestamp().Logger()
	return logger
}

// GetLogger retrieves our global logger object
func GetLogger() *zerolog.Logger {
	// future logic here
	return &logger
}
