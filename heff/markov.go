package heff

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ScanHTML is a basic split function for a Scanner that returns each
// space-separated word of text or HTML tag, with surrounding spaces deleted.
// It will never return an empty string. The definition of space is set by
// unicode.IsSpace.
func ScanHTML(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	var r rune
	start := 0
	for width := 0; start < len(data); start += width {
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}
	if r == '<' {
		// Scan until closing bracket
		for i := start; i < len(data); i++ {
			if data[i] == '>' {
				return i + 1, data[start : i+1], nil
			}
		}
	} else {
		// Scan until space, marking end of word.
		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if unicode.IsSpace(r) {
				return i + width, data[start:i], nil
			}
			if r == '<' {
				return i, data[start:i], nil
			}
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

const (
	nprefix = 2
	nonword = "\n"
)

type tokenPair [nprefix]string

var markov = make(map[tokenPair][]string)

func init() {
	var w1, w2 = nonword, nonword
	var p tokenPair

	s := bufio.NewScanner(strings.NewReader(Src))
	s.Split(ScanHTML)
	for s.Scan() {
		t := s.Text()
		p = tokenPair{w1, w2}
		markov[p] = append(markov[p], t)
		w1, w2 = w2, t
	}

	p = tokenPair{w1, w2}
	markov[p] = append(markov[p], nonword)
}

func Generate(w io.Writer, maxgen int) {
	var w1, w2 = nonword, nonword

	for i := 0; i < maxgen; i++ {
		p := tokenPair{w1, w2}
		suffix := markov[p]
		r := rand.Intn(len(suffix))
		t := suffix[r]
		if t == nonword {
			break
		}
		fmt.Fprint(w, t, "\n")
		w1, w2 = w2, t
	}
}
