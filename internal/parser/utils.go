package parser

import "strings"

// toRef strips the definition prefix from a type reference.
func toRef(s string) string {
	return strings.Replace(s, "#/definitions/", "", 1)
}
