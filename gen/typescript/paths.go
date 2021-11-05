package typescript

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"

	"openapi-gen/gen/parser"
	"openapi-gen/gen/typescript/templates"
)

// GenerateFromPaths generates typescript types from the given paths.
func GenerateFromPaths(hosts []string, basePath string, paths map[string]*parser.Path) string {
	mappedDefs := make([]string, 0)
	mappedMethods := make([]string, 0, len(paths))
	for _, v := range mapByPkg(paths) {
		for _, path := range v {
			if len(path.Parameters) < 1 || len(path.Parameters) == 1 && path.Parameters[0].In == "path" {
				continue
			}

			// Request definitions.
			mappedDefs = append(mappedDefs, generateRequest(path))
			// Request validation.
			validationEligibleProps := make([]*parser.DefinitionProperty, 0, len(path.Parameters))
			for _, param := range path.Parameters {
				if param.In != "body" {
					continue
				}
				switch param.Type {
				case "string", "integer", "array":
					validationEligibleProps = append(validationEligibleProps, param)
				}
			}
			if len(validationEligibleProps) > 0 {
				pathWithEligibleProps := *path
				pathWithEligibleProps.Parameters = validationEligibleProps
				mappedDefs = append(mappedDefs, generateRequestValidation(&pathWithEligibleProps))
			}
			// API client methods.
			mappedMethods = append(mappedMethods, generateMethod(path))
		}
	}
	mappedDefs = append(mappedDefs,
		fmt.Sprintf(templates.HTTPClient, hosts[1], hosts[0], basePath),
		fmt.Sprintf(templates.APIClient, strings.Join(mappedMethods, "\n")),
	)
	return strings.Join(mappedDefs, "\n")
}

var (
	routePathParamRegex = regexp.MustCompile(`{([a-z_]+)}`)
)

func generateMethod(path *parser.Path) string {
	// The method's name but capitalised under camel case.
	operationToCamel := strcase.ToCamel(path.Operation)

	// Extract the path parameters if any.
	routePathParam := ""
	regexResult := routePathParamRegex.FindStringSubmatch(path.Key)
	if len(regexResult) == 2 {
		routePathParam = shortenRoutePathParam(regexResult[1])
	}
	// Append the path parameter if any.
	funcConstPath := "'" + routePathParamRegex.ReplaceAllString(path.Key, "") + "'"
	if routePathParam != "" {
		funcConstPath += " + " + routePathParam
	}

	// Extract the function arguments' definitions.
	var (
		flagPayload  bool
		funcArgsDefs string
	)
	if routePathParam != "" {
		funcArgsDefs = routePathParam + ": string"
	}
	// We assume a payload if the conditions are met.
	if len(path.Parameters) > 1 || len(path.Parameters) == 1 && path.Parameters[0].In == "body" {
		flagPayload = true

		if routePathParam != "" {
			funcArgsDefs += ", "
		}
		funcArgsDefs += "payload: " + operationToCamel + "Request"
	}

	// Assign the HTTP client's method to call.
	funcClientMethod := path.HTTPVerb
	if funcClientMethod == "list" {
		funcClientMethod = "get"
	}
	// Assign the HTTP client method's generics.
	funcClientMethodGenerics := "void, " + operationToCamel + "Response"
	// Assign the HTTP client method's arguments.
	funcClientMethodArgs := "path"
	if flagPayload {
		funcClientMethodArgs += ", payload"
		funcClientMethodGenerics = fmt.Sprintf("%[1]sRequest, %[1]sResponse", operationToCamel)
	}

	return fmt.Sprintf(templates.APIClientMethod,
		path.Operation,
		funcArgsDefs,
		operationToCamel+"Response",
		funcConstPath,
		funcClientMethod,
		funcClientMethodGenerics,
		funcClientMethodArgs,
	)
}

func shortenRoutePathParam(param string) string {
	switch param {
	case "accommodation_id":
		return "accom_id"
	case "organisation_id":
		return "org_id"
	case "member_id":
		return "mbr_id"
	case "location_id":
		return "loc_id"
	case "address_id":
		return "addr_id"
	case "group_id":
		return "grp_id"
	default:
		return param
	}
}

func generateRequest(path *parser.Path) string {
	mappedProps := make([]string, 0, len(path.Parameters))
	for _, prop := range sortProperties(path.Parameters) {
		mappedProps = append(mappedProps, generateRequestObjectProperty(prop))
	}
	return fmt.Sprintf(templates.Request,
		strcase.ToCamel(path.Operation),
		strings.Join(mappedProps, "\n"),
	)
}

func generateRequestValidation(path *parser.Path) string {
	mappedProps := make([]string, 0)
	for _, prop := range sortProperties(path.Parameters) {
		mappedProps = append(mappedProps, generateRequestValidationProperty("\t", prop))
	}
	return fmt.Sprintf(templates.RequestValidation,
		strcase.ToCamel(path.Operation),
		strings.Join(mappedProps, "\n"),
	)
}

func generateRequestValidationProperty(initialIndent string, prop *parser.DefinitionProperty) string {
	result := initialIndent + strcase.ToLowerCamel(prop.Key)
	indent := "\n\t" + initialIndent

	if prop.Ref != "" {

	}

	switch prop.Type {
	case "string":
		result += ": yupString()"
		if pattern := prop.Validation.Pattern; pattern != "" {
			result += indent + fmt.Sprintf(
				// TODO: Message based on pattern
				`.matches(/%s/, 'This field doesn\'t match the required pattern.')`, pattern)
		}
		if maxLength := prop.Validation.MaxLength; maxLength != 0 {
			result += indent + fmt.Sprintf(
				".max(%[1]d, 'This field allows a maximum of %[1]d chars.')", maxLength)
		}
		if minLength := prop.Validation.MinLength; minLength != 0 {
			result += indent + fmt.Sprintf(
				".min(%[1]d, 'This field requires a minimum of %[1]d chars.')", minLength)
		}
		if required := prop.Required; required {
			result += indent + fmt.Sprintf(".required('This field is required.')")
		}
		if format := prop.Format; format != "" {
			switch prop.Format {
			case "email":
				result += indent + fmt.Sprintf(".email('This field must be a valid email address.')")
			case "uri":
				result += indent + fmt.Sprintf(".url('This field must be a valid URL.')")
			}
		}
	case "integer":
		result += ": yupNumber()"
		if required := prop.Required; required {
			result += indent + fmt.Sprintf(".required('This field is required.')")
		}
		if min := prop.Validation.Min; min != 0 {
			result += indent + fmt.Sprintf(".min(%[1]d, 'This field requires a minimum of %[1]d.')", min)
		}
		if max := prop.Validation.Max; max != 0 {
			result += indent + fmt.Sprintf(".max(%[1]d, 'This field allows a maximum of %[1]d.')", max)
		}
	case "array":
		result += ": yupArray()"
		if minItems := prop.Validation.MinItems; minItems != 0 {
			result += indent + fmt.Sprintf(
				".min(%[1]d, 'This field requires a minimum of %[1]d items.')", minItems)
		}
		if maxItems := prop.Validation.MaxItems; maxItems != 0 {
			result += indent + fmt.Sprintf(
				".max(%[1]d, 'This field allows a maximum of %[1]d items.')", maxItems)
		}
	}
	return result + ","
}

// generateRequestObjectProperty generates a typescript object property from the given property.
func generateRequestObjectProperty(prop *parser.DefinitionProperty) string {
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
	case propType == "integer":
		propType = "number"
	case propType == "array":
		propType += "[]"
	}

	result += fmt.Sprintf("\treadonly %s%s: %s;", prop.Key, requiredFlag, propType)
	return result
}

func mapByPkg(paths map[string]*parser.Path) map[string][]*parser.Path {
	pkgMap := make(map[string][]*parser.Path)
	for k, v := range paths {
		key := strings.Split(k, ".")[0]
		if slice, ok := pkgMap[key]; ok {
			pkgMap[key] = append(slice, v)
		} else {
			pkgMap[key] = []*parser.Path{v}
		}
	}
	return pkgMap
}
