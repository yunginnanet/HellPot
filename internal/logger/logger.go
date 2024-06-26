package logger

import (
	"errors"
	"fmt"
	"io"
	"log/syslog"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
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
	NoConsole     bool `koanf:"noconsole"`
	DockerLogging bool `koanf:"docker_logging"`
	// ConsoleTimeFormat sets the time format for the console.
	// The string is passed to time.Format() down the line.
	ConsoleTimeFormat string `koanf:"console_time_format"`

	// TimeDateFilename sets the log file name to include the date and time.
	TimeDateFilename bool `koanf:"use_date_filename"`

	Directory     string `koanf:"directory"`
	File          string `koanf:"log_file"`
	RSyslog       string `koanf:"rsyslog_address"`
	rsyslogTarget string

	ActiveLogFileName string `koanf:"active_log_file_name"`

	Outputs []io.Writer `koanf:"-"`
}

func (c *Configuration) Set(k string, v any) error {
	k = strings.ToLower(k)
	k = strings.TrimPrefix(k, "logger.")
	k = strings.ReplaceAll(k, "__", "_")
	k = strings.ReplaceAll(k, ".", "_")

	ref := reflect.ValueOf(c)
	if ref.Kind() != reflect.Ptr {
		panic("not a pointer")
	}
	ref = ref.Elem()
	if ref.Kind() != reflect.Struct {
		panic("not a struct")
	}

	var field reflect.Value

	for i := 0; i < ref.NumField(); i++ {
		strutTag := ref.Type().Field(i).Tag.Get("koanf")
		if strings.ToLower(strutTag) == strings.ToLower(k) {
			field = ref.Field(i)
			break
		}
	}

	if field == (reflect.Value{}) {
		return fmt.Errorf("field %s does not exist", k)
	}

	if !field.CanSet() {
		return fmt.Errorf("field %s cannot be set", k)
	}

	switch field.Kind() {
	case reflect.Bool:
		if vstr, vstrok := v.(string); vstrok {
			if vb, err := strconv.ParseBool(vstr); err == nil {
				field.SetBool(vb)
				return nil
			}
		}
		if b, ok := v.(bool); ok {
			field.SetBool(b)
			return nil
		}
		return fmt.Errorf("field %s is not a bool", k)
	case reflect.String:
		if s, ok := v.(string); ok {
			field.SetString(s)
			return nil
		}
		return fmt.Errorf("field %s is not a string", k)
	case reflect.Slice:
		if s, ok := v.([]string); ok {
			field.Set(reflect.ValueOf(s))
			return nil
		}
		return fmt.Errorf("field %s is not a slice", k)
	case reflect.Int:
		if i, ok := v.(int); ok {
			field.SetInt(int64(i))
			return nil
		}
		return fmt.Errorf("field %s is not an int", k)
	default:
		return fmt.Errorf("field %s is not a supported type (%T)", k, v)
	}
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

func (c *Configuration) setupSyslog() (error, bool) {
	if c.RSyslog == "" {
		return nil, false
	}
	var (
		err   error
		proto string
		addr  string
		conn  *syslog.Writer
	)
	switch {
	case strings.ToLower(c.RSyslog) == "local":
		proto = ""
		addr = ""
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

	if proto != "" && addr != "" {
		println("dialing syslog server: " + c.rsyslogTarget + "...")
	}

	syslogLogLevel := syslog.LOG_INFO
	if c.Debug || c.Trace {
		syslogLogLevel = syslog.LOG_DEBUG
	}

	if conn, err = syslog.Dial(proto, addr, syslogLogLevel, "HellPot"); err != nil {
		return fmt.Errorf("failed to dial syslog server: %w", err), false
	}

	c.rsyslogTarget = proto + "://" + addr

	c.Outputs = append(c.Outputs, zerolog.SyslogLevelWriter(conn))

	return nil, true
}

func (c *Configuration) SetupOutputs() (err error, rsyslogEnabled bool) {
	if c.Directory != "" || c.File != "" {
		if err = c.setupDirAndFile(); err != nil {
			return fmt.Errorf("failed to setup log file: %w", err), false
		}
	}

	if c.RSyslog != "" {
		if err, rsyslogEnabled = c.setupSyslog(); err != nil {
			return fmt.Errorf("failed to setup syslog: %w", err), false
		}
	}

	consoleSeen := false

	for _, out := range c.Outputs {
		if out == nil {
			return fmt.Errorf("nil output provided"), false
		}
		if out == os.Stdout || out == os.Stderr {
			consoleSeen = true
		}
	}

	if !consoleSeen && !c.NoConsole {
		c.Outputs = append(c.Outputs, os.Stdout)
	}

	return nil, rsyslogEnabled
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
	var err error
	var rsyslogEnabled bool
	if err, rsyslogEnabled = conf.SetupOutputs(); err != nil {
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

	if rsyslogEnabled {
		_log.Info().Str("target", conf.rsyslogTarget).Msg("remote syslog connection established")
	}

	return _log, nil
}
