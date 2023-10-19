package util

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

var ErrLimitReached = errors.New("limit reached")

// Speedometer is an io.Writer wrapper that will limit the rate at which data is written to the underlying target.
//
// It is safe for concurrent use, but writers will block when slowed down.
//
// Optionally, it can be given;
//
//   - a capacity, which will cause it to return an error if the capacity is exceeded.
//
//   - a speed limit, causing slow downs of data written to the underlying writer if the speed limit is exceeded.
type Speedometer struct {
	ceiling    int64
	speedLimit *SpeedLimit
	internal   atomics
	w          io.Writer
}

type atomics struct {
	count    *atomic.Int64
	closed   *atomic.Bool
	start    *sync.Once
	stop     *sync.Once
	birth    *atomic.Pointer[time.Time]
	duration *atomic.Pointer[time.Duration]
	slow     *atomic.Bool
}

func newAtomics() atomics {
	manhattan := atomics{
		count:    new(atomic.Int64),
		closed:   new(atomic.Bool),
		start:    new(sync.Once),
		stop:     new(sync.Once),
		birth:    new(atomic.Pointer[time.Time]),
		duration: new(atomic.Pointer[time.Duration]),
		slow:     new(atomic.Bool),
	}
	manhattan.birth.Store(&time.Time{})
	manhattan.closed.Store(false)
	manhattan.count.Store(0)
	return manhattan
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

const fallbackDelay = 100

func regulateSpeedLimit(speedLimit *SpeedLimit) (*SpeedLimit, error) {
	if speedLimit.Burst <= 0 || speedLimit.Frame <= 0 {
		return nil, errors.New("invalid speed limit")
	}
	if speedLimit.CheckEveryBytes <= 0 {
		speedLimit.CheckEveryBytes = speedLimit.Burst
	}
	if speedLimit.Delay <= 0 {
		speedLimit.Delay = fallbackDelay * time.Millisecond
	}
	return speedLimit, nil
}

func newSpeedometer(w io.Writer, speedLimit *SpeedLimit, ceiling int64) (*Speedometer, error) {
	if w == nil {
		return nil, errors.New("writer cannot be nil")
	}
	var err error
	if speedLimit != nil {
		if speedLimit, err = regulateSpeedLimit(speedLimit); err != nil {
			return nil, err
		}
	}

	return &Speedometer{
		w:          w,
		ceiling:    ceiling,
		speedLimit: speedLimit,
		internal:   newAtomics(),
	}, nil
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
func NewCappedSpeedometer(w io.Writer, capacity int64) (*Speedometer, error) {
	return newSpeedometer(w, nil, capacity)
}

// NewCappedLimitedSpeedometer creates a new Speedometer that wraps the given io.Writer.
// It is a combination of NewLimitedSpeedometer and NewCappedSpeedometer.
func NewCappedLimitedSpeedometer(w io.Writer, speedLimit *SpeedLimit, capacity int64) (*Speedometer, error) {
	return newSpeedometer(w, speedLimit, capacity)
}

func (s *Speedometer) increment(inc int64) (int, error) {
	if s.internal.closed.Load() {
		return 0, io.ErrClosedPipe
	}
	var err error
	if s.ceiling > 0 && s.Total()+inc > s.ceiling {
		_ = s.Close()
		err = ErrLimitReached
		inc = s.ceiling - s.Total()
	}
	s.internal.count.Add(inc)
	return int(inc), err
}

// Running returns true if the Speedometer is still running.
func (s *Speedometer) Running() bool {
	return !s.internal.closed.Load()
}

// Total returns the total number of bytes written to the underlying writer.
func (s *Speedometer) Total() int64 {
	return s.internal.count.Load()
}

// Close stops the Speedometer. No additional writes will be accepted.
func (s *Speedometer) Close() error {
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
	if s.internal.closed.Load() {
		return 0, io.ErrClosedPipe
	}
	s.internal.start.Do(func() {
		now := time.Now()
		s.internal.birth.Store(&now)
	})

	// if no speed limit, just write and record
	if s.speedLimit == nil {
		n, err = s.w.Write(p)
		if err != nil {
			return n, fmt.Errorf("error writing to underlying writer: %w", err)
		}
		return s.increment(int64(len(p)))
	}

	var (
		wErr     error
		accepted int
	)
	accepted, wErr = s.increment(int64(len(p)))

	if wErr != nil {
		return 0, fmt.Errorf("error incrementing: %w", wErr)
	}

	if sErr := s.slowDown(); sErr != nil {
		return 0, fmt.Errorf("error slowing down: %w", sErr)
	}

	var iErr error
	if n, iErr = s.w.Write(p[:accepted]); iErr != nil {
		return n, fmt.Errorf("error writing to underlying writer: %w", iErr)
	}
	return
}
