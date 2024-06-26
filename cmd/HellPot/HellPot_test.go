package main

import (
	"os"
	"testing"
	"time"

	"github.com/yunginnanet/HellPot/internal/http"
)

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

func TestHellPot(t *testing.T) {
	resolvedConf, logFile, stopChan, err := testMain(t)
	if err != nil {
		t.Fatal(err)
	}
	if stopChan == nil {
		t.Fatal("stopChan is nil")
	}
	t.Log("resolvedConf: ", resolvedConf)
	t.Log("logFile: ", logFile)
	time.Sleep(100 * time.Millisecond)
	stopChan <- os.Interrupt
}
