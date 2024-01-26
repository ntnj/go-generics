package gflag

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

type Flag[T any] struct {
	value T
	set   bool
}

func (f *Flag[T]) Get() T {
	return f.value
}

func (f *Flag[T]) Set(v T) {
	f.value = v
	f.set = true
}

func New[T any](name string, def T, usage string) *Flag[T] {
	return InFlagSet[T](flag.CommandLine, name, def, usage)
}

func InFlagSet[T any](fs *flag.FlagSet, name string, def T, usage string) *Flag[T] {
	f := &Flag[T]{value: def}
	parse := func(s string) error {
		v, err := parseFlagValue[T](s)
		if err != nil {
			return err
		}
		f.Set(v)
		return nil
	}
	switch any(def).(type) {
	case bool:
		fs.BoolFunc(name, usage, parse)
	default:
		fs.Func(name, usage, parse)
	}
	return f
}

func parseFlagValue[T any](s string) (T, error) {
	var t T
	switch v := any(t).(type) {
	case flag.Value:
		err := v.Set(s)
		return t, err
	case string:
		reflect.ValueOf(&t).Elem().SetString(s)
		return t, nil
	case bool:
		vv, err := strconv.ParseBool(s)
		reflect.ValueOf(&t).Elem().SetBool(vv)
		return t, err
	case int:
		vv, err := strconv.Atoi(s)
		reflect.ValueOf(&t).Elem().SetInt(int64(vv))
		return t, err
	case int32:
		vv, err := strconv.ParseInt(s, 10, 32)
		reflect.ValueOf(&t).Elem().SetInt(vv)
		return t, err
	case int64:
		vv, err := strconv.ParseInt(s, 10, 64)
		reflect.ValueOf(&t).Elem().SetInt(vv)
		return t, err
	case float64:
		vv, err := strconv.ParseFloat(s, 64)
		reflect.ValueOf(&t).Elem().SetFloat(vv)
		return t, err
	}
	return t, fmt.Errorf("unsupported type: %T", t)
}
