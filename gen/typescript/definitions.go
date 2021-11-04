package typescript

import (
	"fmt"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"

	"openapi-gen/gen/parser"
	"openapi-gen/gen/typescript/constants"
	"openapi-gen/gen/typescript/templates"
	"openapi-gen/internal/utils"
)

// GenerateFromDefinitions generates typescript types from the given definitions.
func GenerateFromDefinitions(defs map[string]*parser.Definition) string {
	mappedDefs := make([]string, 0, len(defs))
	mappedDefs = append(mappedDefs,
		strings.TrimPrefix(constants.Countries, "\n"),
		strings.TrimPrefix(constants.ExtendedDate, "\n"),
	)
	for _, k := range utils.MapIntoSortedKeys(defs) {
		def := defs[k]

		resultDef := ""
		if def.Type == "enum" {
			resultDef = generateEnum(def)
		} else {
			if strings.Contains(def.Key, "Response") ||
				strings.Contains(def.Key, "Request") ||
				strings.Contains(def.Key, "Error") ||
				strings.Contains(def.Key, "ListMembers") {
				resultDef = generateInterface(def)
			} else {
				resultDef = generateClass(def)
			}
		}
		mappedDefs = append(mappedDefs, resultDef)
	}
	return strings.Join(mappedDefs, "\n\n")
}

// generateInterface generates a typescript interface from the given definition.
func generateInterface(def *parser.Definition) string {
	result := ""
	if defDesc := def.Description; defDesc != "" {
		result += "// " + defDesc + "\n"
	}
	result += "interface %s {\n%s\n}"

	mappedProps := make([]string, 0, len(def.Properties))
	for _, prop := range def.Properties {
		mappedProps = append(mappedProps, generateObjectProperty(def.Key, prop))
	}
	return fmt.Sprintf(result, def.Key, strings.Join(mappedProps, "\n"))
}

// generateClass generates a typescript class from the given definition.
func generateClass(def *parser.Definition) string {
	result := strings.TrimPrefix(templates.Class, "\n")

	// Set the class' description.
	if defDesc := def.Description; defDesc != "" {
		result = fmt.Sprintf("/** %s */\n", defDesc) + result
	}
	// Class properties.
	mappedProps := make([]string, 0, len(def.Properties))

	// Constructor properties.
	mappedConstructorProps := make([]string, 0, len(def.Properties))
	for _, prop := range sortProperties(def.Properties) {
		if strings.HasSuffix(prop.Key, "_at") || strings.Contains(def.Key, "ListMembers") {
			prop.Type = "ExtendedDate"
		}
		mappedProps = append(mappedProps, generateObjectProperty(def.Key, prop))

		constructorProp := generateClassConstructorProperty(def.Key, prop)
		mappedConstructorProps = append(mappedConstructorProps, constructorProp)
	}
	endOfClass := ""
	switch {
	case def.Key == "Address":
		endOfClass = "\n" + strings.TrimSuffix(constants.AddressClassMethods, "\n")
	}
	endOfClass += "\n}"

	return fmt.Sprintf(result,
		def.Key,
		strings.Join(mappedProps, "\n"),
		strings.Join(mappedConstructorProps, "\n"),
		endOfClass,
	)
}

func sortProperties(props []*parser.DefinitionProperty) []*parser.DefinitionProperty {
	sort.Slice(props, func(i, j int) bool {
		// Always place "id" property first.
		if props[i].Key == "id" {
			return true
		}
		// Always place dates last.
		if strings.HasSuffix(props[i].Key, "_at") {
			return false
		}
		// Otherwise, sort alphabetically.
		return props[i].Key < props[j].Key
	})
	return props
}

// generateEnum generates a typescript enum from the given definition.
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

// generateObjectProperty generates a typescript object property from the given property.
func generateObjectProperty(pKey string, prop *parser.DefinitionProperty) string {
	result := ""

	propDesc := prop.Description
	if propDesc == "" {
		// When a property is referenced as another, swagger-go will omit the comment.
		switch prop.Key {
		case "whereabouts":
			propDesc = "The member's last signed-in location."
		case "data":
			propDesc = "The response data."
		case "ok":
			propDesc = "Whether the request was successful."
		case "pagination":
			propDesc = "The pagination properties."
		case "address":
			propDesc = "The entity's address."
		case "changed_by_self":
			propDesc = "Whether the whereabouts were updated by the member."
		}
	}
	if propDesc != "" {
		result += "\t/** " + propDesc + " */\n"
	}

	requiredFlag := ""
	if propRequired := prop.Required; !propRequired {
		requiredFlag = "?"
	}
	propType := prop.Type
	switch {
	case prop.Ref != "":
		propType = prop.Ref
	case prop.Key == "country_code" &&
		!strings.Contains(pKey, "Request") &&
		!strings.Contains(pKey, "Response"):
		prop.Key = "country"
		propType = "Country"
	case propType == "integer":
		propType = "number"
	case propType == "array":
		propType += "[]"
	}

	result += fmt.Sprintf("\treadonly %s%s: %s;", prop.Key, requiredFlag, propType)
	return result
}

func generateClassConstructorProperty(key string, prop *parser.DefinitionProperty) (result string) {
	if prop.Ref != "" {
		fmt.Println(prop.Type)
	}
	switch {
	case strings.HasSuffix(prop.Key, "_at") || strings.Contains(key, "ListMembers"):
		return fmt.Sprintf("\t\tthis.%[1]s = new ExtendedDate(data.%[1]s);", prop.Key)
	case prop.Key == "country_code":
		return fmt.Sprintf("\t\tthis.%[1]s = new Country(data.%[1]s);", prop.Key, prop.Type)
	case prop.Ref != "" && prop.Type == "":
		return fmt.Sprintf("\t\tthis.%[1]s = data.%[1]s as %[2]s;", prop.Key, prop.Ref)
	case prop.Ref != "" && prop.Type == "array":
		return fmt.Sprintf("\t\tthis.%[1]s = data.%[1]s.map(item => new %[2]s(item));", prop.Key, prop.Ref)
	case prop.Ref != "":
		return fmt.Sprintf("\t\tthis.%[1]s = new %[2]s(data.%[1]s);", prop.Key, prop.Ref)
	default:
		return fmt.Sprintf("\t\tthis.%[1]s = data.%[1]s;", prop.Key)
	}
}
