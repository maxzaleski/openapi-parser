package internal

import (
	"regexp"
)

var standardTypeRegex = regexp.MustCompile("(Request|Response|List)")

// IsStandardType checks whether the given key is a standard type i.e., not a request or response.
func IsStandardType(key string) bool { return !standardTypeRegex.MatchString(key) }
