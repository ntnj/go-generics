package matchers

import "testing"

func Expect[T any](t testing.TB, v T, m Matcher[T]) {
	t.Helper()
	if err := m.Match(v); err != nil {
		if es := err.Error(); es != "" {
			t.Errorf("\n got: %v\nwant: %v\nwith: %v", v, m, es)
		} else {
			t.Errorf("\n got: %v\nwant: %v", v, m)
		}
	}
}

func Assert[T any](t testing.TB, v T, m Matcher[T]) {
	t.Helper()
	if err := m.Match(v); err != nil {
		if es := err.Error(); es != "" {
			t.Fatalf("\n got: %v\nwant: %v\nwith: %v", v, m, es)
		} else {
			t.Fatalf("\n got: %v\nwant: %v", v, m)
		}
	}
}
