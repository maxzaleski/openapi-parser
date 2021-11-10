package typescript

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"openapi-gen/internal/parser"

	"openapi-gen/gen/typescript/constants"
	"openapi-gen/gen/typescript/templates"
	"openapi-gen/internal"
)

// generateClass generates a typescript class from the given definition.
func generateClass(def *parser.Definition) string {
	template := templates.Class

	// Class description.
	if defDesc := def.Description; defDesc != "" {
		template = toJSDoc("", defDesc) + template
	}
	// Class properties.
	mappedProps := make([]string, 0, len(def.Properties))
	// Class constructor properties.
	mappedConstructorProps := make([]string, 0, len(def.Properties))
	for _, prop := range internal.SortProperties(def.Properties) {
		mappedProps = append(mappedProps, generateObjectProperty("", prop))
		mappedConstructorProps = append(mappedConstructorProps, generateClassConstructorProperty(def.Key, prop))
	}
	// Class methods if any.
	classMethods := ""
	switch {
	case def.Key == "Address":
		classMethods = constants.AddressClassMethods
	}

	return fmt.Sprintf(template,
		def.Key,
		strings.Join(mappedProps, "\n"),
		strings.Join(mappedConstructorProps, "\n"),
		classMethods,
	)
}

// generateClassRequest generates a typescript request class for the given definition.
func generateClassRequest(path *parser.Path) string {
	mappedProps := make([]string, 0, len(path.Parameters))
	for _, prop := range internal.SortProperties(path.Parameters) {
		mappedProps = append(mappedProps, generateObjectProperty("m.", prop))
	}

	return fmt.Sprintf(templates.Request,
		strcase.ToCamel(path.Operation),
		strings.Join(mappedProps, "\n"),
	)
}

// generateClassResponse generates a typescript response class for the given definition.
func generateClassResponse(def *parser.Definition) string {
	// Class's extend type.
	extends := def.Ref
	if !internal.IsErrorType(def.Key) {
		if extendsType := def.Returns; extendsType != "" {
			extends += "<m." + extendsType + ">"
		}
	}
	// Class constructor's super call arguments.
	constructorSuperArgs := setConstructorSuperArgs(def)

	return fmt.Sprintf(templates.Response,
		def.Key,
		extends,
		constructorSuperArgs,
	)
}

// generateClassResponseBody generates a typescript response body class for the given definition.
func generateClassResponseBody(def *parser.Definition) string {
	template := templates.ResponseBody

	className := def.Key
	classExtends := "GenericResponse"
	if len(def.Properties) >= 2 {
		if strings.Contains(strings.ToLower(def.Key), "error") {
			classExtends = "GenericResponse"
			template = templates.ResponseErrorBody
		} else {
			className += "<T>"
			if internal.IsPaginatedResponse(def.Key) {
				classExtends = "PaginatedResponse<T>"
			} else {
				classExtends = "SuccessResponse<T>"
			}
		}
	}

	return fmt.Sprintf(template,
		className,
		classExtends,
	)
}

// setConstructorSuperArgs returns the super call arguments based on the given definition.
func setConstructorSuperArgs(def *parser.Definition) string {
	switch {
	case def.Key == "RegisterOrganisationResponse":
		return constants.ConstructorSuperRegisterOrganisation
	case def.Returns != "":
		if strings.HasSuffix(def.Returns, "[]") {
			return fmt.Sprintf("{ ...data, data: data.data.map(e => new m.%s(e))}",
				strings.TrimSuffix(def.Returns, "[]"))
		} else {
			return fmt.Sprintf("{ ...data, data: new m.%s(data.data)}", def.Returns)
		}
	default:
		return "data"
	}
}

// generateObjectProperty generates a typescript object property from the given definition.
func generateObjectProperty(prefix string, prop *parser.DefinitionProperty) string {
	template := templates.ObjectProperty

	// Property's description.
	//
	// When a property is referenced, swagger-go will omit the comment.
	if propDesc := generateObjectPropertyMissingComment(prop.Description); propDesc != "" {
		template = toJSDoc("\t", propDesc) + template
	}

	// Property's required flag.
	requiredFlag := ""
	if propRequired := prop.Required; !propRequired {
		requiredFlag = "?"
	}
	// Property's type.
	propType := prop.Type
	switch {
	// Account for enums imported from './enums'.
	case !internal.HasConstructor(prop.Ref):
		propType = "e." + prop.Ref
	case prop.Ref != "" && prop.Type != "array":
		propType = prefix + prop.Ref
	default:
		switch prop.Type {
		case "integer":
			propType = "number"
		case "array":
			if internal.IsBasicType(prop.Ref) {
				propType = prop.Ref + "[]"
			} else {
				propType = prefix + prop.Ref + "[]"
			}
		}
	}

	return fmt.Sprintf(template, prop.Key, requiredFlag, propType)
}

// generateObjectPropertyMissingComment generates a comment for the given property.
func generateObjectPropertyMissingComment(key string) string {
	switch key {
	case "whereabouts":
		return "The member's last signed-in location."
	case "data":
		return "The response data."
	case "ok":
		return "Whether the request was successful."
	case "pagination":
		return "The pagination properties."
	case "address":
		return "The entity's address."
	case "changed_by_self":
		return "Whether the whereabouts were updated by the member."
	case "image_fallback":
		return "The image's fallback."
	case "colour":
		return "The view's display colour."
	default:
		return key
	}
}

// generateClassConstructorProperty generates a typescript class constructor property from the given
// definition.
func generateClassConstructorProperty(key string, prop *parser.DefinitionProperty) (result string) {
	switch {
	case strings.HasSuffix(prop.Key, "_at") && !strings.Contains(key, "ListMembers"):
		return fmt.Sprintf("\t\tthis.%[1]s = new ExtendedDate(data.%[1]s);", prop.Key)
	case prop.Key == "country_code":
		return fmt.Sprintf("\t\tthis.%[1]s = new Country(data.%[1]s);", prop.Key, prop.Type)
	case internal.HasConstructor(prop.Ref) && prop.Type == "":
		return fmt.Sprintf("\t\tthis.%[1]s = new %s(data.%[1]s);", prop.Key, prop.Ref)
	case !internal.HasConstructor(prop.Ref) && prop.Type == "":
		return fmt.Sprintf("\t\tthis.%[1]s = data.%[1]s as e.%[2]s;", prop.Key, prop.Ref)
	case prop.Ref != "" && prop.Type == "array":
		return fmt.Sprintf("\t\tthis.%[1]s = data.%[1]s.map(item => new %[2]s(item));", prop.Key, prop.Ref)
	case prop.Ref != "":
		return fmt.Sprintf("\t\tthis.%[1]s = new %[2]s(data.%[1]s);", prop.Key, prop.Ref)
	default:
		return fmt.Sprintf("\t\tthis.%[1]s = data.%[1]s;", prop.Key)
	}
}
