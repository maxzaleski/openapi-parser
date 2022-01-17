package typescript

import (
	"fmt"
	"openapi-generator/internal"
	"strings"

	"openapi-generator/internal/parser"

	"openapi-generator/gen/typescript/templates"
)

// isInterface checks whether the given key should be an interface.
func isInterface(key string) bool {
	if strings.Contains(key, "Request") || strings.Contains(key, "DynamicQuery") {
		return true
	}
	return false
}

// generateInterface generates a typescript interface from the given definition.
func generateInterface(def *parser.Definition) string {
	template := templates.Interface
	if strings.Contains(def.Key, "RequestBody") {
		template = templates.RequestBody
	}

	// Interface's description.
	if defDesc := def.Description; defDesc != "" {
		template = toJSDoc("", defDesc) + template
	}
	// Interface's properties.
	mappedProps := make([]string, 0, len(def.Properties))
	for _, prop := range def.Properties {
		mappedProps = append(mappedProps, generateObjectProperty(def.Key, "m.", prop))
	}

	return fmt.Sprintf(template, def.Key, strings.Join(mappedProps, "\n"))
}

// generateDynamicQueryFilters generates a typescript interface from the given definition
// (assumes model to be `{Prefix}DynamicQueryFilters`).
func generateDynamicQueryFilters(defs map[string]*parser.Definition, def *parser.Definition) string {
	for _, prop := range def.Properties {
		filter := defs[prop.Ref]
		filterProp := filter.Properties[1]

		valueType := filterProp.Ref
		if valueType == "" {
			valueType = filterProp.Type
		}
		// Account for enum values.
		if !internal.HasConstructor(valueType) {
			valueType = "e." + valueType
		}
		prop.Ref = "DynamicQueryFilter<" + valueType + ">"
	}
	return generateInterface(def)
}
