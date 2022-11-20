package matchers

import (
	"errors"
	"fmt"
)

type containsMatcher[T any] struct {
	im Matcher[T]
}

func (m *containsMatcher[T]) Match(v []T) error {
	for _, x := range v {
		if err := m.im.Match(x); err == nil {
			return nil
		}
	}
	return errors.New("")
}

func (m *containsMatcher[T]) String() string {
	return fmt.Sprintf("contains an element which %v", m.im)
}

// Contains checks if the list contains an element matching m.
func Contains[T any](m Matcher[T]) Matcher[[]T] { return &containsMatcher[T]{m} }

// ContainsEq checks if the list argument contains an element equal to v.
func ContainsEq[T comparable](v T) Matcher[[]T] { return Contains(Eq(v)) }

type eachMatcher[T any] struct {
	im Matcher[T]
}

func (m *eachMatcher[T]) Match(v []T) error {
	for _, e := range v {
		if err := m.im.Match(e); err != nil {
			return err
		}
	}
	return nil
}

func (m *eachMatcher[T]) String() string {
	return fmt.Sprintf("only contains elements that %v", m.im)
}

// Each checks if each element in list matches m.
func Each[T any](m Matcher[T]) Matcher[[]T] { return &eachMatcher[T]{m} }

func eqs[T comparable](vs ...T) []Matcher[T] {
	m := make([]Matcher[T], len(vs))
	for i, v := range vs {
		m[i] = Eq(v)
	}
	return m
}

type elementsAreMatcher[T any] struct {
	ims []Matcher[T]
}

func (m *elementsAreMatcher[T]) Match(v []T) error {
	if len(m.ims) != len(v) {
		return errors.New("size mismatch")
	}
	for i, x := range v {
		if err := m.ims[i].Match(x); err != nil {
			return err
		}
	}
	return nil
}

func (m *elementsAreMatcher[T]) String() string {
	if len(m.ims) == 0 {
		return "is empty"
	} else if len(m.ims) == 1 {
		return fmt.Sprintf("has one element which %v", m.ims[0])
	} else {
		s := ""
		for i, im := range m.ims {
			s += fmt.Sprintf("\n  has element #%d which %v", i, im)
		}
		return s
	}
}

// ElementsAre checks each element in list matches in order.
func ElementsAre[T any](m ...Matcher[T]) Matcher[[]T]  { return &elementsAreMatcher[T]{m} }
func ElementsAreEq[T comparable](vs ...T) Matcher[[]T] { return ElementsAre(eqs(vs...)...) }

type unorderedMatcher[T any] struct {
	matchSuperset, matchSubset bool
	ims                        []Matcher[T]
}

func (m *unorderedMatcher[T]) Match(v []T) error {
	if m.matchSubset && m.matchSuperset && len(v) != len(m.ims) {
		return errors.New("size mismatch")
	}
	index := func(ei, mi int) int { return ei*len(m.ims) + mi }
	matches := make([]bool, len(v)*len(m.ims))
	em := make([]bool, len(v))
	mm := make([]bool, len(m.ims))
	for ei, e := range v {
		for mi, im := range m.ims {
			if err := im.Match(e); err == nil {
				matches[index(ei, mi)] = true
				em[ei] = true
				mm[mi] = true
			}
		}
	}
	if m.matchSuperset {
		for mi, im := range m.ims {
			if !mm[mi] {
				return fmt.Errorf("where no element %v", im)
			}
		}
	}
	if m.matchSubset {
		for ei, e := range v {
			if !em[ei] {
				return fmt.Errorf("element %v didn't match", e)
			}
		}
	}

	// Bipartite matching algorithm copied from googletest:gmock-matchers.cc
	ee := make([]int, len(v))
	me := make([]int, len(m.ims))
	for ei := range ee {
		ee[ei] = -1
	}
	for mi := range me {
		me[mi] = -1
	}
	for ei := range v {
		seen := make([]bool, len(m.ims))
		var bpm func(int) bool
		bpm = func(lhs int) bool {
			for mi := range m.ims {
				if seen[mi] || !matches[index(lhs, mi)] {
					continue
				}
				seen[mi] = true
				if me[mi] == -1 || bpm(me[mi]) {
					ee[lhs] = mi
					me[mi] = lhs
					return true
				}
			}
			return false
		}
		bpm(ei)
	}
	mf := 0
	for _, mi := range ee {
		if mi != -1 {
			mf += 1
		}
	}

	if (m.matchSuperset && mf < len(m.ims)) || m.matchSubset && mf < len(v) {
		s := ""
		for ei, mi := range ee {
			if mi != -1 {
				s += fmt.Sprintf("\n  element #%d (%v) matched #%d: %v", ei, v[ei], mi, m.ims[mi])
			}
		}
		return fmt.Errorf("closest match is%v", s)
	}

	return nil
}

func (m *unorderedMatcher[T]) String() string {
	s := ""
	if m.matchSubset && m.matchSuperset {
		s = fmt.Sprintf("has %d elements, which are permutation of:", len(m.ims))
	} else if m.matchSuperset {
		s = "is superset of"
	} else if m.matchSubset {
		s = "is subset of"
	}
	for _, im := range m.ims {
		s += fmt.Sprintf("\n  has element which %v", im)
	}
	return s
}

func UnorderedElementsAre[T any](m ...Matcher[T]) Matcher[[]T] {
	return &unorderedMatcher[T]{true, true, m}
}
func UnorderedElementsAreEq[T comparable](vs ...T) Matcher[[]T] {
	return UnorderedElementsAre(eqs(vs...)...)
}
func IsPermutationOf[T any](m ...Matcher[T]) Matcher[[]T] {
	return UnorderedElementsAre(m...)
}
func IsPermutationOfEq[T comparable](vs ...T) Matcher[[]T] {
	return IsPermutationOf(eqs(vs...)...)
}

func IsSupersetOf[T any](m ...Matcher[T]) Matcher[[]T]  { return &unorderedMatcher[T]{true, false, m} }
func IsSupersetOfEq[T comparable](vs ...T) Matcher[[]T] { return IsSupersetOf(eqs(vs...)...) }

func IsSubsetOf[T any](m ...Matcher[T]) Matcher[[]T]  { return &unorderedMatcher[T]{false, true, m} }
func IsSubsetOfEq[T comparable](vs ...T) Matcher[[]T] { return IsSubsetOf(eqs(vs...)...) }
