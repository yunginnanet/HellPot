package main

import (
	"os"
	"testing"
	"time"
)

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
