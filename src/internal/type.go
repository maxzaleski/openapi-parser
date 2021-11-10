package internal

import "strings"

// IsStandardType checks whether the given key is a standard type i.e., not a request or response.
func IsStandardType(key string) bool {
	return !strings.Contains(key, "Request") && !strings.Contains(key, "Response")
}
