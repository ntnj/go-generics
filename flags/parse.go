package flags

import (
	"fmt"
	"reflect"
	"strconv"
)

func parseValue(s string, v any) error {
	r := reflect.ValueOf(v).Elem()
	switch r.Kind() {
	case reflect.String:
		r.SetString(s)
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		r.SetBool(b)
	case reflect.Int:
		i, err := strconv.ParseInt(s, 0, strconv.IntSize)
		if err != nil {
			return err
		}
		r.SetInt(i)
	case reflect.Int64:
		i, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			return err
		}
		r.SetInt(i)
	case reflect.Uint:
		i, err := strconv.ParseUint(s, 0, strconv.IntSize)
		if err != nil {
			return err
		}
		r.SetUint(i)
	case reflect.Uint64:
		i, err := strconv.ParseUint(s, 0, 64)
		if err != nil {
			return err
		}
		r.SetUint(i)
	case reflect.Float64:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		r.SetFloat(f)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil
}
