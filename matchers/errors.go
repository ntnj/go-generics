package matchers

import "errors"

func WrapsError(err error) Matcher[error] {
	return createSimple(func(x error) bool { return errors.Is(x, err) }, "wraps %v", err)
}
