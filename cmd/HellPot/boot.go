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
	"github.com/yunginnanet/HellPot/internal/logger"
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
