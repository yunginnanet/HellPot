package toml

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func (d *decoder) handleLine(lines []string, i int, myTable map[string]interface{}) (err error, readMore bool) {
	var key string
	var value interface{}

	// TODO: does this account for multiline strings?
	split := strings.SplitN(lines[i], "=", 2) //nolint:gomnd
	if len(split) != 2 {
		return nil, false
	}

	key = strings.TrimSpace(split[0])

	valueString := strings.TrimSpace(split[1])

	if valueString == "[" {
		d.inArray = true
		var continued []byte
		if continued, err = d.buf.ReadBytes(']'); err != nil {
			return fmt.Errorf("%w: %w in mthe middle of array %w", ErrInvalidTOML, io.ErrUnexpectedEOF, err), false
		}
		valueString += string(continued)
		valueString = strings.ReplaceAll(valueString, "\n", "")
	}

	switch {
	case d.inArray:
		value = handleArray(valueString)
		d.inArray = false
		readMore = true
	default:
		value = handleValue(valueString)
	}

	if value == nil || key == "" || len(valueString) == 0 {
		return
	}

	fmt.Printf("key: %s, value (%T): %v\n", key, value, value)

	myTable[key] = value

	return
}

func (d *decoder) readData() error {
	d.buf = bufs.Get()
	defer bufs.MustPut(d.buf)

	n, err := d.buf.ReadFrom(d.utf8)
	if err == nil && n == 0 {
		err = io.ErrUnexpectedEOF
	}
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidTOML, err)
	}

	for d.buf.Len() > 0 {
		var tableName []byte
		var tableNameErr error

		if tableName, tableNameErr = d.nextTableName(); tableNameErr != nil && !errors.Is(tableNameErr, io.EOF) {
			return tableNameErr
		}

		myTable, mapAssertOK := d.tables[string(tableName)].(map[string]interface{})

		if !mapAssertOK {
			d.tables[string(tableName)] = make(map[string]interface{})
			myTable = d.tables[string(tableName)].(map[string]interface{}) //nolint:forcetypeassert
		}

	readTable:

		var tableData []byte
		if tableData, err = d.buf.ReadBytes('['); err != nil && !errors.Is(err, io.EOF) {
			return fmt.Errorf("%w: %w", ErrInvalidTOML, err)
		}

		tableDataStr := string(tableData)

		lines := strings.Split(tableDataStr, "\n")

		if strings.Contains(lines[len(lines)-1], "[") && !strings.Contains(lines[len(lines)-1], "=") {
			_ = d.buf.UnreadByte()
		}

		var readMore bool

		for i := 0; i < len(lines); i++ {
			if err, readMore = d.handleLine(lines, i, myTable); err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return err
			}
			if readMore {
				goto readTable
			}
		}

		if len(d.tables) == 0 {
			return fmt.Errorf("%w: empty TOML, %w", ErrInvalidTOML, io.ErrUnexpectedEOF)
		}
	}
	return nil
}
