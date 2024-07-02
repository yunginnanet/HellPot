package toml

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"git.tcp.direct/kayos/common/pool"
)

var bufs = pool.NewBufferFactory()

var ErrMissingTag = errors.New("missing toml tag")

type subStruct struct {
	name   string
	ref    reflect.Value
	parent *subStruct
}

func (sub *subStruct) writeTableHeader(buf *pool.Buffer) {
	parent := sub.parent
	var parents = make([]string, 0)
	for parent != nil {
		parents = append(parents, parent.name)
		parent = parent.parent
	}
	slices.Reverse(parents)
	parents = append(parents, sub.name)
	if len(parents) > 1 {
		writeTableHeaders(buf, parents[0], parents[1:]...)
	}
	if len(parents) == 1 {
		writeTableHeaders(buf, parents[0], parents[1:]...)
	}
}

func mustValidateTableName(tableName string) {
	for _, b := range tableName {
		if !unicode.IsLetter(b) && b != '.' && b != '_' {
			panic("table name must contain only letters, bad character: " + string(b))
		}
	}
}

func shouldSkip(field reflect.StructField) (skip bool) {
	if !field.IsExported() || field.Tag.Get("toml") == "-" {
		skip = true
		return
	}
	return
}

func writeTableHeaders(buf *pool.Buffer, tableName string, subTableNames ...string) {
	mustValidateTableName(tableName)
	_, _ = buf.WriteString("[")
	_, _ = buf.WriteString(tableName)
	if len(subTableNames) > 0 {
		for _, sub := range subTableNames {
			mustValidateTableName(sub)
			_, _ = buf.WriteString(".")
			_, _ = buf.WriteString(sub)
		}
	}
	_, _ = buf.WriteString("]")
	_, _ = buf.WriteString("\n")
}

func handlePanic(err error) error {
	var r any
	if r = recover(); r == nil {
		return err
	}
	if _, ok := r.(error); !ok {
		panic(r)
	}
	ve := &reflect.ValueError{}
	if !errors.As(r.(error), &ve) { //nolint:forcetypeassert
		panic(r)
	}
	if err == nil {
		err = ve
		return err
	}
	errs := []error{err, ve}
	err = errors.Join(errs...)
	return err
}

func handleString(s string, buf *pool.Buffer) error {
	delim := '"'
	if strings.Contains(s, "\"") {
		delim = '\''
	}
	if strings.Contains(s, "'") && delim == '\'' {
		return errors.New("string contains both single and double quotes, escaping is not supported")
	}

	_, _ = buf.WriteRune(delim)
	_, _ = buf.WriteString(s)
	_, _ = buf.WriteRune(delim)

	return nil
}

func handleBottomLevelField(field reflect.StructField, ref reflect.Value, buf *pool.Buffer) error {
	if field.Tag.Get("toml") == "" {
		return fmt.Errorf("%w: bottom level field", ErrMissingTag)
	}
	if shouldSkip(field) {
		return nil
	}
	_, _ = buf.WriteString(field.Tag.Get("toml"))
	_, _ = buf.WriteString(" = ")
	//goland:noinspection GoSwitchMissingCasesForIotaConsts
	switch ref.Kind() {
	case reflect.String:
		if err := handleString(ref.String(), buf); err != nil {
			return err
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		_, _ = buf.WriteString(strconv.FormatInt(ref.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		_, _ = buf.WriteString(strconv.FormatUint(ref.Uint(), 10))
	case reflect.Bool:
		_, _ = buf.WriteString(strconv.FormatBool(ref.Bool()))
	case reflect.Slice:
		//goland:noinspection GoSwitchMissingCasesForIotaConsts
		switch ref.Type().Elem().Kind() { //nolint:exhaustive
		case reflect.String:
			if ref.Len() == 0 {
				return nil
			}
			_, _ = buf.WriteString("[")
			for i := 0; i < ref.Len(); i++ {
				if err := handleString(ref.Index(i).String(), buf); err != nil {
					return err
				}
				if i < ref.Len()-1 {
					_, _ = buf.WriteString(", ")
				}
			}
			_, _ = buf.WriteString("]")
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if ref.Len() == 0 {
				return nil
			}
			_, _ = buf.WriteString("[")
			for i := 0; i < ref.Len(); i++ {
				_, _ = buf.WriteString(strconv.FormatInt(ref.Index(i).Int(), 10))
				if i < ref.Len()-1 {
					_, _ = buf.WriteString(", ")
				}
			}
			_, _ = buf.WriteString("]")
		}
	}

	_, _ = buf.WriteString("\n")

	return nil
}

func allChildrenEmpty(sub reflect.Value) (empty bool, err error) {
	if sub.Kind() != reflect.Struct {
		return false, errors.New("input to allChildrenEmpty must be a struct")
	}
	for i := 0; i < sub.Type().NumField(); i++ {
		if shouldSkip(sub.Type().Field(i)) {
			continue
		}
		if sub.Type().Field(i).Tag.Get("toml") == "" {
			return false, fmt.Errorf("%w: during allChildrenEmpty call (string branch)", ErrMissingTag)
		}
		val := reflect.ValueOf(sub.Field(i).Interface())
		if !val.IsZero() {
			// additionalMsg = "field '" + sub.Type().Field(i).Name + "' is not zero: " + val.Kind().String()
			return false, nil
		}
	}
	return true, nil
}

func (sub *subStruct) handleSubStructField(field reflect.StructField, ref reflect.Value, buf *pool.Buffer) error {
	if field.Tag.Get("toml") == "" {
		return fmt.Errorf("%w: sub struct field", ErrMissingTag)
	}
	if shouldSkip(field) {
		return nil
	}
	if field.Type.Kind() == reflect.Ptr {
		ref = ref.Elem()
		if !ref.IsValid() {
			return nil
		}
	}

	ft := ref.Type()

	switch ft.Kind() { //nolint:exhaustive
	case reflect.Struct:
		child := &subStruct{name: field.Tag.Get("toml"), ref: ref, parent: sub}
		if empty, err := allChildrenEmpty(ref); empty || err != nil {
			return err
		}
		_, _ = buf.WriteString("\n")
		child.writeTableHeader(buf)
		for i := 0; i < ft.NumField(); i++ {
			if err := child.handleSubStructField(ft.Field(i), ref.Field(i), buf); err != nil {
				return err
			}
		}
	default:
		if err := handleBottomLevelField(field, ref, buf); err != nil {
			return err
		}
	}

	return nil
}

func handleFieldTopLevel(field reflect.StructField, ref reflect.Value, buf *pool.Buffer) error {
	if field.Type.Kind() != reflect.Struct {
		return errors.New("input struct must contain only sub-structs")
	}
	if field.Tag.Get("toml") == "" {
		return fmt.Errorf("%w: top level field", ErrMissingTag)
	}

	fieldCount := field.Type.NumField()
	if fieldCount < 1 {
		return errors.New("sub-struct must have at least one field")
	}

	if shouldSkip(field) {
		return nil
	}

	if empty, err := allChildrenEmpty(ref); empty || err != nil {
		return err
	}

	child := &subStruct{name: field.Tag.Get("toml"), parent: nil, ref: ref}
	child.writeTableHeader(buf)

	for i := 0; i < ref.NumField(); i++ {
		if err := child.handleSubStructField(ref.Type().Field(i), ref.Field(i), buf); err != nil {
			return err
		}
	}

	return nil
}

func MarshalTOML(v interface{}) (data []byte, err error) {
	defer func() { err = handlePanic(err) }()
	tof := reflect.TypeOf(v)
	if tof.Kind() != reflect.Struct {
		return nil, errors.New("input must be a struct")
	}
	elem := reflect.ValueOf(v)
	fieldCount := elem.NumField()
	if fieldCount < 1 {
		return nil, errors.New("input struct must have at least one field")
	}
	buf := bufs.Get()
	defer bufs.MustPut(buf)
	for i := 0; i < fieldCount; i++ {
		field := tof.Field(i)
		if shouldSkip(field) {
			continue
		}
		if err = handleFieldTopLevel(field, elem.Field(i), buf); err != nil {
			return nil, err
		}
		_, _ = buf.WriteString("\n")
	}

	return bytes.TrimSpace(buf.Bytes()), nil
}
