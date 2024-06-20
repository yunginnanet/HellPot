package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/rs/zerolog"

	"github.com/knadh/koanf/parsers/toml"
	viper "github.com/knadh/koanf/v2"
)

// generic vars
var (
	noColorForce = false
	customconfig = false
	home         string
	snek         = viper.New(".")
)

func init() {
	home, _ = os.UserHomeDir()
	if home == "" {
		home = os.Getenv("HOME")
	}
	if home == "" {
		println("WARNING: could not determine home directory")
	}
}

// exported generic vars
var (
	// Trace is the value of our trace (extra verbose)  on/off toggle as per the current configuration.
	Trace bool
	// Debug is the value of our debug (verbose) on/off toggle as per the current configuration.
	Debug bool
	// Filename returns the current location of our toml config file.
	Filename string
)

func writeConfig() string {
	prefConfigLocation, _ := os.UserConfigDir()

	if prefConfigLocation == "" {
		home, _ = os.UserHomeDir()
		prefConfigLocation = filepath.Join(home, ".config", Title)
	}

	if _, err := os.Stat(prefConfigLocation); os.IsNotExist(err) {
		if err = os.MkdirAll(prefConfigLocation, 0o750); err != nil && !errors.Is(err, os.ErrExist) {
			println("error writing new config: " + err.Error())
			os.Exit(1)
		}
	}
	Filename = filepath.Join(prefConfigLocation, "config.toml")

	tomld, terr := toml.Parser().Marshal(snek.All())
	if terr != nil {
		fmt.Println("Failed to marshal new configuration file: " + terr.Error())
		os.Exit(1)
	}

	if err := os.WriteFile(Filename, tomld, 0o600); err != nil {
		println("error writing new config: " + err.Error())
		os.Exit(1)
	}

	return Filename
}

// Init will initialize our toml configuration engine and define our default configuration values which can be written to a new configuration file if desired
func Init() {
	argParse()

	if customconfig {
		associateExportedVariables()
		return
	}

	setDefaults()

	chosen := ""

	uconf, _ := os.UserConfigDir()

	switch runtime.GOOS {
	case "windows":
		//
	default:
		if _, err := os.Stat(filepath.Join("/etc/", Title, "config.toml")); err == nil {
			chosen = filepath.Join("/etc/", Title, "config.toml")
		}
	}

	if chosen == "" && uconf == "" && home != "" {
		uconf = filepath.Join(home, ".config")
	}

	if chosen == "" && uconf != "" {
		_ = os.MkdirAll(filepath.Join(uconf, Title), 0750)
		chosen = filepath.Join(uconf, Title, "config.toml")
	}

	if chosen == "" {
		pwd, _ := os.Getwd()
		if _, err := os.Stat("./config.toml"); err == nil {
			chosen = "./config.toml"
		} else {
			if _, err := os.Stat(filepath.Join(pwd, "config.toml")); err == nil {
				chosen = filepath.Join(pwd, "config.toml")
			}
		}
	}

	loadErr := snek.Load(file.Provider(chosen), toml.Parser())

	if chosen == "" || loadErr != nil {
		println("No configuration file found, writing new configuration file...")
		chosen = writeConfig()
	}
	Filename = chosen

	if loadErr = snek.Load(file.Provider(chosen), toml.Parser()); loadErr != nil {
		fmt.Println("failed to load default config file: ", loadErr.Error())
		os.Exit(1)
	}

	/*	snek.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
		snek.SetEnvPrefix(Title)
		snek.AutomaticEnv()
	*/
	associateExportedVariables()
}

func loadCustomConfig(path string) {
	Filename, _ = filepath.Abs(path)

	if err := snek.Load(file.Provider(Filename), toml.Parser()); err != nil {
		fmt.Println("failed to load specified config file: ", err.Error())
		os.Exit(1)
	}

	customconfig = true
}

func processOpts() {
	// string options and their exported variables
	stringOpt := map[string]*string{
		"http.bind_addr":             &HTTPBind,
		"http.bind_port":             &HTTPPort,
		"http.real_ip_header":        &HeaderName,
		"logger.directory":           &logDir,
		"logger.console_time_format": &ConsoleTimeFormat,
		"deception.server_name":      &FakeServerName,
	}
	// string slice options and their exported variables
	strSliceOpt := map[string]*[]string{
		"http.router.paths":            &Paths,
		"http.uagent_string_blacklist": &UseragentBlacklistMatchers,
	}
	// bool options and their exported variables
	boolOpt := map[string]*bool{
		"performance.restrict_concurrency": &RestrictConcurrency,
		"http.use_unix_socket":             &UseUnixSocket,
		"logger.debug":                     &Debug,
		"logger.trace":                     &Trace,
		"logger.nocolor":                   &NoColor,
		"logger.docker_logging":            &DockerLogging,
		"http.router.makerobots":           &MakeRobots,
		"http.router.catchall":             &CatchAll,
	}
	// integer options and their exported variables
	intOpt := map[string]*int{
		"performance.max_workers": &MaxWorkers,
	}

	for key, opt := range stringOpt {
		*opt = snek.String(key)
	}
	for key, opt := range strSliceOpt {
		*opt = snek.Strings(key)
	}
	for key, opt := range boolOpt {
		*opt = snek.Bool(key)
	}
	for key, opt := range intOpt {
		*opt = snek.Int(key)
	}
}

func associateExportedVariables() {
	_ = snek.Load(env.Provider("HELLPOT_", ".", func(s string) string {
		s = strings.TrimPrefix(s, "HELLPOT_")
		s = strings.ToLower(s)
		s = strings.ReplaceAll(s, "__", " ")
		s = strings.ReplaceAll(s, "_", ".")
		s = strings.ReplaceAll(s, " ", "_")
		return s
	}), nil)

	processOpts()

	if noColorForce {
		NoColor = true
	}

	if UseUnixSocket {
		UnixSocketPath = snek.String("http.unix_socket_path")
		parsedPermissions, err := strconv.ParseUint(snek.String("http.unix_socket_permissions"), 8, 32)
		if err == nil {
			UnixSocketPermissions = uint32(parsedPermissions)
		}
	}

	// We set exported variables here so that it tracks when accessed from other packages.

	if Debug || forceDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		Debug = true
	}
	if Trace || forceTrace {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		Trace = true
	}
}
