package logger

import (
	"errors"
	"fmt"
	"io"
	"log/syslog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

// Configuration represents the configuration for the logger.
type Configuration struct {
	Debug         bool `koanf:"debug"`
	Trace         bool `koanf:"trace"`
	NoColor       bool `koanf:"nocolor"`
	DockerLogging bool `koanf:"docker_logging"`
	// ConsoleTimeFormat sets the time format for the console.
	// The string is passed to time.Format() down the line.
	ConsoleTimeFormat string

	// TimeDateFilename sets the log file name to include the date and time.
	TimeDateFilename bool `koanf:"use_date_filename"`

	Directory string `koanf:"directory"`
	File      string `koanf:"log_file"`
	RSyslog   string `koanf:"rsyslog_address"`

	ActiveLogFileName string `koanf:"active_log_file_name"`

	Outputs []io.Writer `koanf:"-"`
}

func (c *Configuration) findFallbackDir() error {
	locs := []string{"/var/log"}
	uconf, err := os.UserHomeDir()
	if err == nil {
		locs = append(locs, filepath.Join(uconf, ".local", "share"))
	}

	var errs []error

	for _, loc := range locs {
		if _, err = os.Stat(loc); err == nil {
			var locStat os.FileInfo
			if locStat, err = os.Stat(filepath.Join(loc, "HellPot")); err == nil {
				if locStat.IsDir() {
					c.Directory = filepath.Join(loc, "HellPot")
					c.File = "HellPot.log"
					return nil
				}
				errs = append(errs, fmt.Errorf("HellPot directory exists but is not a directory"))
			} else if !errors.Is(err, os.ErrNotExist) {
				errs = append(errs, err)
			}

			if err = os.MkdirAll(filepath.Join(loc, "HellPot"), 0750); err == nil {
				c.Directory = filepath.Join(loc, "HellPot")
				c.File = "HellPot.log"
				return nil
			}
			errs = append(errs, fmt.Errorf("failed to create HellPot directory"))
		}
	}

	return errors.Join(errs...)
}

func (c *Configuration) Validate() error {
	if c.DockerLogging {
		c.Outputs = []io.Writer{os.Stdout}
	}
	if c.Directory == "" && c.File == "" && c.ActiveLogFileName == "" && c.RSyslog == "" && !c.DockerLogging {
		if err := c.findFallbackDir(); err != nil {
			return fmt.Errorf("failed to find a fallback log directory: %w", err)
		}
	}
	if len(c.Outputs) == 0 && c.Directory == "" && c.File == "" && c.RSyslog == "" {
		return ErrNoOutputs
	}
	if c.File != "" && c.Directory != "" && filepath.Dir(c.File) != c.Directory && filepath.Base(c.File) != c.File {
		return ErrBothFileAndDirectory
	}
	return nil
}

func (c *Configuration) setupDirAndFile() error {
	switch {
	case c.Directory != "":
		println(c.Directory)
		stat, err := os.Stat(c.Directory)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("failed to access specified log directory: %w", err)
		}
		if errors.Is(err, os.ErrNotExist) {
			if err = os.MkdirAll(c.Directory, 0750); err != nil {
				return fmt.Errorf("failed to create specified log directory: %w", err)
			}
		}
		if stat != nil && !stat.IsDir() {
			return fmt.Errorf("specified log directory is not a directory")
		}
		if c.File == "" {
			c.File = "HellPot.log"
		}
	case c.Directory == "" && c.File != "":
		println(c.File)
		stat, err := os.Stat(filepath.Dir(c.File))
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("failed to access specified log directory: %w", err)
		}
		if errors.Is(err, os.ErrNotExist) {
			if err = os.MkdirAll(filepath.Dir(c.File), 0750); err != nil {
				return fmt.Errorf("failed to create specified log directory: %w", err)
			}
		}
		if stat != nil && !stat.IsDir() {
			panic("specified log directory is not a directory, but it should be...? please report this issue on github")
		}
	case c.Directory == "" && c.File == "" && c.ActiveLogFileName == "" && c.RSyslog == "" && !c.DockerLogging:
		return fmt.Errorf("no log directory or file specified")
	}
	var f *os.File
	var err error
	if c.TimeDateFilename {
		og := filepath.Base(c.File)
		ext := filepath.Ext(og)
		og = strings.TrimSuffix(og, ext)
		c.File = filepath.Join(
			filepath.Dir(c.File),
			fmt.Sprintf("%s-%s%s", og, time.Now().Format("2006-01-02-15-04-05"), ext),
		)
	}
	if f, err = os.OpenFile(filepath.Join(c.Directory, c.File), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600); err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	c.Outputs = append(c.Outputs, f)
	c.ActiveLogFileName = f.Name()
	return nil
}

func (c *Configuration) setupSyslog() error {
	var (
		err   error
		proto string
		addr  string
		conn  *syslog.Writer
	)
	switch {
	case strings.Contains(c.RSyslog, "://"):
		proto = strings.Split(c.RSyslog, "://")[0]
		addr = strings.Split(c.RSyslog, "://")[1]
	case strings.Contains(c.RSyslog, ":"):
		proto = "udp"
		addr = c.RSyslog
	default:
		proto = "udp"
		addr = c.RSyslog + ":514"
	}
	if conn, err = syslog.Dial(proto, addr, syslog.LOG_INFO, "HellPot"); err != nil {
		return fmt.Errorf("failed to dial syslog server: %w", err)
	}

	c.Outputs = append(c.Outputs, zerolog.SyslogLevelWriter(conn))

	return nil
}

func (c *Configuration) SetupOutputs() error {
	if c.Directory != "" || c.File != "" {
		if err := c.setupDirAndFile(); err != nil {
			return fmt.Errorf("failed to setup log file: %w", err)
		}
	}
	if c.RSyslog != "" {
		if err := c.setupSyslog(); err != nil {
			return fmt.Errorf("failed to setup syslog: %w", err)
		}
	}

	return nil
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

var (
	ErrNoOutputs = errors.New(
		"no outputs provided, if console only logging is desired, set docker_logging to true",
	)
	ErrBothFileAndDirectory = errors.New(
		"cannot specify both file and directory unless file is a child of directory",
	)
)

var _log zerolog.Logger

func New(conf *Configuration) (zerolog.Logger, error) {
	if err := conf.Validate(); err != nil {
		return zerolog.Logger{}, fmt.Errorf("invalid logger configuration: %w", err)
	}
	if err := conf.SetupOutputs(); err != nil {
		return zerolog.Logger{}, fmt.Errorf("failed to setup logger outputs: %w", err)
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
