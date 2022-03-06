package config

import "os"

var usage = []string{
	"\n" + Title + " v" + Version + " Usage\n",
	"-c <config.toml> - Specify config file",
	"--nocolor - disable color and banner ",
	"--banner - show banner + version and exit",
	"--genconfig - write default config to 'default.toml' then exit",
}

func printUsage() {
	println(usage)
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
		case "--debug", "-v",
			forceDebug = true
		case "--trace", "-vv",
			forceTrace = true
		case "--nocolor":
			noColorForce = true
		case "--banner":
			BannerOnly = true
		case "-c", "--config":
			if len(os.Args) <= i-1 {
				panic("syntax error! expected file after -c")
			}
			loadCustomConfig(os.Args[i+1])
		default:
			continue
		}
	}
}
