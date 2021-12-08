package parser

import (
	"regexp"
	"strings"
)

// toRef strips the definition prefix from a type reference.
func toRef(s string) string {
	return strings.Replace(s, "#/definitions/", "", 1)
}

var descriptionRegex = regexp.MustCompile(`^(.*\.)`)

// extractDescription extracts the description from a string.
func extractDescription(s string) string {
	matches := descriptionRegex.FindStringSubmatch(s)
	if len(matches) >= 2 {
		return matches[1]
	}
	return ""
}
