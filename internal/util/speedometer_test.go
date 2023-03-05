package util

import (
	"bufio"
	"errors"
	"io"
	"testing"
	"time"
)

func writeStuff(target *bufio.Writer, count int) error {
	write := func() error {
		_, err := target.WriteString("a")
		if err != nil {
			return err
		}
		return nil
	}

	if count < 0 {
		for {
			if err := write(); err != nil {
				return err
			}
		}
	}
	for i := 0; i < count; i++ {
		if err := write(); err != nil {
			return err
		}
	}
	return nil
}

func Test_Speedometer(t *testing.T) {
	type results struct {
		total   int64
		written int
		rate    float64
		err     error
	}

	isIt := func(want, have results) {
		if have.total != want.total {
			t.Errorf("total: want %d, have %d", want.total, have.total)
		}
		if have.written != want.written {
			t.Errorf("written: want %d, have %d", want.written, have.written)
		}
		if have.rate != want.rate {
			t.Errorf("rate: want %f, have %f", want.rate, have.rate)
		}
		if have.err != want.err {
			t.Errorf("wantErr: want %v, have %v", want.err, have.err)
		}
	}

	sp := NewSpeedometer(io.Discard)
	errChan := make(chan error)
	go func() {
		errChan <- writeStuff(sp.BufioWriter, -1)
	}()
	time.Sleep(1 * time.Second)
	_ = sp.Close()
	err := <-errChan
	cnt, err := sp.Write([]byte("a"))
	isIt(results{err: io.ErrClosedPipe, written: 0}, results{err: err, written: cnt})
	sp.Reset()
	cnt, err = sp.Write([]byte("a"))
	isIt(results{err: nil, written: 1, total: 1}, results{err: err, written: cnt, total: sp.Total()})
	cnt, err = sp.Write([]byte("aa"))
	isIt(results{err: nil, written: 2, total: 3}, results{err: err, written: cnt, total: sp.Total()})
	cnt, err = sp.Write([]byte("a"))
	isIt(results{err: nil, written: 1, total: 4}, results{err: err, written: cnt, total: sp.Total()})
	cnt, err = sp.Write([]byte("a"))
	isIt(results{err: nil, written: 1, total: 5}, results{err: err, written: cnt, total: sp.Total()})
	go func() {
		errChan <- writeStuff(sp.BufioWriter, -1)
	}()
	var count = 0
	for sp.Running() {
		select {
		case err = <-errChan:
			if !errors.Is(err, io.ErrClosedPipe) {
				t.Errorf("unexpected error: %v", err)
			} else {
				if count < 5 {
					t.Errorf("too few iterations: %d", count)
				}
				t.Logf("final rate: %v per second", sp.Rate())
			}
		default:
			if count > 5 {
				_ = sp.Close()
			}
			t.Logf("rate: %v per second", sp.Rate())
			time.Sleep(1 * time.Second)
			count++
		}
	}
}
