/*
Package heffalump attempts to encapsulate the original work by carlmjohnson on heffalump
https://github.com/carlmjohnson/heffalump
*/
package heffalump

import (
	"bufio"
	"io"
	"sync"
)

const DefaultBuffSize = 100 * 1 << 10

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

// NewDefaultHeffalump instantiates a new default Heffalump from a MarkovMap created using
// using the default source text.
func NewDefaultHeffalump() *Heffalump {
	return NewHeffalump(NewDefaultMarkovMap(), DefaultBuffSize)
}

// WriteHell writes markov chain heffalump hell to the provided io.Writer
func (h *Heffalump) WriteHell(bw *bufio.Writer) (int64, error) {
	var n int64
	var err error

	/*	defer func() {
		if r := recover(); r != nil {
			log.Error().Interface("caller", r).Msg("panic recovered!")
		}
	}()*/

	buf := h.pool.Get().([]byte)

	if _, err = bw.WriteString("<html>\n<body>\n"); err != nil {
		h.pool.Put(buf)
		return n, err
	}
	if n, err = io.CopyBuffer(bw, h.mm, buf); err != nil {
		h.pool.Put(buf)
		return n, nil
	}

	h.pool.Put(buf)
	return n, nil
}
