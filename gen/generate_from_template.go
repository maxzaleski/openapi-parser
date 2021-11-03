package gen

import (
	"fmt"
	"log"
	"strings"

	"github.com/iancoleman/strcase"

	"openapi-gen/gen/parser"
)

const interfaceTemplate = `%s
interface %s {
%s
}
`

const propertyTemplate = `	%s
	%s: %s;`

const enumTemplate = `enum %s {
%s
}
`

func generateFromTemplate(doc *parser.Document, extn Extension, logger *log.Logger) string {
	result := "// Generated with love.\n\n"
	mappedDefinitions := make([]string, 0)
	for _, k := range mapIntoSortedKeys(doc.Definitions) {
		def := doc.Definitions[k]
		defDesc := def.Description
		if defDesc != "" {
			defDesc = "// " + defDesc
		}
		if def.Type == "enum" {
			entries := make([]string, 0)
			for _, entry := range def.EnumEntries {
				if entry == "" {
					continue
				}
				entries = append(
					entries, fmt.Sprintf("\t%s = '%s'", strcase.ToScreamingSnake(entry), entry))
			}
			mappedDefinitions = append(mappedDefinitions, fmt.Sprintf(enumTemplate, def.Key, strings.Join(entries, ",\n")))
		} else {
			properties := make([]string, 0)
			for _, prop := range def.Properties {
				// Type.
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
				// Desc.
				propDesc := prop.Description
				if propDesc != "" {
					propDesc = "// " + propDesc
				}
				properties = append(
					properties, fmt.Sprintf(propertyTemplate, propDesc, prop.Key, propType))
			}
			mappedDefinitions = append(mappedDefinitions, fmt.Sprintf(interfaceTemplate,
				defDesc,
				def.Key,
				strings.Join(properties, "\n"),
			))
		}
	}
	result += strings.Join(mappedDefinitions, "\n")
	return result
}
