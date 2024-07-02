package config

import (
	"flag"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/yunginnanet/HellPot/internal/extra"
	"github.com/yunginnanet/HellPot/internal/version"
)

var CLIFlags = flag.NewFlagSet("config", flag.ContinueOnError)

var (
	sliceDefs = make(map[string][]string)
	slicePtrs = make(map[string]*string)
)

func addCLIFlags() {
	parse := func(k string, v interface{}, nestedName string) {
		switch casted := v.(type) {
		case bool:
			CLIFlags.Bool(nestedName, casted, "set "+k)
		case string:
			CLIFlags.String(nestedName, casted, "set "+k)
		case int:
			CLIFlags.Int(nestedName, casted, "set "+k)
		case float64:
			CLIFlags.Float64(nestedName, casted, "set "+k)
		case []string:
			sliceDefs[nestedName] = casted
			joined := strings.Join(sliceDefs[nestedName], ",")
			slicePtrs[nestedName] = CLIFlags.String(nestedName, joined, "set "+k)
		}
	}

	for key, val := range Defaults.val {
		if _, ok := val.(map[string]interface{}); !ok {
			parse(key, val, key)
			continue
		}
		nested, ok := val.(map[string]interface{})
		if !ok {
			// linter was confused by the above check
			panic("unreachable, if you see this you have entered a real life HellPot")
		}
		for k, v := range nested {
			nestedName := key + "." + k
			parse(k, v, nestedName)
		}
	}
}

var replacer = map[string][]string{
	"-h": {"-help"},
	"-v": {"-version"},
	"-c": {"-config"},
	"-g": {"-bespoke.enable_grimoire", "true", "-bespoke.grimoire_file"},
}

func InitCLI() {
	newArgs := make([]string, 0)
	for _, arg := range os.Args {
		if repl, ok := replacer[arg]; ok {
			newArgs = append(newArgs, repl...)
			continue
		}
		// check for unit test flags
		if !strings.HasPrefix(arg, "-test.") {
			newArgs = append(newArgs, arg)
		}
	}

	newArgs = slices.Compact(newArgs)

	CLIFlags.Bool("banner", false, "show banner and version then exit")
	CLIFlags.Bool("genconfig", false, "write default config to stdout then exit")
	CLIFlags.Bool("help", false, "show this help and exit")
	CLIFlags.String("config", "", "specify config file")
	CLIFlags.String("version", "", "show version and exit")

	addCLIFlags()

	if err := CLIFlags.Parse(newArgs[1:]); err != nil {
		println(err.Error())
		os.Exit(2)
	}

	for defaultKey, defaultVal := range Defaults.val {
		switch defaultVal.(type) {
		case bool:
			parsedBool, pErr := strconv.ParseBool(CLIFlags.Lookup(defaultKey).Value.String())
			if pErr != nil {
				continue
			}
			if parsedBool == Defaults.val[defaultKey].(bool) {
				fl := CLIFlags.Lookup(defaultKey)
				*fl = flag.Flag{}
				fl = nil
			}
		case string:
			if CLIFlags.Lookup(defaultKey).Value.String() == Defaults.val[defaultKey].(string) ||
				CLIFlags.Lookup(defaultKey).Value.String() == "" {
				fl := CLIFlags.Lookup(defaultKey)
				*fl = flag.Flag{}
				fl = nil
			}
		}

	}

	if os.Getenv("HELLPOT_CONFIG") != "" {
		if err := CLIFlags.Set("config", os.Getenv("HELLPOT_CONFIG")); err != nil {
			panic(err)
		}
	}
	if CLIFlags.Lookup("help").Value.String() == "true" {
		CLIFlags.Usage()
		os.Exit(0)
	}
	if CLIFlags.Lookup("version").Value.String() == "true" {
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
