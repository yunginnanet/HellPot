package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/yunginnanet/HellPot/internal/http"
)

func testMain(t *testing.T) (string, string, chan os.Signal, error) {
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

func TestHellPot(t *testing.T) {
	logDir := filepath.Join(t.TempDir(), "logs")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		t.Fatal(err)
	}
	confFile := filepath.Join(t.TempDir(), "HellPot_test.toml")
	t.Setenv("HELLPOT_LOGGER_DIRECTORY", logDir)
	t.Setenv("HELLPOT_CONFIG", confFile)

	resolvedConf, logFile, stopChan, err := testMain(t)
	if err != nil {
		t.Fatal(err)
	}
	if stopChan == nil {
		t.Fatal("stopChan is nil")
	}
	if resolvedConf == "" {
		t.Fatal("resolvedConf is empty")
	}
	if logFile == "" {
		t.Fatal("logFile is empty")
	}
	if _, err = os.Stat(logFile); err != nil {
		t.Fatal(err)
	}
	if resolvedConf != confFile {
		t.Errorf("expected %s, got %s", confFile, resolvedConf)
	}
	if logFile != filepath.Join(logDir, "HellPot.log") {
		t.Errorf("expected %s, got %s", filepath.Join(logDir, "HellPot.log"), logFile)
	}
	t.Log("resolvedConf: ", resolvedConf)
	t.Log("logFile: ", logFile)
	time.Sleep(100 * time.Millisecond)
	stopChan <- os.Interrupt
}
