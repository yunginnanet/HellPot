package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const (
	// Version roughly represents the applications current version.
	Version = "0.3"
	// Title is the name of the application used throughout the configuration process.
	Title = "HellPot"
)

var (
	// BannerOnly when toggled causes HellPot to only print the banner and version then exit.
	BannerOnly = false
	// GenConfig when toggled causes HellPot to write its default config to the cwd and then exit.
	GenConfig = false
	// NoColor when true will disable the banner and any colored console output.
	NoColor bool
)

// "http"
var (
	// BindAddr is defined via our toml configuration file. It is the address that HellPot listens on.
	BindAddr string
	// BindPort is defined via our toml configuration file. It is the port that HellPot listens on.
	BindPort string
	// Paths are defined via our toml configuration file. These are the paths that HellPot will present for "robots.txt"
	//       These are also the paths that HellPot will respond for. Other paths will throw a warning and will serve a 404.
	Paths []string

	// UseUnixSocket determines if we will listen for HTTP connections on a unix socket.
	UseUnixSocket bool

	// UnixSocketPath is defined via our toml configuration file. It is the path of the socket HellPot listens on
	// if UseUnixSocket, also defined via our toml configuration file, is set to true.
	UnixSocketPath        = ""
	UnixSocketPermissions uint32
)

// "performance"
var (
	RestrictConcurrency bool
	MaxWorkers          int
)

// "deception"
var (
	// FakeServerName is our configured value for the "Server: " response header when serving HTTP clients
	FakeServerName string
)

var (
	// Filename returns the current location of our toml config file.
	Filename string
)

var (
	f   *os.File
	err error
)

var (
	noColorForce    = false
	customconfig    = false
	home            string
	configLocations []string
)

var (
	// Debug is our global debug toggle
	Debug bool

	prefConfigLocation string
	snek               *viper.Viper
)

func init() {
	if home, err = os.UserHomeDir(); err != nil {
		panic(err)
	}
	prefConfigLocation = home + "/.config/" + Title
	snek = viper.New()
}

func writeConfig() {
	if runtime.GOOS != "windows" {
		if _, err := os.Stat(prefConfigLocation); os.IsNotExist(err) {
			if err = os.Mkdir(prefConfigLocation, 0755); err != nil {
				println("error writing new config: " + err.Error())
			}
		}
		newconfig := prefConfigLocation + "/" + "config.toml"
		if err = snek.SafeWriteConfigAs(newconfig); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		Filename = newconfig
		return
	}

	newconfig := "hellpot-config"
	snek.SetConfigName(newconfig)
	if err = snek.MergeInConfig(); err != nil {
		if err = snek.SafeWriteConfigAs(newconfig + ".toml"); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	Filename = newconfig
}

// Init will initialize our toml configuration engine and define our default configuration values which can be written to a new configuration file if desired
func Init() {
	snek.SetConfigType("toml")
	snek.SetConfigName("config")

	argParse()

	if customconfig {
		associate()
		return
	}

	acquireClue()

	setDefaults()

	for _, loc := range configLocations {
		snek.AddConfigPath(loc)
	}

	if err = snek.MergeInConfig(); err != nil {
		writeConfig()
	}

	if len(Filename) < 1 {
		Filename = snek.ConfigFileUsed()
	}

	associate()
}

func setDefaults() {
	var (
		configSections = []string{"logger", "http", "performance", "deception"}
		deflogdir      = home + "/.config/" + Title + "/logs/"
		defNoColor     = false
	)

	if runtime.GOOS == "windows" {
		deflogdir = "logs/"
		defNoColor = true
	}

	Opt := make(map[string]map[string]interface{})

	Opt["logger"] = map[string]interface{}{
		"debug":             true,
		"directory":         deflogdir,
		"nocolor":           defNoColor,
		"use_date_filename": true,
	}
	Opt["http"] = map[string]interface{}{
		"use_unix_socket":         false,
		"unix_socket_path":        "/var/run/hellpot",
		"unix_socket_permissions": "0666",
		"bind_addr":               "127.0.0.1",
		"bind_port":               "8080",
		"paths": []string{
			"wp-login.php",
			"wp-login",
		},
	}
	Opt["performance"] = map[string]interface{}{
		"restrict_concurrency": false,
		"max_workers":          256,
	}
	Opt["deception"] = map[string]interface{}{
		"server_name": "nginx",
	}

	for _, def := range configSections {
		snek.SetDefault(def, Opt[def])
	}

	if GenConfig {
		if err = snek.SafeWriteConfigAs("./config.toml"); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

}

func acquireClue() {
	configLocations = append(configLocations, "./")

	if runtime.GOOS != "windows" {
		configLocations = append(configLocations, prefConfigLocation)
		configLocations = append(configLocations, "/etc/"+Title+"/")
		configLocations = append(configLocations, "../")
		configLocations = append(configLocations, "../../")
	}
}

func loadCustomConfig(path string) {
	if f, err = os.Open(path); err != nil {
		println("Error opening specified config file: " + path)
		panic("config file open fatal error: " + err.Error())
	}
	buf, err := ioutil.ReadAll(f)
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

func printUsage() {
	println("\n" + Title + " v" + Version + " Usage\n")
	println("-c <config.toml> - Specify config file")
	println("--nocolor - disable color and banner ")
	println("--banner - show banner + version and exit")
	println("--genconfig - write default config to 'default.toml' then exit")
	os.Exit(0)
}

// TODO: should probably just make a proper CLI with flags or something
func argParse() {
	for i, arg := range os.Args {
		switch arg {
		case "-h":
			printUsage()
		case "--genconfig":
			GenConfig = true
		case "--nocolor":
			noColorForce = true
		case "--banner":
			BannerOnly = true
		case "--config":
			fallthrough
		case "-c":
			if len(os.Args) <= i-1 {
				panic("syntax error! expected file after -c")
			}
			loadCustomConfig(os.Args[i+1])
		default:
			continue
		}
	}
}

func bl(key string) bool {
	return snek.GetBool(key)
}
func st(key string) string {
	return snek.GetString(key)
}
func sl(key string) []string {
	return snek.GetStringSlice(key)
}
func it(key string) int {
	return snek.GetInt(key)
}


func associate() {
	BindAddr = st("http.bind_addr")
	BindPort = st("http.bind_port")
	Paths = sl("http.paths")
	UseUnixSocket = bl("http.use_unix_socket")
	//
	Debug = bl("logger.debug")
	logDir = st("logger.directory")
	NoColor = bl("logger.nocolor")
	//
	FakeServerName = st("deception.server_name")
	//
	RestrictConcurrency = bl("performance.restrict_concurrency")
	MaxWorkers = it("performance.max_workers")

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
