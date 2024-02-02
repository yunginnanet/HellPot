package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// generic vars
var (
	noColorForce       = false
	customconfig       = false
	home               string
	prefConfigLocation string
	snek               = viper.New()
)

// exported generic vars
var (
	// Trace is the value of our trace (extra verbose)  on/off toggle as per the current configuration.
	Trace bool
	// Debug is the value of our debug (verbose) on/off toggle as per the current configuration.
	Debug bool
	// Filename returns the current location of our toml config file.
	Filename string
	// UseCustomHeffalump decides if a custom Heffalump is to be used
	UseCustomHeffalump = false
	// Grimoire returns the current location of a possible source of suffering file
	Grimoire string
)

func writeConfig() {
	if _, err := os.Stat(prefConfigLocation); os.IsNotExist(err) {
		if err = os.MkdirAll(prefConfigLocation, 0o750); err != nil {
			println("error writing new config: " + err.Error())
			os.Exit(1)
		}
	}
	Filename = prefConfigLocation + "/" + "config.toml"
	if err := snek.SafeWriteConfigAs(Filename); err != nil {
		fmt.Println("Failed to write new configuration file to '" + Filename + "': " + err.Error())
		os.Exit(1)
	}
}

// Init will initialize our toml configuration engine and define our default configuration values which can be written to a new configuration file if desired
func Init() {
	snek.SetConfigType("toml")
	snek.SetConfigName("config")

	argParse()

	if customconfig {
		associateExportedVariables()
		return
	}

	setDefaults()

	for _, loc := range getConfigPaths() {
		snek.AddConfigPath(loc)
	}

	if err := snek.MergeInConfig(); err != nil {
		println("Error reading configuration file: " + err.Error())
		println("Writing new configuration file...")
		writeConfig()
	}

	if len(Filename) < 1 {
		Filename = snek.ConfigFileUsed()
	}

	snek.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	snek.SetEnvPrefix(Title)

	snek.AutomaticEnv()

	associateExportedVariables()
}

func getConfigPaths() (paths []string) {
	paths = append(paths, "./")
	//goland:noinspection GoBoolExpressions
	if runtime.GOOS != "windows" {
		paths = append(paths,
			prefConfigLocation, "/etc/"+Title+"/", "../", "../../")
	}
	return
}

func loadCustomConfig(path string) {
	/* #nosec */
	cf, err := os.Open(path)
	if err != nil {
		println("Error opening specified config file: " + path)
		println(err.Error())
		os.Exit(1)
	}

	Filename, err = filepath.Abs(path)
	if len(Filename) < 1 || err != nil {
		Filename = path
	}

	defer func(f *os.File) {
		if fcerr := f.Close(); fcerr != nil {
			fmt.Println("failed to close file handler for config file: ", fcerr.Error())
		}
	}(cf)

	buf, err1 := io.ReadAll(cf)
	err2 := snek.ReadConfig(bytes.NewBuffer(buf))

	switch {
	case err1 != nil:
		fmt.Println("config file read fatal error during i/o: ", err1.Error())
		os.Exit(1)
	case err2 != nil:
		fmt.Println("config file read fatal error during parse: ", err2.Error())
		os.Exit(1)
	default:
		break
	}

	customconfig = true
}

func processOpts() {
	// string options and their exported variables
	stringOpt := map[string]*string{
		"http.bind_addr":        &HTTPBind,
		"http.bind_port":        &HTTPPort,
		"http.real_ip_header":   &HeaderName,
		"logger.directory":      &logDir,
		"deception.server_name": &FakeServerName,
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
		*opt = snek.GetString(key)
	}
	for key, opt := range strSliceOpt {
		*opt = snek.GetStringSlice(key)
	}
	for key, opt := range boolOpt {
		*opt = snek.GetBool(key)
	}
	for key, opt := range intOpt {
		*opt = snek.GetInt(key)
	}
}

func associateExportedVariables() {
	processOpts()

	if noColorForce {
		NoColor = true
	}

	if UseUnixSocket {
		UnixSocketPath = snek.GetString("http.unix_socket_path")
		parsedPermissions, err := strconv.ParseUint(snek.GetString("http.unix_socket_permissions"), 8, 32)
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
