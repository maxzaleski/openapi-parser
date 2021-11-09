package typescript

import (
	"fmt"
	"strings"

	"openapi-gen/internal/parser"

	"openapi-gen/gen/typescript/templates"
)

// isInterface checks whether the given key should be an interface.
func isInterface(key string) bool {
	if strings.Contains(key, "Request") || strings.Contains(key, "ListMembers") {
		return true
	}
	return false
}

// generateInterface generates a typescript interface from the given definition.
func generateInterface(def *parser.Definition) string {
	template := templates.Interface

	// Interface's description.
	if defDesc := def.Description; defDesc != "" {
		template = toJSDoc("", defDesc) + template
	}
	// Interface's properties.
	mappedProps := make([]string, 0, len(def.Properties))
	for _, prop := range def.Properties {
		mappedProps = append(mappedProps, generateObjectProperty("", prop))
	}

	return fmt.Sprintf(template, def.Key, strings.Join(mappedProps, "\n"))
}
