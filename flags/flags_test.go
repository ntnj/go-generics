package flags

import (
	"flag"
	"testing"

	m "github.com/ntnj/go-generics/matchers"
)

type id int64

func TestSliceFlags(t *testing.T) {
	var is Slice[int]
	var ss Slice[string]
	var ids Slice[id]
	var fs flag.FlagSet
	fs.Var(&is, "is", "")
	fs.Var(&ss, "ss", "")
	fs.Var(&ids, "ids", "")
	if err := fs.Parse([]string{"-is=3,43,-2", "-ids=1,2,3,4", "-ss=a,b,c,2"}); err != nil {
		t.Fatal(err)
	}
	m.Expect(t, is.Get(), m.ElementsAreEq(3, 43, -2))
	m.Expect(t, ids.Get(), m.ElementsAreEq[id](1, 2, 3, 4))
	m.Expect(t, ss.Get(), m.ElementsAreEq("a", "b", "c", "2"))
}
