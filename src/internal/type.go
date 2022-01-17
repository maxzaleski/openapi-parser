package internal

import (
	"regexp"
)

var standardTypeRegex = regexp.MustCompile("(Request|Response|List)")

// IsStandardModel checks whether the given key is a standard type i.e., not a request or response.
func IsStandardModel(key string) bool { return !standardTypeRegex.MatchString(key) }
