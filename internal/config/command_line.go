package config

import (
	"flag"
	"io"
	"os"
	"strings"

	"github.com/yunginnanet/HellPot/internal/extra"
	"github.com/yunginnanet/HellPot/internal/version"
)

var CLIFlags = flag.NewFlagSet("cli", flag.ExitOnError)

func init() {
	for _, arg := range os.Args {
		// check for unit test flags
		if strings.Contains(arg, "test.testlogfile") {
			// we're in a unit test, bail
			return
		}
		if strings.Contains(arg, "test.v") {
			return
		}
	}

	CLIFlags.Bool("logger-debug", false, "force debug logging")
	CLIFlags.Bool("logger-trace", false, "force trace logging")
	CLIFlags.Bool("logger-nocolor", false, "force no color logging")
	CLIFlags.String("bespoke-grimoire", "", "specify a custom file used for text generation")
	CLIFlags.Bool("banner", false, "show banner and version then exit")
	CLIFlags.Bool("genconfig", false, "write default config to stdout then exit")
	CLIFlags.Bool("h", false, "show this help and exit")
	CLIFlags.Bool("help", false, "show this help and exit")
	CLIFlags.String("c", "", "specify config file")
	CLIFlags.String("config", "", "specify config file")
	CLIFlags.String("version", "", "show version and exit")
	CLIFlags.String("v", "", "show version and exit")
	if err := CLIFlags.Parse(os.Args[1:]); err != nil {
		println(err.Error())
		// flag.ExitOnError will call os.Exit(2)
	}
	if CLIFlags.Lookup("h").Value.String() == "true" || CLIFlags.Lookup("help").Value.String() == "true" {
		CLIFlags.Usage()
		os.Exit(0)
	}
	if CLIFlags.Lookup("version").Value.String() == "true" || CLIFlags.Lookup("v").Value.String() == "true" {
		_, _ = os.Stdout.WriteString("HellPot version: " + version.Version + "\n")
		os.Exit(0)
	}
	if CLIFlags.Lookup("genconfig").Value.String() == "true" {
		if n, err := io.Copy(os.Stdout, Defaults.IO); err != nil || n == 0 {
			if err == nil {
				err = io.EOF
			}
			panic(err)
		}
		os.Exit(0)
	}
	if CLIFlags.Lookup("banner").Value.String() == "true" {
		extra.Banner()
		os.Exit(0)
	}
}
