package toml

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"

	"git.tcp.direct/kayos/common/pool"
	"golang.org/x/text/encoding/unicode"
)

var (
	ErrInvalidTOML       = errors.New("invalid TOML")
	ErrBadTarget         = errors.New("bad target")
	ErrMismatchingSchema = errors.New("mismatching schema types")
)

type decoder struct {
	buf     *pool.Buffer
	utf8    io.Reader
	target  interface{}
	inArray bool
	tables  map[string]interface{}
}

func newDecoder(input []byte, v interface{}) *decoder {
	d := &decoder{
		target: v,
		utf8:   unicode.UTF8.NewDecoder().Reader(bytes.NewReader(input)),
		tables: make(map[string]interface{}),
	}
	return d
}

func (d *decoder) start() error {
	if err := d.readData(); err != nil {
		return err
	}
	return d.handleTables()
}

func trimTest(data []byte) (trimmed []byte) {
	for _, c := range []byte{' ', '\n', '\t', '\r', '[', ']', '='} {
		trimmed = bytes.ReplaceAll(data, []byte{c}, []byte(""))
	}
	return
}

func UnmarshalTOML(data []byte, v interface{}) error {
	var err error
	if data, err = unicode.UTF8.NewDecoder().Bytes(data); err != nil {
		return fmt.Errorf("unicode decode failure: %w", err)
	}

	if !bytes.Contains(data, []byte("[")) {
		return fmt.Errorf("%w: missing '['", ErrInvalidTOML)
	}
	if !bytes.Contains(data, []byte("]")) {
		return fmt.Errorf("%w: missing ']'", ErrInvalidTOML)
	}
	if !bytes.Contains(data, []byte("=")) {
		return fmt.Errorf("%w: missing '='", ErrInvalidTOML)
	}
	if len(trimTest(data)) < 2 {
		return fmt.Errorf("%w: empty TOML", ErrInvalidTOML)
	}

	ref := reflect.ValueOf(v)
	if ref.Kind() != reflect.Ptr {
		return fmt.Errorf("%w: non-pointer", ErrBadTarget)
	}
	if ref.IsNil() {
		return fmt.Errorf("%w: nil pointer", ErrBadTarget)
	}
	if ref.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("%w: non-struct", ErrBadTarget)
	}

	data = bytes.ReplaceAll(data, []byte("\r"), []byte(""))

	return newDecoder(data, v).start()
}
