package config

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var (
	// CurrentLogFile is used for accessing the location of the currently used log file across packages.
	CurrentLogFile string
	logFile        io.Writer
	logDir         string
	logger         zerolog.Logger
)

func prepLogDir() {
	logDir = snek.String("logger.directory")
	if logDir == "" {
		logDir = filepath.Join(home, ".local", "share", Title, "logs")
	}
	_ = os.MkdirAll(logDir, 0666)
}

// StartLogger instantiates an instance of our zerolog loggger so we can hook it in our main package.
// While this does return a logger, it should not be used for additional retrievals of the logger. Use GetLogger().
func StartLogger(pretty bool, targets ...io.Writer) zerolog.Logger {
	logFileName := "HellPot"

	if snek.Bool("logger.use_date_filename") {
		tn := strings.ReplaceAll(time.Now().Format(time.RFC822), " ", "_")
		tn = strings.ReplaceAll(tn, ":", "-")
		logFileName = logFileName + "_" + tn
	}

	var err error

	switch {
	case len(targets) > 0:
		logFile = io.MultiWriter(targets...)
	default:
		prepLogDir()
		CurrentLogFile = path.Join(logDir, logFileName+".log")
		//nolint:lll
		logFile, err = os.OpenFile(CurrentLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666) // #nosec G304 G302 -- we are not using user input to create the file
		if err != nil {
			println("cannot create log file: " + err.Error())
			os.Exit(1)
		}
	}

	var logWriter = logFile

	if pretty {
		logWriter = zerolog.MultiLevelWriter(zerolog.ConsoleWriter{TimeFormat: ConsoleTimeFormat, NoColor: NoColor, Out: os.Stdout}, logFile)
	}

	logger = zerolog.New(logWriter).With().Timestamp().Logger()
	return logger
}

// GetLogger retrieves our global logger object.
func GetLogger() *zerolog.Logger {
	// future logic here
	return &logger
}
