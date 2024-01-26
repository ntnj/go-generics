package gflag

import (
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
	test()
}
