package toml

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

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
	switch field := targetElem; field.Type().Elem().Kind() { //nolint:exhaustive
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

func findStructTag(struc reflect.Value, sought string) (int, error) {
	for i := 0; i < struc.NumField(); i++ {
		if struc.Type().Field(i).Tag.Get("toml") == sought {
			return i, nil
		}
	}
	return -1, fmt.Errorf("%w: %s", ErrMismatchingSchema, sought)

}

func isIntIsh(a any) (int64, bool) {
	var aii int64
	switch ai := a.(type) {
	case int64:
		aii = ai
	case int32:
		aii = int64(ai)
	case int16:
		aii = int64(ai)
	case int8:
		aii = int64(ai)
	case uint:
		aii = int64(ai)
	case uint8:
		aii = int64(ai)
	case uint16:
		aii = int64(ai)
	case uint32:
		aii = int64(ai)
	case uint64:
		aii = int64(ai)
	default:
		return -1, false
	}
	return aii, true
}

func (d *decoder) applySlice(targetElem reflect.Value, field reflect.StructField, i int, tk string, tv any) error {
	switch field.Type.Elem().Kind() {
	case reflect.String:
		var strok bool
		var strSlice []string
		if strSlice, strok = tv.([]string); !strok {
			return fmt.Errorf("expected []string for %s, got %T", tk, tv)
		}
		targetElem.Field(i).Grow(len(strSlice))
		for si, sii := range strSlice {
			fmt.Printf("[key: %s]: setting %s at index %d\n", tk, sii, si)
			targetElem.Field(i).Index(si).SetString(sii)
		}
		return nil
	case reflect.Int:
		// var
	}

	return nil
}

func (d *decoder) applyValue(targetElem reflect.Value, tk string, tv any) error {
	for i := 0; i < targetElem.NumField(); i++ {
		field := targetElem.Type().Field(i)
		if !field.IsExported() {
			continue
		}
		if field.Tag.Get("toml") != tk {
			continue
		}
		println("found : ", field.Type.Kind().String(), " in struct for toml tag: ", field.Tag.Get("toml"))
		//nolint:exhaustive
		switch field.Type.Kind() {
		case reflect.String:
			if _, ok := tv.(string); !ok {
				return fmt.Errorf("%w: expected string, got %T", ErrMismatchingSchema, tv)
			}
			targetElem.Field(i).SetString(tv.(string)) //nolint:forcetypeassert
		case reflect.Int64, reflect.Int32, reflect.Int16,
			reflect.Int8, reflect.Uint64, reflect.Uint32, reflect.Int:
			var iok bool
			var i64 int64
			if i64, iok = isIntIsh(tv); !iok {
				return fmt.Errorf("%w: expected ~int, got %T", ErrMismatchingSchema, tv)
			}
			targetElem.Field(i).SetInt(i64)
		case reflect.Bool:
			var b bool
			var bok bool
			if b, bok = tv.(bool); !bok {
				return fmt.Errorf("%w: expected bool, got %T", ErrMismatchingSchema, tv)
			}
			targetElem.Field(i).SetBool(b)
		case reflect.Slice:
			if err := d.applySlice(targetElem.Field(i), field, i, tk, tv); err != nil {
				return err
			}
		default:
			return fmt.Errorf("%w: unsupported type %s", ErrMismatchingSchema, field.Type.Kind().String())
		}
	}

	return nil
}
func (d *decoder) handleTable(table map[string]interface{}, k string) error {
	targetElem := reflect.ValueOf(d.target).Elem()

	fieldIndex, fieldErr := findStructTag(targetElem, k)
	if fieldErr != nil {
		return fieldErr
	}
	println("field index: ", fieldIndex)
	if targetElem.Field(fieldIndex).Type().Kind() != reflect.Struct {
		return fmt.Errorf("%w: expected struct, got %T", ErrMismatchingSchema, targetElem.Field(fieldIndex).Type())
	}

	for tk, tv := range table {
		if err := d.applyValue(targetElem, tk, tv); err != nil {
			return err
		}

	}

	return nil
}

func (d *decoder) handleTables() error {
	spew.Dump(d.tables)

	for k, v := range d.tables {
		vTable, tableOK := v.(map[string]interface{})
		if !tableOK {
			return fmt.Errorf("%w: expected table, got %T", ErrMismatchingSchema, v)
		}
		println("table:", spew.Sdump(vTable))

		if err := d.handleTable(vTable, k); err != nil {
			return err
		}
	}

	return nil
}
