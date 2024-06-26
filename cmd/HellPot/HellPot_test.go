package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/yunginnanet/HellPot/internal/config"
	"github.com/yunginnanet/HellPot/internal/http"
)

func testMain(t *testing.T) (string, string, chan os.Signal, *config.Parameters, error) {
	t.Setenv("HELLPOT_TEST_MODE", "true")
	t.Helper()
	stopChan := make(chan os.Signal, 1)

	log, logFile, resolvedConf, realConfig, err := setup(stopChan)
	if err == nil {
		printInfo(log, resolvedConf, realConfig)
		go func() {
			terr := http.Serve(realConfig)
			if terr != nil {
				t.Error("failed to serve HTTP: " + terr.Error())
				return
			}
		}()
	}
	//goland:noinspection GoNilness
	return resolvedConf, logFile, stopChan, realConfig, err
}

func TestHellPot(t *testing.T) {
	tDir := filepath.Join(t.TempDir(), strconv.Itoa(int(time.Now().Unix())))
	logDir := filepath.Join(tDir, "logs")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		t.Fatal(err)
	}
	confFile := filepath.Join(tDir, "HellPot_test.toml")
	t.Setenv("HELLPOT_LOGGER_DIRECTORY", logDir)
	t.Setenv("HELLPOT_LOGGER_RSYSLOG_ADDRESS", "local")
	t.Setenv("HELLPOT_CONFIG", confFile)

	resolvedConf, logFile, stopChan, realConfig, err := testMain(t)
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
	time.Sleep(25 * time.Millisecond) // sync maybe
	logDat, err := os.ReadFile(logFile)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(logDat), "🔥 Starting HellPot 🔥") {
		t.Errorf("expected log to contain '🔥 Starting HellPot 🔥', got %s", logDat)
	}
	if !strings.Contains(string(logDat), logFile) {
		t.Errorf("expected log to contain '%s'", logFile)
	}
	if !strings.Contains(string(logDat), resolvedConf) {
		t.Errorf("expected log to contain '%s'", resolvedConf)
	}
	if !strings.Contains(string(logDat), "PID: "+strconv.Itoa(os.Getpid())) {
		t.Errorf("expected log to contain 'PID: %d', got %s", os.Getpid(), logDat)
	}
	t.Log("resolvedConf: ", resolvedConf)
	t.Log("logFile: ", logFile)
	t.Log("realConfig: ", realConfig)
	time.Sleep(100 * time.Millisecond)
	stopChan <- os.Interrupt
}
