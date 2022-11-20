package matchers

import (
	"regexp"
	"strings"
)

func ContainsRegex(r string) Matcher[string] {
	re := regexp.MustCompile(r)
	return createSimple(func(x string) bool { return re.MatchString(x) }, "contains regex %v", r)
}

func ContainsRegexPOSIX(r string) Matcher[string] {
	re := regexp.MustCompilePOSIX(r)
	return createSimple(func(x string) bool { return re.MatchString(x) }, "contains regex %v", r)
}

func EndsWith(s string) Matcher[string] {
	return createSimple(func(x string) bool { return strings.HasSuffix(x, s) }, "ends with %v", s)
}

func StartsWith(s string) Matcher[string] {
	return createSimple(func(x string) bool { return strings.HasPrefix(x, s) }, "starts with %v", s)
}

func HasSubstr(s string) Matcher[string] {
	return createSimple(func(x string) bool { return strings.Contains(x, s) }, "has substring %v", s)
}
