package logger

import (
	"errors"
	"io"
	"os"
	"sync"

	"github.com/rs/zerolog"
)

// Configuration represents the configuration for the logger.
type Configuration struct {
	Directory     string `koanf:"directory"`
	Debug         bool   `koanf:"debug"`
	Trace         bool   `koanf:"trace"`
	NoColor       bool   `koanf:"nocolor"`
	DockerLogging bool   `koanf:"docker_logging"`
	// ConsoleTimeFormat sets the time format for the console.
	// The string is passed to time.Format() down the line.
	ConsoleTimeFormat string

	Outputs []io.Writer `koanf:"-"`
}

var once = &sync.Once{}

func GetLoggerOnce() *zerolog.Logger {
	var ret *zerolog.Logger
	once.Do(func() {
		ret = &_log
	})
	if ret == nil {
		panic("i said once you fool")
	}
	return ret
}

var ErrNoOutputs = errors.New("no outputs provided")

var _log zerolog.Logger

func New(conf *Configuration) (zerolog.Logger, error) {
	if len(conf.Outputs) == 0 {
		return zerolog.Logger{}, ErrNoOutputs
	}
	for i, output := range conf.Outputs {
		if output == os.Stdout || output == os.Stderr {
			cw := zerolog.ConsoleWriter{Out: output, TimeFormat: conf.ConsoleTimeFormat, NoColor: conf.NoColor}
			conf.Outputs = append(conf.Outputs[:i], conf.Outputs[i+1:]...)
			conf.Outputs = append(conf.Outputs, cw)
		}
	}
	_log = zerolog.New(zerolog.MultiLevelWriter(conf.Outputs...)).With().Timestamp().Logger()
	_log = _log.Level(zerolog.InfoLevel)
	if conf.Debug {
		_log = _log.Level(zerolog.DebugLevel)
	}
	if conf.Trace {
		_log = _log.Level(zerolog.TraceLevel)
	}
	return _log, nil
}
