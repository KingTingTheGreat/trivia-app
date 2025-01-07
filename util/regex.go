package util

import (
	"regexp"
	"strings"
)

var invalidRanges = regexp.MustCompile("^[a-zA-Z0-9]+$")
var invalidChars = "_ -()*^@$/!:;,.<>{}+][]"

func HasInvalidChar(s string) bool {
	if invalidRanges.MatchString(s) {
		return false
	}

	return !strings.ContainsAny(s, invalidChars)
}
