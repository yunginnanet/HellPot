package toml

import (
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"strings"
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
	return -1, fmt.Errorf("%w: missing struct tag for '%s'", ErrMismatchingSchema, sought)

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
		if ai > math.MaxInt64 {
			panic("uint too large to convert to int64")
		}
		aii = int64(ai)
	case uint8:
		aii = int64(ai)
	case uint16:
		aii = int64(ai)
	case uint32:
		aii = int64(ai)
	case uint64:
		if ai > math.MaxInt64 {
			panic("uint64 too large to convert to int64")
		}
		aii = int64(ai)
	default:
		return -1, false
	}
	return aii, true
}

func toInts(tv any) ([]int, bool) {
	if _, ok := tv.([]int); ok {
		return tv.([]int), true
	}
	if _, ok := tv.([]string); !ok {
		return nil, false
	}
	// convert []string to []int if all elements are int-ish
	var isIntSlice = true
	var intSlice []int
	for _, ss := range tv.([]string) {
		n, err := strconv.Atoi(strings.TrimSpace(ss))
		if err != nil {
			isIntSlice = false
			break
		}
		if intSlice == nil {
			intSlice = make([]int, 0, len(tv.([]string)))
		}
		// println("n: ", n)
		intSlice = append(intSlice, n)
	}
	if !isIntSlice {
		return nil, false
	}

	return intSlice, true
}

func (d *decoder) applySlice(targetElem reflect.Value, field reflect.StructField, tk string, tv any) error {
	switch field.Type.Elem().Kind() { //nolint:exhaustive
	case reflect.String:
		if _, strOK := tv.([]string); !strOK {
			return fmt.Errorf("expected []string for %s, got %T", tk, tv)
		}
	case reflect.Int:
		var intSlice []int
		var intOK bool
		if intSlice, intOK = toInts(tv); !intOK {
			return fmt.Errorf("expected []int for %s, got %T", tk, tv)
		}
		targetElem.Set(reflect.ValueOf(intSlice))
		return nil
	case reflect.Int64:
		if _, intOK := tv.([]int64); !intOK {
			return fmt.Errorf("expected []int64 for %s, got %T", tk, tv)
		}
	case reflect.Int32:
		if _, intOK := tv.([]int32); !intOK {
			return fmt.Errorf("expected []int32 for %s, got %T", tk, tv)
		}
	case reflect.Uint64:
		if _, intOK := tv.([]uint64); !intOK {
			return fmt.Errorf("expected []uint64 for %s, got %T", tk, tv)
		}
	case reflect.Float64:
		if _, floatOK := tv.([]float64); !floatOK {
			return fmt.Errorf("expected []float64 for %s, got %T", tk, tv)
		}
	case reflect.Float32:
		if _, floatOK := tv.([]float32); !floatOK {
			return fmt.Errorf("expected []float32 for %s, got %T", tk, tv)
		}
	case reflect.Bool:
		if _, boolOK := tv.([]bool); !boolOK {
			return fmt.Errorf("expected []bool for %s, got %T", tk, tv)
		}
	default:
		return fmt.Errorf("unsupported type %v for %s", field.Type.Elem().Kind(), tk)
	}

	targetElem.Set(reflect.ValueOf(tv))

	return nil
}

func (d *decoder) applyValue(targetElem reflect.Value, tk string, tv any) error {
	for i := 0; i < targetElem.NumField(); i++ {

		field := targetElem.Type().Field(i)

		if field.Type.Kind() == reflect.Ptr {
			if targetElem.Field(i).IsNil() {
				targetElem.Field(i).Set(reflect.New(field.Type.Elem()))
			}
			if targetElem.Field(i).Elem().Kind() != reflect.Struct {
				continue
			}
			targetElem = targetElem.Field(i).Elem()
		}

		if !field.IsExported() {
			continue
		}
		if field.Tag.Get("toml") != tk {
			continue
		}

		// println("found : ", field.Type.Kind().String(), " in struct for toml tag: ", field.Tag.Get("toml"))
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
			if err := d.applySlice(targetElem.Field(i), field, tk, tv); err != nil {
				return err
			}
		default:
			return fmt.Errorf("%w: unsupported type %s", ErrMismatchingSchema, field.Type.Kind().String())
		}
	}

	return nil
}

func (d *decoder) handleSubTableKey(targetElem reflect.Value, k string) (*reflect.Value, string, error) {
	if len(strings.Split(k, ".")) < 2 {
		return nil, "", fmt.Errorf("%w: expected at least two parts in key %s", ErrMismatchingSchema, k)
	}
	subStructField, subStructErr := findStructTag(targetElem, strings.Split(k, ".")[0])
	if subStructErr != nil {
		return nil, "", subStructErr
	}
	targetSubStructElem := reflect.ValueOf(d.target).Elem().Field(subStructField)
	if targetSubStructElem.Kind() == reflect.Ptr {
		if targetSubStructElem.IsNil() {
			return nil, "", fmt.Errorf("sub-struct for '%s' is nil", k)
		}
		targetSubStructElem = targetSubStructElem.Elem()
	}
	if targetSubStructElem.Kind() != reflect.Struct {
		return nil, "", fmt.Errorf("%w: expected struct for '%s', got %T", ErrMismatchingSchema, k, targetSubStructElem)
	}
	targetElem = targetSubStructElem
	k = strings.Split(k, ".")[1]

	return &targetElem, k, nil
}
func (d *decoder) handleTable(table map[string]interface{}, k string) error {
	targetElem := reflect.ValueOf(d.target).Elem()

	if strings.Contains(k, ".") {
		var err error
		var targetElemPtr *reflect.Value
		if targetElemPtr, k, err = d.handleSubTableKey(targetElem, k); err != nil {
			return err
		}
		targetElem = *targetElemPtr
	}

	fieldIndex, fieldErr := findStructTag(targetElem, k)
	if fieldErr != nil {
		return fieldErr
	}
	//  println("field index: ", fieldIndex)

	switch targetElem.Type().Field(fieldIndex).Type.Kind() { //nolint:exhaustive
	case reflect.Ptr:
		if targetElem.Field(fieldIndex).IsNil() {
			return fmt.Errorf("%w: expected struct for '%s', got nil ptr", ErrMismatchingSchema, k)
		}
		if targetElem.Field(fieldIndex).Elem().Kind() != reflect.Struct {
			return fmt.Errorf(
				"%w: expected struct for '%s', got %s",
				ErrMismatchingSchema, k, targetElem.Field(fieldIndex).Elem().Type().Kind().String(),
			)
		}
		targetElem = targetElem.Field(fieldIndex).Elem()
	case reflect.Struct:
		targetElem = targetElem.Field(fieldIndex)
	default:
		return fmt.Errorf(
			"%w: expected struct for '%s', got %s",
			ErrMismatchingSchema, k, targetElem.Field(fieldIndex).Type().Kind().String(),
		)
	}

	for tk, tv := range table {
		if err := d.applyValue(targetElem, tk, tv); err != nil {
			return err
		}
	}

	return nil
}

func (d *decoder) handleTables() error {
	// spew.Dump(d.tables)

	for k, v := range d.tables {
		vTable, tableOK := v.(map[string]interface{})
		if !tableOK {
			return fmt.Errorf("%w: expected table, got %T", ErrMismatchingSchema, v)
		}
		// println("table:", spew.Sdump(vTable))

		if err := d.handleTable(vTable, k); err != nil {
			return err
		}
	}

	return nil
}
