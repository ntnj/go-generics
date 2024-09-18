package options

import (
	"flag"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

type Option[T any] struct {
	key   string
	name  string
	desc  string
	value T
	set   bool
}

func (o *Option[T]) Get() T {
	return o.value
}

type option struct {
}

type Registry map[string]option

var defaultRegistry = make(Registry)

func New[T any](name string, def T, desc string) *Option[T] {
	return new(defaultRegistry, name, def, desc)
}

func new[T any](reg Registry, name string, def T, desc string) *Option[T] {
	o := &Option[T]{
		key:   name,
		name:  name,
		desc:  desc,
		value: def,
	}
	if _, ok := reg[name]; ok {
		log.Fatalf("duplicate option: %s", name)
	}
	reg[name] = option{}
	return o
}

func ParseCommandLine() {}

func parseValue[T any](s string) (T, error) {
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
