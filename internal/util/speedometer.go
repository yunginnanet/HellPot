package util

import (
	"bufio"
	"errors"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

var ErrLimitReached = errors.New("limit reached")

type Speedometer struct {
	cap         int64
	speedLimit  SpeedLimit
	birth       *time.Time
	duration    time.Duration
	internal    atomics
	w           io.Writer
	BufioWriter *bufio.Writer
}

type atomics struct {
	closed *atomic.Bool
	count  *int64
	start  *sync.Once
	stop   *sync.Once
}

func NewSpeedometer(w io.Writer) *Speedometer {
	z := int64(0)
	speedo := &Speedometer{
		w:   w,
		cap: -1,
		internal: atomics{
			count:  &z,
			closed: new(atomic.Bool),
			stop:   new(sync.Once),
			start:  new(sync.Once),
		},
	}
	speedo.internal.closed.Store(false)
	speedo.BufioWriter = bufio.NewWriter(speedo)
	return speedo
}

type SpeedLimit struct {
	Bytes           int
	Per             time.Duration
	CheckEveryBytes int
	Delay           time.Duration
}

func NewLimitedSpeedometer(w io.Writer, speedLimit SpeedLimit) *Speedometer {
	z := int64(0)
	speedo := &Speedometer{
		w:          w,
		cap:        -1,
		speedLimit: speedLimit,
		internal: atomics{
			count:  &z,
			closed: new(atomic.Bool),
			stop:   new(sync.Once),
			start:  new(sync.Once),
		},
	}
	speedo.internal.closed.Store(false)
	speedo.BufioWriter = bufio.NewWriterSize(speedo, speedLimit.CheckEveryBytes)
	return speedo
}

func NewCappedSpeedometer(w io.Writer, cap int64) *Speedometer {
	z := int64(0)
	speedo := &Speedometer{
		w:   w,
		cap: cap,
		internal: atomics{
			count:  &z,
			closed: new(atomic.Bool),
			stop:   new(sync.Once),
			start:  new(sync.Once),
		},
	}
	speedo.internal.closed.Store(false)
	speedo.BufioWriter = bufio.NewWriterSize(speedo, int(cap)/10)
	return speedo
}

func (s *Speedometer) increment(inc int64) (int, error) {
	if s.internal.closed.Load() {
		return 0, io.ErrClosedPipe
	}
	var err error
	if s.cap > 0 && s.Total()+inc > s.cap {
		_ = s.Close()
		err = ErrLimitReached
		inc = s.cap - s.Total()
	}
	atomic.AddInt64(s.internal.count, inc)
	return int(inc), err
}

func (s *Speedometer) Running() bool {
	return !s.internal.closed.Load()
}

func (s *Speedometer) Total() int64 {
	return atomic.LoadInt64(s.internal.count)
}

func (s *Speedometer) Reset() {
	s.internal.count = new(int64)
	s.internal.closed = new(atomic.Bool)
	s.internal.start = new(sync.Once)
	s.internal.stop = new(sync.Once)
	s.BufioWriter = bufio.NewWriter(s)
	s.internal.closed.Store(false)
}

func (s *Speedometer) Close() error {
	s.internal.stop.Do(func() {
		s.internal.closed.Store(true)
		stopped := time.Now()
		s.duration = stopped.Sub(*s.birth)
	})
	return nil
}

func (s *Speedometer) Rate() float64 {
	if s.internal.closed.Load() {
		return float64(s.Total()) / s.duration.Seconds()
	}
	return float64(s.Total()) / time.Since(*s.birth).Seconds()
}

func (s *Speedometer) Write(p []byte) (n int, err error) {
	s.internal.start.Do(func() {
		tn := time.Now()
		s.birth = &tn
	})
	accepted, err := s.increment(int64(len(p)))
	if err != nil {
		wn, innerErr := s.w.Write(p[:accepted])
		if innerErr != nil {
			err = innerErr
		}
		if wn < accepted {
			err = io.ErrShortWrite
		}
		return wn, err
	}
	return s.w.Write(p)
}

/*func BenchmarkHeffalump_WriteHell(b *testing.B) {

}*/
