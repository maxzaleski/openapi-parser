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
	mappedDefs := []string{
		strings.TrimPrefix(constants.Countries, "\n"),
		strings.TrimPrefix(constants.ExtendedDate, "\n"),
		strings.TrimPrefix(constants.GenericResponse, "\n"),
		strings.TrimPrefix(constants.SuccessResponse, "\n"),
	}
	for _, k := range utils.MapIntoSortedKeys(defs) {
		def := defs[k]

		resultDef := ""
		if def.Type == "enum" {
			resultDef = generateEnum(def)
		} else {
			if strings.Contains(def.Key, "Request") || strings.Contains(def.Key, "ListMembers") {
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
		result += "/** " + defDesc + " */\n"
	}
	result += templates.Interface

	mappedProps := make([]string, 0, len(def.Properties))
	for _, prop := range def.Properties {
		mappedProps = append(mappedProps, generateObjectProperty(def.Key, prop))
	}
	return fmt.Sprintf(result, def.Key, strings.Join(mappedProps, "\n"))
}

// generateClass generates a typescript class from the given definition.
func generateClass(def *parser.Definition) string {
	if strings.HasSuffix(def.Key, "ResponseBody") {
		return generateClassResponseBody(def)
	}

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
		mappedConstructorProps = append(mappedConstructorProps, generateClassConstructorProperty(def.Key, prop))
	}

	// Append class methods if any.
	methods := ""
	switch {
	case def.Key == "Address":
		methods = "\n" + strings.TrimSuffix(constants.AddressClassMethods, "\n")
	}

	return fmt.Sprintf(result,
		def.Key,
		strings.Join(mappedProps, "\n"),
		strings.Join(mappedConstructorProps, "\n"),
		methods,
	)
}

func generateClassResponseBody(def *parser.Definition) string {
	template := templates.ResponseBody
	className := def.Key
	classExtends := "GenericResponse"
	if len(def.Properties) == 2 {
		if strings.Contains(strings.ToLower(def.Key), "error") {
			classExtends = "GenericResponse"
			template = templates.ResponseErrorBody
		} else {
			className += "<T>"
			classExtends = "SuccessResponse<T>"
		}
	}

	return fmt.Sprintf(template,
		className,
		classExtends,
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
		switch def.Key {
		case "Colour":
			mappedEntries = append(mappedEntries, generateEnumColourProperty(entry))
		default:
			mappedEntries = append(mappedEntries,
				fmt.Sprintf("\t%s = '%s',", strcase.ToScreamingSnake(entry), entry))
		}
	}
	return fmt.Sprintf(result, def.Key, strings.Join(mappedEntries, "\n"))
}

func generateEnumColourProperty(entry string) string {
	colourName := ""
	switch entry {
	case "1":
		colourName = "AMBER"
	case "2":
		colourName = "ORANGE"
	case "3":
		colourName = "TOMATO"
	case "4":
		colourName = "RED"
	case "5":
		colourName = "CRIMSON"
	case "6":
		colourName = "PINK"
	case "7":
		colourName = "PLUM"
	case "8":
		colourName = "PURPLE"
	case "9":
		colourName = "VIOLET"
	case "10":
		colourName = "INDIGO"
	case "11":
		colourName = "BLUE"
	case "12":
		colourName = "CYAN"
	case "13":
		colourName = "TEAL"
	case "14":
		colourName = "GREEN"
	case "15":
		colourName = "GRASS"
	}
	return fmt.Sprintf("\t%s = %s,", colourName, entry)
}

// generateObjectProperty generates a typescript object property from the given property.
func generateObjectProperty(pKey string, prop *parser.DefinitionProperty) string {
	result := ""

	propDesc := prop.Description
	// When a property is referenced as another, swagger-go will omit the comment.
	if propDesc == "" {
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
		case "image_fallback":
			propDesc = "The image's fallback."
		case "colour":
			propDesc = "The view's display colour."
		}
	}
	if propDesc != "" {
		if isStandardType(pKey) {
			switch prop.Key {
			case "country_code":
				propDesc = "The entity's country."
			}
		}
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
		// re: we are only interested in the Colour enum.
		if propType == "ImageFallback" {
			propType = "Colour"
		}
	case prop.Key == "country_code" && isStandardType(pKey):
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

func isStandardType(key string) bool {
	return !strings.Contains(key, "Request") && !strings.Contains(key, "Response")
}

func generateClassConstructorProperty(key string, prop *parser.DefinitionProperty) (result string) {
	switch {
	case strings.HasSuffix(prop.Key, "_at") || strings.Contains(key, "ListMembers"):
		return fmt.Sprintf("\t\tthis.%[1]s = new ExtendedDate(data.%[1]s);", prop.Key)
	case prop.Key == "country_code":
		return fmt.Sprintf("\t\tthis.%[1]s = new Country(data.%[1]s);", prop.Key, prop.Type)
	case !hasConstructor(prop.Ref) && prop.Type == "":
		return fmt.Sprintf("\t\tthis.%[1]s = data.%[1]s as %[2]s;", prop.Key, prop.Ref)
	case hasConstructor(prop.Ref) && prop.Type == "":
		return fmt.Sprintf("\t\tthis.%[1]s = new %s(data.%[1]s);", prop.Key, prop.Ref)
	case prop.Ref != "" && prop.Type == "array":
		return fmt.Sprintf("\t\tthis.%[1]s = data.%[1]s.map(item => new %[2]s(item));", prop.Key, prop.Ref)
	case prop.Ref != "":
		return fmt.Sprintf("\t\tthis.%[1]s = new %[2]s(data.%[1]s);", prop.Key, prop.Ref)
	default:
		return fmt.Sprintf("\t\tthis.%[1]s = data.%[1]s;", prop.Key)
	}
}

func hasConstructor(ref string) bool {
	switch ref {
	case
		"APIError",
		"Colour",
		"EntityType",
		"ErrorCode",
		"ErrorType",
		"RelationshipWithMember",
		"Role":
		return false
	default:
		return true
	}
}
