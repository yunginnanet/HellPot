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
		if err = os.MkdirAll(prefConfigLocation, 0755); err != nil {
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

	setConfigFileLocations()
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

	associateExportedVariables()
}

func setDefaults() {
	var (
		configSections = []string{"logger", "http", "performance", "deception", "ssh"}
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

func setConfigFileLocations() {
	configLocations = append(configLocations, "./")

	if runtime.GOOS != "windows" {
		configLocations = append(configLocations, prefConfigLocation)
		configLocations = append(configLocations, "/etc/"+Title+"/")
		configLocations = append(configLocations, "../")
		configLocations = append(configLocations, "../../")
		configLocations = append(configLocations, "./")

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
