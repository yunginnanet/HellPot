package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/yunginnanet/HellPot/internal/config"
	"github.com/yunginnanet/HellPot/internal/extra"
	"github.com/yunginnanet/HellPot/internal/http"
	"github.com/yunginnanet/HellPot/internal/logger"
	"github.com/yunginnanet/HellPot/internal/version"
)

var (
	runningConfig *config.Parameters
)

func writeConfig(target string) bool {
	var f *os.File
	var err error
	f, err = os.Create(target) // #nosec G304 -- go home gosec, you're drunk
	if err != nil {
		println("failed to create config file: " + err.Error())
		return false
	}
	if _, err = io.Copy(f, config.Defaults.IO); err != nil {
		println("failed to write default config to file: " + err.Error())
		_ = f.Close()
		return false
	}
	println("wrote default config to " + target)
	runningConfig, _ = config.Setup(f)
	_ = f.Close()
	return true
}

func searchConfig() string {
	var resolvedConf string
	uconf, _ := os.UserConfigDir()
	if uconf == "" && os.Getenv("HOME") != "" {
		uconf = filepath.Join(os.Getenv("HOME"), ".config")
	}

	for _, path := range []string{
		"/etc/HellPot/config.toml",
		"/usr/local/etc/HellPot/config.toml",
		"./config.toml",
		filepath.Join(uconf, "HellPot", "config.toml"),
	} {
		if _, err := os.Stat(path); err == nil {
			resolvedConf = path
			break
		}
	}
	return resolvedConf
}

func readConfig(resolvedConf string) error {
	var err error
	var setupErr error
	var f *os.File

	if resolvedConf == "" {
		return fmt.Errorf("%w: provided config file is an empty string", io.EOF)
	}

	f, err = os.Open(resolvedConf) // #nosec G304 go home gosec, you're drunk
	if err == nil {
		runningConfig, setupErr = config.Setup(f)
	}
	switch {
	case setupErr != nil:
		println("failed to setup config: " + setupErr.Error())
		err = setupErr
	case err != nil:
		println("failed to open config file for reading: " + err.Error())
		println("trying to create it....")
		wroteOK := writeConfig(resolvedConf)
		if wroteOK {
			break
		}
		println("failed to create config file, cannot continue")
		return fmt.Errorf("failed to create config file: %w", err)
	case runningConfig != nil:
		_ = f.Close()
	}

	return err
}

func resolveConfig() (usingDefaults bool, resolvedConf string, err error) {
	if config.CLIFlags != nil {
		confRoot := config.CLIFlags.Lookup("config")
		if confRoot != nil && confRoot.Value.String() != "" {
			resolvedConf = confRoot.Value.String()
		}
	}

	if resolvedConf == "" && os.Getenv("HELLPOT_CONFIG_FILE") != "" {
		resolvedConf = os.Getenv("HELLPOT_CONFIG_FILE")
	}

	if resolvedConf == "" {
		resolvedConf = searchConfig()
	}

	if err = readConfig(resolvedConf); err != nil && !errors.Is(err, io.EOF) {
		return false, "", err
	}

	if runningConfig == nil {
		println("still no config, trying defaults...")
		if runningConfig, err = config.Setup(nil); err != nil || runningConfig == nil {
			if err == nil {
				err = errors.New("unknown failure resulting in missing configuration, cannot continue")
			}
			return false, "", err
		}
		return true, "", nil
	}

	return false, resolvedConf, nil
}

func setup(stopChan chan os.Signal) (log zerolog.Logger, logFile string, resolvedConf string, err error) {
	var usingDefaults bool

	if usingDefaults, resolvedConf, err = resolveConfig(); err != nil {
		return
	}

	//goland:noinspection GoNilness // we check for nil above
	if log, err = logger.New(runningConfig.Logger); err != nil {
		return
	}

	logFile = runningConfig.Logger.ActiveLogFileName

	if usingDefaults {
		runningConfig.UsingDefaults = true
		log.Warn().Msg("continuing with default configuration in ")
		for i := 5; i > 0; i-- {
			print(strconv.Itoa(i))
			for i := 0; i < 5; i++ {
				time.Sleep(200 * time.Millisecond)
				print(".")
			}
		}
	}

	if //goland:noinspection GoNilness
	!runningConfig.Logger.NoColor {
		extra.Banner()
	}

	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	return
}

func testMain(t interface{ Error(...any) }) (string, string, chan os.Signal, error) {
	stopChan := make(chan os.Signal, 1)
	log, logFile, resolvedConf, err := setup(stopChan)
	if err == nil {
		log.Info().Msg("🔥 Starting HellPot 🔥")
		go func() {
			terr := http.Serve(runningConfig)
			if terr != nil {
				t.Error("failed to serve HTTP: " + terr.Error())
				close(stopChan)
			}
		}()
	}
	//goland:noinspection GoNilness
	return resolvedConf, logFile, stopChan, err
}

func main() {
	stopChan := make(chan os.Signal, 1)
	log, _, resolvedConf, err := setup(stopChan)

	if err != nil {
		println("failed to start: " + err.Error())
		os.Exit(1)
	}

	log.Info().Msg("🔥 Starting HellPot 🔥")
	log.Info().Msg("Version: " + version.Version)
	log.Info().Msg("PID: " + strconv.Itoa(os.Getpid()))
	log.Info().Msg("Using config file: " + resolvedConf)
	if runningConfig.UsingDefaults {
		log.Warn().Msg("Using default configuration")
	}
	if runningConfig.Logger.RSyslog != "" {
		log.Info().Msg("Logging to syslog: " + runningConfig.Logger.RSyslog)
	}
	if runningConfig.Logger.ActiveLogFileName != "" {
		log.Info().Msg("Logging to file: " + runningConfig.Logger.ActiveLogFileName)
	}
	if runningConfig.Logger.DockerLogging &&
		runningConfig.Logger.File == "" &&
		runningConfig.Logger.Directory == "" &&
		runningConfig.Logger.RSyslog == "" {
		log.Info().Msg("Only logging to standard output")
	}

	log.Debug().Msg("Debug logging enabled")
	log.Trace().Msg("Trace logging enabled")

	go func() {
		log.Fatal().Err(http.Serve(runningConfig)).Msg("HTTP error")
	}()

	sig := <-stopChan // wait for SIGINT
	log.Warn().Interface("signal_received", sig).
		Msg("Shutting down server...")
}
