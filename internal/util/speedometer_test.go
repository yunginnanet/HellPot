package util

import (
	"errors"
	"io"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type testWriter struct {
	t     *testing.T
	total int64
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	atomic.AddInt64(&w.total, int64(len(p)))
	return len(p), nil
}

func writeStuff(target io.Writer, count int) error {
	write := func() error {
		_, err := target.Write([]byte("a"))
		if err != nil {
			return err
		}
		return nil
	}

	if count < 0 {
		var err error
		for err = write(); err == nil; err = write() {
			time.Sleep(5 * time.Millisecond)
		}
		return err
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

	var (
		errChan = make(chan error, 10)
		err     error
		cnt     int
	)

	t.Run("EarlyClose", func(t *testing.T) {
		sp, nerr := NewSpeedometer(&testWriter{t: t})
		if nerr != nil {
			t.Errorf("unexpected error: %v", nerr)
		}
		go func() {
			errChan <- writeStuff(sp, -1)
		}()
		time.Sleep(1 * time.Second)
		if closeErr := sp.Close(); closeErr != nil {
			t.Errorf("wantErr: want %v, have %v", nil, closeErr)
		}
		err = <-errChan
		if !errors.Is(err, io.ErrClosedPipe) {
			t.Errorf("wantErr: want %v, have %v", io.ErrClosedPipe, err)
		}
		cnt, err = sp.Write([]byte("a"))
		isIt(results{err: io.ErrClosedPipe, written: 0}, results{err: err, written: cnt})
	})

	t.Run("Basic", func(t *testing.T) {
		sp, nerr := NewSpeedometer(&testWriter{t: t})
		if nerr != nil {
			t.Errorf("unexpected error: %v", nerr)
		}
		cnt, err = sp.Write([]byte("a"))
		isIt(results{err: nil, written: 1, total: 1}, results{err: err, written: cnt, total: sp.Total()})
		cnt, err = sp.Write([]byte("aa"))
		isIt(results{err: nil, written: 2, total: 3}, results{err: err, written: cnt, total: sp.Total()})
		cnt, err = sp.Write([]byte("a"))
		isIt(results{err: nil, written: 1, total: 4}, results{err: err, written: cnt, total: sp.Total()})
		cnt, err = sp.Write([]byte("a"))
		isIt(results{err: nil, written: 1, total: 5}, results{err: err, written: cnt, total: sp.Total()})
	})

	t.Run("ConcurrentWrites", func(t *testing.T) {
		count := int64(0)
		sp, nerr := NewSpeedometer(&testWriter{t: t})
		if nerr != nil {
			t.Errorf("unexpected error: %v", nerr)
		}
		wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				var counted int
				var err error
				counted, err = sp.Write([]byte("a"))
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				atomic.AddInt64(&count, int64(counted))
				wg.Done()
			}()
		}
		wg.Wait()
		isIt(results{err: nil, written: 1, total: 100}, results{err: err, written: cnt, total: sp.Total()})
	})

	t.Run("GottaGoFast", func(t *testing.T) {
		sp, nerr := NewSpeedometer(&testWriter{t: t})
		if nerr != nil {
			t.Errorf("unexpected error: %v", nerr)
		}
		go func() {
			errChan <- writeStuff(sp, -1)
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
				time.Sleep(100 * time.Millisecond)
				t.Logf("rate: %v per second", sp.Rate())
				count++
			}
		}
	})

	// test limiter with speedlimit
	t.Run("CantGoFast", func(t *testing.T) {
		t.Run("10BytesASecond", func(t *testing.T) {
			sp, nerr := NewLimitedSpeedometer(&testWriter{t: t}, &SpeedLimit{
				Burst:           10,
				Frame:           time.Second,
				CheckEveryBytes: 1,
				Delay:           100 * time.Millisecond,
			})
			if nerr != nil {
				t.Errorf("unexpected error: %v", nerr)
			}
			for i := 0; i < 15; i++ {
				if _, err = sp.Write([]byte("a")); err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				/*if sp.IsSlow() {
					t.Errorf("unexpected slow state")
				}*/
				t.Logf("rate: %v per second", sp.Rate())
				if sp.Rate() > 10 {
					t.Errorf("speeding in a school zone (expected under %d): %v", sp.speedLimit.Burst, sp.Rate())
				}
			}
		})

		t.Run("1000BytesPer5SecondsMeasuredEvery5000Bytes", func(t *testing.T) {
			sp, nerr := NewLimitedSpeedometer(&testWriter{t: t}, &SpeedLimit{
				Burst:           1000,
				Frame:           2 * time.Second,
				CheckEveryBytes: 5000,
				Delay:           500 * time.Millisecond,
			})

			if nerr != nil {
				t.Errorf("unexpected error: %v", nerr)
			}

			for i := 0; i < 4999; i++ {
				if _, err = sp.Write([]byte("a")); err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if i%1000 == 0 {
					t.Logf("rate: %v per second", sp.Rate())
				}
				if sp.Rate() < 1000 {
					t.Errorf("shouldn't have slowed down yet (expected over %d): %v", sp.speedLimit.Burst, sp.Rate())
				}
			}
			if _, err = sp.Write([]byte("a")); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			for i := 0; i < 10; i++ {
				if _, err = sp.Write([]byte("a")); err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				t.Logf("rate: %v per second", sp.Rate())
				if sp.Rate() > 1000 {
					t.Errorf("speeding in a school zone (expected under %d): %v", sp.speedLimit.Burst, sp.Rate())
				}
			}
		})
	})

	// test capped speedometer
	t.Run("OnlyALittle", func(t *testing.T) {
		sp, nerr := NewCappedSpeedometer(&testWriter{t: t}, 1024)
		if nerr != nil {
			t.Errorf("unexpected error: %v", nerr)
		}
		for i := 0; i < 1024; i++ {
			if _, err = sp.Write([]byte("a")); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if sp.Total() > 1024 {
				t.Errorf("shouldn't have written more than 1024 bytes")
			}
		}
		if _, err = sp.Write([]byte("a")); err == nil {
			t.Errorf("expected error when writing over capacity")
		}
	})

}
