package flags

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Slice[T ~bool | ~int | ~int64 | ~uint | ~uint64 | ~float64 | ~string] []T

func (s *Slice[T]) Get() []T {
	return *s
}

func (s *Slice[T]) String() string {
	ret := make([]string, len(*s))
	for i, v := range *s {
		ret[i] = fmt.Sprint(v)
	}
	return strings.Join(ret, ",")
}

func (s *Slice[T]) Set(value string) error {
	for _, val := range strings.Split(value, ",") {
		var t T
		r := reflect.ValueOf(&t).Elem()
		switch r.Kind() {
		case reflect.String:
			r.SetString(val)
		case reflect.Bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			r.SetBool(b)
		case reflect.Int:
			i, err := strconv.ParseInt(val, 0, strconv.IntSize)
			if err != nil {
				return err
			}
			r.SetInt(i)
		case reflect.Int64:
			i, err := strconv.ParseInt(val, 0, 64)
			if err != nil {
				return err
			}
			r.SetInt(i)
		case reflect.Uint:
			i, err := strconv.ParseUint(val, 0, strconv.IntSize)
			if err != nil {
				return err
			}
			r.SetUint(i)
		case reflect.Uint64:
			i, err := strconv.ParseUint(val, 0, 64)
			if err != nil {
				return err
			}
			r.SetUint(i)
		case reflect.Float64:
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			r.SetFloat(f)
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		*s = append(*s, t)
	}
	return nil
}

type SliceVar[T flag.Value] []T

func (s *SliceVar[T]) String() string {
	ret := make([]string, len(*s))
	for i, v := range *s {
		ret[i] = fmt.Sprint(v)
	}
	return strings.Join(ret, ",")
}

func (s *SliceVar[T]) Set(value string) error {
	for _, val := range strings.Split(value, ",") {
		var v T
		if err := v.Set(val); err != nil {
			return err
		}
		*s = append(*s, v)
	}
	return nil
}
