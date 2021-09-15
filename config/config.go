package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const (
	version = "0.2"
	title   = "HellPot"
)

var (
	// BindAddr is defined via our toml configuration file. It is the address that HellPot listens on.
	BindAddr string
	// BindPort is defined via our toml configuration file. It is the port that HellPot listens on.
	BindPort string
	// Paths are defined via our toml configuration file. These are the paths that HellPot will present for "robots.txt"
	//       These are also the paths that HellPot will respond for. Other paths will throw a warning and will serve a 404.
	Paths []string
)

var (
	// Filename returns the current location of our toml config file.
	Filename string
)

var (
	f   *os.File
	err error

	customconfig    = false
	home            string
	configLocations []string
)

/*
Opt represents our program options.
    nitially the values that are defined in Opt will be used to define details.
	Beyond that, default values will be replaced by options from our config file.
*/
var Opt map[string]map[string]interface{}

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
	prefConfigLocation = home + "/.config/" + title
	Opt = make(map[string]map[string]interface{})
	snek = viper.New()
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
		if _, err := os.Stat(prefConfigLocation); os.IsNotExist(err) {
			if err = os.Mkdir(prefConfigLocation, 0755); err != nil {
				panic(err)
			}
		}

		newconfig := prefConfigLocation + "/" + "config.toml"

		if err = snek.SafeWriteConfigAs(newconfig); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		Filename = newconfig
	}

	if len(Filename) < 1 {
		Filename = snek.ConfigFileUsed()
	}

	associate()
}

func setDefaults() {
	var configSections = []string{"logger", "http"}

	Opt["logger"] = map[string]interface{}{
		"debug":             true,
		"directory":         home + "/.config/" + title + "/logs/",
		"use_date_filename": true,
	}
	Opt["http"] = map[string]interface{}{
		"bind_addr": "127.0.0.1",
		"bind_port": "8080",
		"paths": []string{
			"wp-login.php",
			"wp-login",
		},
	}

	for _, def := range configSections {
		snek.SetDefault(def, Opt[def])
	}
}

func acquireClue() {
	configLocations = append(configLocations, prefConfigLocation)
	configLocations = append(configLocations, "/etc/"+title+"/")
	configLocations = append(configLocations, "./")
	configLocations = append(configLocations, "../")
	configLocations = append(configLocations, "../../")
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

// TODO: should probably just make a proper CLI with flags or something
func argParse() {
	for i, arg := range os.Args {
		switch arg {
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

func associate() {
	var newOpt map[string]map[string]interface{}
	newOpt = Opt
	for category, opt := range Opt {
		for optname, value := range opt {
			if snek.IsSet(category + "." + optname) {
				newOpt[category][optname] = value
			}
		}
	}

	Opt = newOpt
	Debug = snek.GetBool("logger.debug")
	logDir = snek.GetString("logger.directory")
	BindAddr = snek.GetString("http.bind_addr")
	BindPort = snek.GetString("http.bind_port")
	Paths = snek.GetStringSlice("http.paths")

	if Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
