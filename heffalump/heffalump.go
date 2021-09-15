/*
Package heffalump attempts to encapsulate the original work by carlmjohnson on heffalump
https://github.com/carlmjohnson/heffalump
 */
package heffalump

import (
	"io"
	"sync"

	"github.com/yunginnanet/HellPot/config"
)

var log = config.GetLogger()

// DefaultHeffalump represents a Heffalump type
var DefaultHeffalump = NewHeffalump(DefaultMarkovMap, 100*1<<10)

// Heffalump represents our buffer pool and markov map from Heffalump
// https://github.com/carlmjohnson/heffalump
type Heffalump struct {
	pool     sync.Pool
	buffsize int
	mm       MarkovMap
}

// NewHeffalump instantiates a new Heffalump for markov generation and buffer/io operations
// https://github.com/carlmjohnson/heffalump
func NewHeffalump(mm MarkovMap, buffsize int) *Heffalump {
	return &Heffalump{
		pool:     sync.Pool{},
		buffsize: buffsize,
		mm:       mm,
	}
}

func (h *Heffalump) getBuffer() []byte {
	x := h.pool.Get()
	if buf, ok := x.([]byte); ok {
		return buf
	}
	return make([]byte, h.buffsize)
}

func (h *Heffalump) putBuffer(buf []byte) {
	h.pool.Put(buf)
}

// WriteHell writes markov chain heffalump hell to the provided io.Writer
// https://github.com/carlmjohnson/heffalump
func (h *Heffalump) WriteHell(w io.Writer) int64 {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Interface("caller", r).Msg("panic recovered!")
		}
	}()

	buf := h.getBuffer()
	defer h.putBuffer(buf)

	if _, err := io.WriteString(w, "<HTML>\n<BODY>\n"); err != nil {
		log.Debug().Caller().Err(err).Msg("WriteString_fail")
	}

	var n int64
	var err error
	if n, err = io.CopyBuffer(w, h.mm, buf); err != nil {
		log.Debug().Caller().Err(err).Msg("CopyBuffer_fail")
	}

	return n
}
