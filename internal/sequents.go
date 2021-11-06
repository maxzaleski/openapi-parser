package internal

import (
	"strings"

	"openapi-gen/gen/parser"
)

// HasConstructor checks whether the given key has a constructor.
func HasConstructor(ref string) bool {
	switch ref {
	case
		"APIError",
		"Colour",
		"EntityType",
		"ViewEntityType",
		"ErrorCode",
		"ErrorType",
		"RelationshipWithMember",
		"MemberRole":
		return false
	default:
		return true
	}
}

// IsErrorType checks whether the given key is an error type.
func IsErrorType(key string) bool {
	return strings.Contains(strings.ToLower(key), "error")
}

// IsPathParameter checks whether the givens properties correspond to a path parameter.
func IsPathParameter(props []*parser.DefinitionProperty) bool {
	return len(props) == 1 && props[0].In != "path"
}

// IsSuitedForAPIMethod checks whether the given properties are suited for an API method.
func IsSuitedForAPIMethod(props []*parser.DefinitionProperty) bool {
	return len(props) > 1 || IsPathParameter(props)
}

// IsPropSuitableForValidation checks whether the given property is suitable for validation.
func IsPropSuitableForValidation(t string) bool {
	switch t {
	case "string", "number", "array":
		return true
	default:
		return false
	}
}
