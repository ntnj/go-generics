package gflag

import (
	"errors"
	"flag"
	"reflect"
	"strconv"
)

type IFlag interface {
	// Interface defined since golang doesn't allow generic methods.
	addTo(fs *FlagSet)
}

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

func (f *Flag[T]) AddTo(fs *FlagSet) *Flag[T] {
	f.addTo(fs)
	return f
}

func (f *Flag[T]) addTo(fs *FlagSet) {

}

func New[T any](name, short string, def T, usage string) *Flag[T] {
	return nil
}

func Struct[T FlagStruct]() *Flag[T] {
	return nil
}

type FlagSet struct {
	parsed bool
}

func (fs *FlagSet) Parse(args []string) error {
	if fs.parsed {
		return errors.New("already parsed")
	}
	return nil
}

func (fs *FlagSet) Parsed() bool {
	return fs.parsed
}

type FlagStruct interface {
	private()
}

type Config struct {
	FlagStruct
	FlagA int `flag:"name=flaga,short=a,default=23"`
}

func parseFlagValue[T any](s string) (T, error) {
	var t T
	switch v := any(t).(type) {
	case flag.Value:
		err := v.Set(s)
		return t, err
	case int:
		vv, err := strconv.Atoi(s)
		reflect.ValueOf(&t).Elem().SetInt(int64(vv))
		return t, err
	case int32:
		vv, err := strconv.ParseInt(s, 10, 32)
		reflect.ValueOf(&t).Elem().SetInt(int64(vv))
		return t, err
	}
	return t, nil
}

func testf(sf FlagStruct) {

}

func test() {
	sf := Config{}
	// sf := FlagSet{}
	testf(sf)
	// sf.private()
}
