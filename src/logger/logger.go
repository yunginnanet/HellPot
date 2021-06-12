package logger

import (
	//"fmt"
	"github.com/rs/zerolog"
	//"github.com/rs/zerolog/log"
	"github.com/yunginnanet/HellPot/src/config"
	"os"
	"time"
)

var (
	logFile *os.File
	err     error
)

var GlobalLogger zerolog.Logger

func LogInit() {
	if err := os.MkdirAll(config.LogDir, 0755); err != nil {
		panic("cannot create log directory: " + config.LogDir + "(" + err.Error() + ")")
	}
	if logFile, err = os.OpenFile(config.LogDir+time.Now().Format(time.RFC3339)+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		panic("cannot create log file: " + err.Error())
	}
	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stderr}, logFile)
	GlobalLogger = zerolog.New(multi).With().Timestamp().Logger()
}
