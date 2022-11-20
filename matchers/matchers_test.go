package matchers

import "testing"

func pass[T any](t *testing.T, v T, m Matcher[T]) {
	t.Helper()
	err := m.Match(v)
	t.Logf("matcher=\n%v", m)
	if err != nil {
		t.Errorf("fail: %v", err)
	}
}

func fail[T any](t *testing.T, v T, m Matcher[T]) {
	t.Helper()
	err := m.Match(v)
	t.Logf("matcher=\n%v", m)
	if err == nil {
		t.Errorf("expected error: %v", err)
	}
}

func TestSimple(t *testing.T) {
	pass(t, 2, Eq(2))
	pass(t, "abc", Eq("abc"))

	fail(t, 3, Eq(2))
	fail(t, "cba", Eq("abc"))

	pass(t, 4, Lt(5))
	fail(t, 5, Lt(4))

	pass(t, 5, Gt(4))
	fail(t, 4, Gt(5))

	pass(t, 5, AllOf(Eq(5), Lt(6), Gt(4)))
	fail(t, 5, AllOf(Eq(5), Gt(6)))

	pass(t, 5, AnyOf(Eq(5), Gt(6)))
	fail(t, 5, AnyOf(Eq(4), Gt(6)))
}

func TestString(t *testing.T) {
	pass(t, "abc", EndsWith("bc"))
	pass(t, "abc", StartsWith("ab"))
	pass(t, "abcd", HasSubstr("bc"))
	pass(t, "abcbc", ContainsRegex("(bc){2}"))
}

func TestList(t *testing.T) {
	pass(t, []int{2, 4}, Contains(Eq(2)))
	fail(t, []int{2, 4}, Contains(Eq(3)))

	pass(t, []int{2, 4, 6}, Each(Lt(8)))
	fail(t, []int{2, 4, 6}, Each(Lt(5)))

	pass(t, []int{2, 4, 6}, ElementsAre(Eq(2), Eq(4), Eq(6)))
	fail(t, []int{2, 4, 6}, ElementsAre(Eq(2), Eq(6), Eq(4)))

	pass(t, []int{2, 4, 6}, UnorderedElementsAre(Eq(2), Eq(4), Eq(6)))
	pass(t, []int{2, 4, 6}, UnorderedElementsAre(Eq(2), Eq(6), Eq(4)))

	pass(t, []int{2, 4, 6}, IsSupersetOf(Eq(2), Eq(6)))
	fail(t, []int{2, 4, 6}, IsSupersetOf(Eq(2), Eq(5)))

	pass(t, []int{2, 4}, IsSubsetOf(Eq(2), Eq(6), Eq(4)))
}
