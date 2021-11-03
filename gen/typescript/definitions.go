package typescript

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"openapi-gen/gen/parser"
	"openapi-gen/internal/utils"
)

// GenerateFromDefinitions generates typescript types from the given definitions.
func GenerateFromDefinitions(defs map[string]*parser.Definition) string {
	mappedDefs := make([]string, 0, len(defs))
	for _, k := range utils.MapIntoSortedKeys(defs) {
		def := defs[k]

		resultDef := ""
		if def.Type == "enum" {
			resultDef = generateEnum(def)
		} else {
			resultDef = generateInterface(def)
		}
		mappedDefs = append(mappedDefs, resultDef)
	}
	return strings.Join(mappedDefs, "\n\n")
}

func generateInterface(def *parser.Definition) string {
	result := ""
	if defDesc := def.Description; defDesc != "" {
		result += "// " + defDesc + "\n"
	}
	result += "interface %s {\n%s\n}"

	mappedProps := make([]string, 0, len(def.Properties))
	for _, prop := range def.Properties {
		mappedProps = append(mappedProps, generateProperty(prop))
	}
	return fmt.Sprintf(result, def.Key, strings.Join(mappedProps, "\n"))
}

func generateEnum(def *parser.Definition) string {
	result := "enum %s {\n%s\n}"

	mappedEntries := make([]string, 0, len(def.Properties))
	for _, entry := range def.EnumEntries {
		if entry == "" {
			continue
		}
		mappedEntries = append(mappedEntries,
			fmt.Sprintf("\t%s = '%s',", strcase.ToScreamingSnake(entry), entry))
	}
	return fmt.Sprintf(result, def.Key, strings.Join(mappedEntries, "\n"))
}

func generateProperty(prop *parser.DefinitionProperty) string {
	result := ""

	propDesc := prop.Description
	if propDesc == "" {
		// When a property is referenced as another, swagger-go will omit the comment.
		switch prop.Key {
		case "whereabouts":
			propDesc = "The member's last signed-in location."
		}
	}
	if propDesc != "" {
		result += "\t// " + propDesc + "\n"
	}

	requiredFlag := ""
	if propRequired := prop.Required; !propRequired {
		requiredFlag = "?"
	}
	propType := prop.Type
	if propRef := prop.Ref; propRef != "" {
		propType = propRef
	}
	switch prop.Type {
	case "integer":
		propType = "number"
	case "array":
		propType += "[]"
	}

	result += fmt.Sprintf("\t%s%s: %s;", prop.Key, requiredFlag, propType)
	return result
}
