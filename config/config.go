package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// generic vars
var (
	f                  *os.File
	err                error
	noColorForce       = false
	customconfig       = false
	home               string
	prefConfigLocation string
	snek               *viper.Viper
)

// exported generic vars
var (
	// Debug is the value of our debug on/off toggle as per the current configuration.
	Debug              bool
	// Filename returns the current location of our toml config file.
	Filename string
)

func init() {
	if home, err = os.UserHomeDir(); err != nil {
		panic(err)
	}
	prefConfigLocation = home + "/.config/" + Title
	snek = viper.New()
}

func writeConfig() {
	if runtime.GOOS == "windows" {
		newconfig := "hellpot-config"
		snek.SetConfigName(newconfig)
		if err = snek.MergeInConfig(); err != nil {
			if err = snek.SafeWriteConfigAs(newconfig + ".toml"); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
		return
	}

	if _, err := os.Stat(prefConfigLocation); os.IsNotExist(err) {
		if err = os.MkdirAll(prefConfigLocation, 0o755); err != nil {
			println("error writing new config: " + err.Error())
			os.Exit(1)
		}
	}

	newconfig := prefConfigLocation + "/" + "config.toml"
	if err = snek.SafeWriteConfigAs(newconfig); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	Filename = newconfig
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

	if err = snek.MergeInConfig(); err != nil {
		writeConfig()
	}

	if len(Filename) < 1 {
		Filename = snek.ConfigFileUsed()
	}

	associateExportedVariables()
}

func getConfigPaths() (paths []string) {
	paths = append(paths, "./")

	if runtime.GOOS != "windows" {
		paths = append(paths,
			prefConfigLocation, "/etc/"+Title+"/", "../", "../../")
	}

	return
}

func loadCustomConfig(path string) {
	if f, err = os.Open(path); err != nil {
		println("Error opening specified config file: " + path)
		panic("config file open fatal error: " + err.Error())
	}
	buf, err := io.ReadAll(f)
	err2 := snek.ReadConfig(bytes.NewBuffer(buf))
	switch {
	case err != nil:
		fmt.Println("config file read fatal error: ", err.Error())
	case err2 != nil:
		fmt.Println("config file read fatal error: ", err2.Error())
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
		"logger.directory":      &logDir,
		"deception.server_name": &FakeServerName,
	}
	// string slice options and their exported variables
	strSliceOpt := map[string]*[]string{
		"http.paths": &Paths,
	}
	// bool options and their exported variables
	boolOpt := map[string]*bool{
		"http.use_unix_socket":             &UseUnixSocket,
		"logger.debug":                     &Debug,
		"performance.restrict_concurrency": &RestrictConcurrency,
		"logger.nocolor":                   &NoColor,
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

	if Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
