package bob // package name must match the package name in bob_test.go

import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

func Hey(s string) string {
	s = strings.TrimSpace(s)

	switch {
	case s == "":
		return "Fine. Be that way!"
	case any(s, unicode.IsUpper) && !any(s, unicode.IsLower):
		return "Whoa, chill out!"
	case s[len(s)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}
}

func any(s string, test func(rune) bool) bool {
	for _, c := range s {
		if test(c) {
			return true
		}
	}
	return false
}
