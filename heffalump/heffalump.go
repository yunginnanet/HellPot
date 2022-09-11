/*
Package heffalump attempts to encapsulate the original work by carlmjohnson on heffalump
https://github.com/carlmjohnson/heffalump
*/
package heffalump

import (
	"bufio"
	"io"
	"sync"

	"github.com/yunginnanet/HellPot/internal/config"
)

var log = config.GetLogger()

// DefaultHeffalump represents a Heffalump type
var DefaultHeffalump *Heffalump

// Heffalump represents our buffer pool and markov map from Heffalump
type Heffalump struct {
	pool     *sync.Pool
	buffsize int
	mm       MarkovMap
}

// NewHeffalump instantiates a new Heffalump for markov generation and buffer/io operations
func NewHeffalump(mm MarkovMap, buffsize int) *Heffalump {
	return &Heffalump{
		pool: &sync.Pool{New: func() interface{} {
			b := make([]byte, buffsize)
			return b
		}},
		buffsize: buffsize,
		mm:       mm,
	}
}

// WriteHell writes markov chain heffalump hell to the provided io.Writer
func (h *Heffalump) WriteHell(bw *bufio.Writer) (int64, error) {
	var n int64
	var err error

	defer func() {
		if r := recover(); r != nil {
			log.Error().Interface("caller", r).Msg("panic recovered!")
		}
	}()

	buf := h.pool.Get().([]byte)
	defer h.pool.Put(buf)

	if _, err = bw.WriteString("<html>\n<body>\n"); err != nil {
		return n, err
	}
	if n, err = io.CopyBuffer(bw, h.mm, buf); err != nil {
		return n, nil
	}

	return n, nil
}
