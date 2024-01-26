package gflag

import (
	"flag"
	"testing"

	m "github.com/ntnj/go-generics/matchers"
)

func parse[T any](s string) T {
	v, err := parseFlagValue[T](s)
	if err != nil {
		panic(err)
	}
	return v
}

func TestParse(t *testing.T) {
	m.Expect(t, parse[int32]("123"), m.Eq[int32](123))
	m.Expect(t, parse[int]("123"), m.Eq(123))
	m.Expect(t, parse[string]("123"), m.Eq("123"))
}

func TestFlags(t *testing.T) {
	fs := flag.NewFlagSet("Test", flag.ContinueOnError)
	i := InFlagSet(fs, "an_int", 42, "help an int")
	s := InFlagSet(fs, "a_string", "hello", "help a string")
	b := InFlagSet(fs, "a_bool", false, "help a bool")
	f := InFlagSet(fs, "a_float", 1.9, "help a float")
	if err := fs.Parse([]string{"-an_int", "32", "-a_string", "world", "-a_bool", "-a_float", "2.9"}); err != nil {
		t.Error(err)
	}
	m.Expect(t, i.Get(), m.Eq(32))
	m.Expect(t, s.Get(), m.Eq("world"))
	m.Expect(t, b.Get(), m.Eq(true))
	m.Expect(t, f.Get(), m.Eq(2.9))
}
