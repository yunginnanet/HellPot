package config

import (
	"os"
)

var (
	forceDebug = false
	forceTrace = false
)

var argBoolMap = map[string]*bool{
	"--debug": &forceDebug, "-v": &forceDebug, "--trace": &forceTrace, "-vv": &forceTrace,
	"--nocolor": &noColorForce, "--banner": &BannerOnly, "--genconfig": &GenConfig,
}

// TODO: should probably just make a proper CLI with flags or something
func argParse() {
	for i, arg := range os.Args {
		if t, ok := argBoolMap[arg]; ok {
			*t = true
			continue
		}
		switch arg {
		case "-h", "--help":
			CLI.printUsage()
		case "-c", "--config":
			if len(os.Args) < i+2 {
				println("missing config file after -c/--config")
				os.Exit(1)
			}
			loadCustomConfig(os.Args[i+1])
		case "-b", "--book":
			if len(os.Args) < i+2 {
				println("missing book file after -b/--book")
				os.Exit(1)
			}
			BookFilename = os.Args[i+1]
			UseCustomHeffalump = true
		default:
			continue
		}
	}
}
