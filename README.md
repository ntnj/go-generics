# Go Generics

## Matchers
[Matchers](matchers/README.md) provides a collection of utilities to write [googletest](https://github.com/google/googletest) inspired tests in golang.

```go
import m "github.com/ntnj/go-generics/matchers"

func TestM(t *testing.T) {
  got := 42
  m.Expect(t, got, m.Eq(42))

  list := []int{2,4,6}
  m.Expect(t, list, m.Contains(m.Eq(4)))
  m.Expect(t, list, m.IsPermutationOf(m.Eq(4), m.Eq(6), m.Eq(2)))
}
```