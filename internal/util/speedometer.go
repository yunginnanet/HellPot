package util

import (
	"errors"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

var ErrLimitReached = errors.New("limit reached")

// Speedometer is a wrapper around an io.Writer that will limit the rate at which data is written to the underlying writer.
//
// It is safe for concurrent use, but writers will block when slowed down.
//
// Optionally, it can be given;
//
//   - a capacity, which will cause it to return an error if the capacity is exceeded.
//
//   - a speed limit, causing slow downs of data written to the underlying writer if the speed limit is exceeded.
type Speedometer struct {
	cap        int64
	speedLimit *SpeedLimit
	internal   atomics
	hardLock   *sync.RWMutex
	w          io.Writer
}

type atomics struct {
	closed   *atomic.Bool
	count    *int64
	start    *sync.Once
	stop     *sync.Once
	birth    *atomic.Pointer[time.Time]
	duration *atomic.Pointer[time.Duration]
	slow     *atomic.Bool
}

// SpeedLimit is used to limit the rate at which data is written to the underlying writer.
type SpeedLimit struct {
	// Burst is the number of bytes that can be written to the underlying writer per Frame.
	Burst int
	// Frame is the duration of the frame in which Burst can be written to the underlying writer.
	Frame time.Duration
	// CheckEveryBytes is the number of bytes written before checking if the speed limit has been exceeded.
	CheckEveryBytes int
	// Delay is the duration to delay writing if the speed limit has been exceeded during a Write call. (blocking)
	Delay time.Duration
}

func regulateSpeedLimit(speedLimit *SpeedLimit) (*SpeedLimit, error) {
	if speedLimit.Burst <= 0 || speedLimit.Frame <= 0 {
		return nil, errors.New("invalid speed limit")
	}
	if speedLimit.CheckEveryBytes <= 0 {
		speedLimit.CheckEveryBytes = speedLimit.Burst
	}
	if speedLimit.Delay <= 0 {
		speedLimit.Delay = 100 * time.Millisecond
	}
	return speedLimit, nil
}

func newSpeedometer(w io.Writer, speedLimit *SpeedLimit, cap int64) (*Speedometer, error) {
	if w == nil {
		return nil, errors.New("writer cannot be nil")
	}
	var err error
	if speedLimit != nil {
		if speedLimit, err = regulateSpeedLimit(speedLimit); err != nil {
			return nil, err
		}
	}
	z := int64(0)
	speedo := &Speedometer{
		w:          w,
		cap:        cap,
		speedLimit: speedLimit,
		hardLock:   &sync.RWMutex{},
		internal: atomics{
			count:    &z,
			birth:    new(atomic.Pointer[time.Time]),
			duration: new(atomic.Pointer[time.Duration]),
			closed:   new(atomic.Bool),
			stop:     new(sync.Once),
			start:    new(sync.Once),
			slow:     new(atomic.Bool),
		},
	}
	speedo.internal.birth.Store(&time.Time{})
	speedo.internal.closed.Store(false)
	return speedo, nil
}

// NewSpeedometer creates a new Speedometer that wraps the given io.Writer.
// It will not limit the rate at which data is written to the underlying writer, it only measures it.
func NewSpeedometer(w io.Writer) (*Speedometer, error) {
	return newSpeedometer(w, nil, -1)
}

// NewLimitedSpeedometer creates a new Speedometer that wraps the given io.Writer.
// If the speed limit is exceeded, writes to the underlying writer will be limited.
// See SpeedLimit for more information.
func NewLimitedSpeedometer(w io.Writer, speedLimit *SpeedLimit) (*Speedometer, error) {
	return newSpeedometer(w, speedLimit, -1)
}

// NewCappedSpeedometer creates a new Speedometer that wraps the given io.Writer.
// If len(written) bytes exceeds cap, writes to the underlying writer will be ceased permanently for the Speedometer.
func NewCappedSpeedometer(w io.Writer, cap int64) (*Speedometer, error) {
	return newSpeedometer(w, nil, cap)
}

func (s *Speedometer) increment(inc int64) (int, error) {
	if s.internal.closed.Load() || !s.hardLock.TryRLock() {
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

// Running returns true if the Speedometer is still running.
func (s *Speedometer) Running() bool {
	return !s.internal.closed.Load()
}

// Total returns the total number of bytes written to the underlying writer.
func (s *Speedometer) Total() int64 {
	return atomic.LoadInt64(s.internal.count)
}

// Close stops the Speedometer. No additional writes will be accepted.
func (s *Speedometer) Close() error {
	s.hardLock.TryLock()
	if s.internal.closed.Load() {
		return io.ErrClosedPipe
	}
	s.internal.stop.Do(func() {
		s.internal.closed.Store(true)
		stopped := time.Now()
		birth := s.internal.birth.Load()
		duration := stopped.Sub(*birth)
		s.internal.duration.Store(&duration)
	})
	return nil
}

/*func (s *Speedometer) IsSlow() bool {
	return s.internal.slow.Load()
}*/

// Rate returns the rate at which data is being written to the underlying writer per second.
func (s *Speedometer) Rate() float64 {
	if s.internal.closed.Load() {
		return float64(s.Total()) / s.internal.duration.Load().Seconds()
	}
	return float64(s.Total()) / time.Since(*s.internal.birth.Load()).Seconds()
}

func (s *Speedometer) slowDown() error {
	switch {
	case s.speedLimit == nil:
		return nil
	case s.speedLimit.Burst <= 0 || s.speedLimit.Frame <= 0,
		s.speedLimit.CheckEveryBytes <= 0, s.speedLimit.Delay <= 0:
		return errors.New("invalid speed limit")
	default:
		//
	}
	if s.Total()%int64(s.speedLimit.CheckEveryBytes) != 0 {
		return nil
	}
	s.internal.slow.Store(true)
	for s.Rate() > float64(s.speedLimit.Burst)/s.speedLimit.Frame.Seconds() {
		time.Sleep(s.speedLimit.Delay)
	}
	s.internal.slow.Store(false)
	return nil
}

// Write writes p to the underlying writer, following all defined speed limits.
func (s *Speedometer) Write(p []byte) (n int, err error) {
	if !s.hardLock.TryRLock() {
		return 0, io.ErrClosedPipe
	}
	s.internal.start.Do(func() {
		now := time.Now()
		s.internal.birth.Store(&now)
	})
	accepted, err := s.increment(int64(len(p)))
	if err != nil {
		wn, innerErr := s.w.Write(p[:accepted])
		if innerErr != nil {
			err = innerErr
		}
		return wn, err
	}
	if err = s.slowDown(); err != nil {
		return 0, err
	}
	return s.w.Write(p)
}
