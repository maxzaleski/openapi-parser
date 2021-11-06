package typescript

import (
	"fmt"
	"strings"

	"openapi-gen/gen/parser"
	"openapi-gen/gen/typescript/constants"
	"openapi-gen/gen/typescript/templates"
	"openapi-gen/internal"
)

// Generate generates the typescript types for the given spec.
func Generate(doc *parser.Document) string {
	out := []string{
		constants.Imports,
		generateStandardTypes(doc.Definitions),
		generateResponseTypes(doc.Responses),
		generateRequestTypes(doc.Paths),
		generateRestClient(doc.Hosts, doc.BasePath),
		generateAPIClient(doc.Paths),
	}
	return strings.Join(out, "\n\n")
}

// generateStandardTypes generates typescript types from the given definitions.
func generateStandardTypes(defs map[string]*parser.Definition) string {
	mappedDefs := []string{
		constants.Countries,
		constants.ExtendedDate,
		constants.GenericResponse,
		constants.SuccessResponse,
	}
	for _, k := range internal.SortMapKeysAlphabetically(defs) {
		def := internal.OverrideDefinition(defs[k])
		resultType := ""
		switch {
		case def.Type == "enum":
			resultType = generateEnum(def)
		case strings.HasSuffix(def.Key, "ResponseBody"):
			resultType = generateClassResponseBody(def)
		case isInterface(def.Key):
			resultType = generateInterface(def)
		default:
			resultType = generateClass(def)
		}
		mappedDefs = append(mappedDefs, resultType)
	}

	return strings.Join(mappedDefs, "\n\n")
}

// generateResponseTypes generates typescript types from the given responses.
func generateResponseTypes(defs map[string]*parser.Definition) string {
	mappedDefs := make([]string, 0, len(defs))
	for _, k := range internal.SortMapKeysAlphabetically(defs) {
		def := internal.OverrideResponseDefinition(defs[k])
		mappedDefs = append(mappedDefs, generateClassResponse(def))
	}
	return strings.Join(mappedDefs, "\n\n")
}

// generateRequestTypes generates typescript types from the given paths.
func generateRequestTypes(defs map[string]*parser.Path) string {
	mappedDefs := make([]string, 0)
	for _, paths := range internal.MapByPkg(defs) {
		for _, path := range paths {
			// Check if the path is suitable.
			if !internal.IsSuitedForAPIMethod(path.Parameters) {
				continue
			}

			// Request types.
			mappedDefs = append(mappedDefs, generateClassRequest(path))
			// Request validation objects.
			validationEligibleProps := make([]*parser.DefinitionProperty, 0, len(path.Parameters))
			for _, param := range path.Parameters {
				if param.In != "body" {
					continue
				}
				if internal.IsPropSuitableForValidation(param.Type) {
					validationEligibleProps = append(validationEligibleProps, param)
				}
			}
			if len(validationEligibleProps) > 0 {
				pathWithEligibleProps := *path
				pathWithEligibleProps.Parameters = validationEligibleProps
				mappedDefs = append(mappedDefs, generateRequestClassValidationObject(&pathWithEligibleProps))
			}
		}
	}

	return strings.Join(mappedDefs, "\n\n")
}

// generateRestClient generates the rest client code.
func generateRestClient(hosts []string, basePath string) string {
	return fmt.Sprintf(templates.RestClient, hosts[1], hosts[0], basePath)
}

// generateAPIClient generates the API client code for the given spec.
func generateAPIClient(defs map[string]*parser.Path) string {
	// The client's methods.
	mappedMethods := make([]string, 0, len(defs))
	for _, paths := range internal.MapByPkg(defs) {
		for _, path := range paths {
			// Check if the path is suitable.
			if !internal.IsSuitedForAPIMethod(path.Parameters) {
				continue
			}
			// API client's methods.
			mappedMethods = append(mappedMethods, generateAPIMethod(path))
		}
	}

	return fmt.Sprintf(templates.APIClient, strings.Join(mappedMethods, "\n"))
}
