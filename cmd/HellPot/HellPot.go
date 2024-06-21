package main

import (
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

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
	if f, err = os.Create(target); err != nil {
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

func main() {
	conf := config.CLIFlags.Lookup("config").Value
	if conf.String() == "" {
		conf = config.CLIFlags.Lookup("c").Value
	}

	usingDefaults := true
	resolvedConf := conf.String()

	uconf, _ := os.UserConfigDir()
	if uconf == "" && os.Getenv("HOME") != "" {
		uconf = filepath.Join(os.Getenv("HOME"), ".config")
	}

	if resolvedConf == "" {
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
	}

	var setupErr error
	var err error
	var f *os.File

	f, err = os.Open(resolvedConf)
	if err == nil {
		runningConfig, setupErr = config.Setup(f)
	}
	switch {
	case setupErr != nil:
		println("failed to setup config: " + setupErr.Error())
	case err != nil:
		println("failed to open config file for reading: " + err.Error())
		println("trying to create it....")
		wroteOK := writeConfig(resolvedConf)
		if wroteOK {
			break
		}
	case runningConfig != nil:
		usingDefaults = false
		_ = f.Close()
	}

	if runningConfig == nil {
		if runningConfig, err = config.Setup(nil); err != nil || runningConfig == nil {
			panic("failed to setup default config...\n" + err.Error())
		}
	}

	log, _ := logger.New(runningConfig.Logger)

	if usingDefaults {
		log.Warn().Msg("continuing with default configuration in ")
		for i := 5; i > 0; i-- {
			print(strconv.Itoa(i))
			for i := 0; i < 5; i++ {
				time.Sleep(200 * time.Millisecond)
				print(".")
			}
		}
	}

	if !runningConfig.Logger.NoColor {
		extra.Banner()
	}

	log.Info().Msg("🔥 Starting HellPot 🔥")
	log.Info().Msg("Version: " + version.Version)
	log.Info().Msg("PID: " + strconv.Itoa(os.Getpid()))
	log.Info().Msg("Using config file: " + resolvedConf)
	if usingDefaults {
		log.Warn().Msg("Using default configuration")
	}
	log.Debug().Msg("Debug logging enabled")
	log.Trace().Msg("Trace logging enabled")

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Fatal().Err(http.Serve(runningConfig)).Msg("HTTP error")
	}()

	<-stopChan // wait for SIGINT
	log.Warn().Msg("Shutting down server...")
}
