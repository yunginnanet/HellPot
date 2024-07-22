package toml

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"git.tcp.direct/kayos/common/pool"
	"github.com/davecgh/go-spew/spew"
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

func (d *decoder) nextTableName() (tableName []byte, err error) {
	if _, err = d.buf.ReadBytes('['); err != nil {
		return nil, fmt.Errorf("%w (%w): %w", ErrInvalidTOML, io.ErrUnexpectedEOF, err)
	}

	if tableName, err = d.buf.ReadBytes(']'); err != nil {
		if errors.Is(err, io.EOF) {
			err = fmt.Errorf("%w: missing ']'", io.ErrUnexpectedEOF)
		}
		return nil, fmt.Errorf("%w: %w", ErrInvalidTOML, err)
	}

	tableName = tableName[:len(tableName)-1]

	return tableName, nil
}

func handleBool(v any) (bool, error) {
	vBool, boolOK := v.(bool)
	vInt, intOK := v.(int64)
	vStr, strOK := v.(string)
	switch {
	case intOK:
		if vInt == 0 || vInt == 1 {
			vBool = vInt == 1
			boolOK = true
		}
	case strOK:
		parsedBool, parseErr := strconv.ParseBool(vStr)
		if parseErr != nil {
			return false, fmt.Errorf("%w: expected bool-ish: %w", ErrMismatchingSchema, parseErr)
		}
		vBool = parsedBool
		boolOK = true
	case boolOK:
	default:
	}
	if !boolOK {
		return false, fmt.Errorf("%w: expected bool, got %T", ErrMismatchingSchema, v)
	}

	return vBool, nil
}

func handleSlice(targetElem reflect.Value, j int, v any) error {
	switch field := targetElem; field.Type().Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vInt, intOK := v.([]int64)
		if !intOK {
			return fmt.Errorf("%w: expected int slice, got %T", ErrMismatchingSchema, v)
		}
		slice := reflect.MakeSlice(field.Type().Elem(), len(vInt), len(vInt))
		for i := 0; i < len(vInt); i++ {
			slice.Index(i).SetInt(vInt[i])
		}
		targetElem.Field(j).Set(slice)
	case reflect.String:
		vStr, strOK := v.([]string)
		if !strOK {
			return fmt.Errorf("%w: expected string slice, got %T", ErrMismatchingSchema, v)
		}
		slice := reflect.MakeSlice(field.Type().Elem(), len(vStr), len(vStr))
		for i := 0; i < len(vStr); i++ {
			slice.Index(i).SetString(vStr[i])
		}
		targetElem.Field(j).Set(slice)
	default:
		return fmt.Errorf("%w: unsupported slice type %s", ErrMismatchingSchema, field.Type().Elem().Kind().String())
	}
	return nil
}

func handleBottom(targetElem reflect.Value, v any, j int) error {
	if !targetElem.CanSet() {
		return fmt.Errorf("%w: cannot set field", ErrBadTarget)
	}

	println(targetElem.Type().Kind().String())

	switch targetElem.Kind() {
	case reflect.String:
		vStr, strOK := v.(string)
		if !strOK {
			return fmt.Errorf("%w: expected string, got %T", ErrMismatchingSchema, v)
		}
		targetElem.SetString(vStr)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vInt, intOK := v.(int64)
		if !intOK {
			return fmt.Errorf("%w: expected int, got %T", ErrMismatchingSchema, v)
		}
		targetElem.SetInt(vInt)
	case reflect.Slice:
		if sliceErr := handleSlice(targetElem, j, v); sliceErr != nil {
			return sliceErr
		}
	case reflect.Bool:
		vBool, boolErr := handleBool(v)
		if boolErr != nil {
			return boolErr
		}
		targetElem.Field(j).SetBool(vBool)
	}
	return nil
}

func handleArray(valueString string) (value any) {
	valueString = strings.Trim(valueString, "[]")
	asplit := strings.Split(valueString, ",")
	intVals := make([]int64, 0, len(asplit))
	for aindex := range asplit {
		asplit[aindex] = strings.Trim(strings.TrimSpace(asplit[aindex]), "\"'")
		avalInt, intErr := strconv.ParseInt(valueString, 10, 64)
		if intErr == nil {
			intVals = append(intVals, avalInt)
		}
	}

	if len(intVals) == len(asplit) && len(intVals) > 0 {
		value = intVals
	} else {
		value = asplit
	}

	return value
}

func handleValue(valueString string) (value any) {
	valueString = strings.Trim(valueString, "\"'")
	valBool, boolErr := strconv.ParseBool(valueString)
	valInt, intErr := strconv.ParseInt(valueString, 10, 64)
	switch {
	case boolErr == nil:
		value = valBool
	case intErr == nil:
		value = valInt
	default:
		value = valueString
	}
	return value
}

func (d *decoder) handleLine(lines []string, i int, myTable map[string]interface{}) (err error, readMore bool) {
	var key string
	var value interface{}
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

func findStructTag(struc reflect.Value, sought string) (int, error) {
	for i := 0; i < struc.NumField(); i++ {
		if struc.Type().Field(i).Tag.Get("toml") == sought {
			return i, nil
		}
	}
	return -1, fmt.Errorf("%w: %s", ErrMismatchingSchema, sought)

}

func (d *decoder) handleTables() error {
	spew.Dump(d.tables)

	for k, v := range d.tables {
		if !strings.Contains(k, ".") {
			vTable, tableOK := v.(map[string]interface{})
			if !tableOK {
				return fmt.Errorf("%w: expected table, got %T", ErrMismatchingSchema, v)
			}
			println("table:", vTable)

			targetElem := reflect.ValueOf(d.target).Elem()

			fieldIndex, fieldErr := findStructTag(targetElem, k)
			if fieldErr != nil {
				continue
			}
			println("field index: ", fieldIndex)
			if targetElem.Field(fieldIndex).Type().Kind() != reflect.Struct {
				return fmt.Errorf("%w: expected struct, got %T", ErrMismatchingSchema, targetElem.Field(fieldIndex).Type())
			}
		}
	}

	return nil
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
