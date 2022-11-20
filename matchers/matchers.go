package matchers

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Matcher[T any] interface {
	fmt.Stringer
	Match(v T) error
}

type simpleMatcher[T any] struct {
	fn   func(x T) bool
	desc string
	args []any
}

func (m *simpleMatcher[T]) Match(v T) error {
	if !m.fn(v) {
		return errors.New("")
	}
	return nil
}

func (m *simpleMatcher[T]) String() string {
	return fmt.Sprintf(m.desc, m.args...)
}

func createSimple[T any](fn func(x T) bool, desc string, args ...any) Matcher[T] {
	if false {
		// To allow go vet to validate pattern.
		fmt.Printf(desc, args...)
	}
	return &simpleMatcher[T]{fn, desc, args}
}

func Eq[T comparable](v T) Matcher[T] {
	return createSimple(func(x T) bool { return x == v }, "is equal to %v", v)
}

func Ge[T constraints.Ordered](v T) Matcher[T] {
	return createSimple(func(x T) bool { return x >= v }, "is greater or equal to %v", v)
}

func Gt[T constraints.Ordered](v T) Matcher[T] {
	return createSimple(func(x T) bool { return x > v }, "is greater than %v", v)
}

func Le[T constraints.Ordered](v T) Matcher[T] {
	return createSimple(func(x T) bool { return x <= v }, "is less or equal to %v", v)
}

func Lt[T constraints.Ordered](v T) Matcher[T] {
	return createSimple(func(x T) bool { return x < v }, "is less than %v", v)
}

// func IsNil[T any]() Matcher[T] {
//  // TODO: use reflect
// 	return createSimple(func(x T) bool { return x == nil }, "is nil")
// }

type notMatcher[T any] struct {
	im Matcher[T]
}

func (m *notMatcher[T]) Match(v T) error {
	if err := m.im.Match(v); err == nil {
		return errors.New("")
	}
	return nil
}

func (m *notMatcher[T]) String() string {
	return fmt.Sprintf("not %v", m.im.String())
}

func Not[T any](m Matcher[T]) Matcher[T] {
	return &notMatcher[T]{m}
}

type allOfMatcher[T any] struct {
	ims []Matcher[T]
}

func (m *allOfMatcher[T]) Match(v T) error {
	for _, im := range m.ims {
		if err := im.Match(v); err != nil {
			return err
		}
	}
	return nil
}

func (m *allOfMatcher[T]) String() string {
	var ss []string
	for _, im := range m.ims {
		ss = append(ss, im.String())
	}
	return strings.Join(ss, ", and ")
}

func AllOf[T any](ms ...Matcher[T]) Matcher[T] {
	return &allOfMatcher[T]{ms}
}

type anyOfMatcher[T any] struct {
	ims []Matcher[T]
}

func (m *anyOfMatcher[T]) Match(v T) error {
	for _, im := range m.ims {
		if err := im.Match(v); err == nil {
			return nil
		}
	}
	return errors.New("")
}

func (m *anyOfMatcher[T]) String() string {
	var ss []string
	for _, im := range m.ims {
		ss = append(ss, im.String())
	}
	return strings.Join(ss, ", or ")
}

func AnyOf[T any](ms ...Matcher[T]) Matcher[T] {
	return &anyOfMatcher[T]{ms}
}
