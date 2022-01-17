package typescript

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"openapi-generator/gen/typescript/constants"
	"openapi-generator/gen/typescript/templates"
	"openapi-generator/internal"
	"openapi-generator/internal/parser"
	"openapi-generator/internal/slog"
	"regexp"
	"strings"
)

var dynanicQueryFilterRegex = regexp.MustCompile(`[aA-zZ]+DynamicQueryFilter[A-Z][aA-zZ]+`)

// generateModelTypes generates typescript types from the given definitions.
func generateModelTypes(defs map[string]*parser.Definition, logger slog.Logger) string {
	mappedDefs := []string{
		constants.ModelsImports,
		constants.ExtendedDate,
		constants.DynamicQueryFilterGeneric,
	}
	for _, k := range internal.SortMapKeysAlphabetically(defs) {
		def := internal.OverrideDefinition(defs[k])
		logger.Printf("saw '%s'", def.Key)

		resultType := ""
		switch {
		case dynanicQueryFilterRegex.MatchString(k):
			// Individual filters are not included because they all follow the same interface which is included as a constant.
			continue
		case strings.HasSuffix(k, "DynamicQueryFilters"):
			resultType = generateDynamicQueryFilters(defs, def)
		case isInterface(def.Key) || strings.HasSuffix(def.Key, "Data"):
			resultType = generateInterface(def)
		default:
			resultType = generateClass(def)
		}

		logger.Printf("generated '%s'", def.Key)
		mappedDefs = append(mappedDefs, resultType)
	}
	logger.Printf("[generateModelTypes] received %d :: mapped %d", len(defs), len(mappedDefs)-3)

	return strings.Join(mappedDefs, "\n\n")
}

// generateEnumTypes generates typescript enum types from the given definitions.
func generateEnumTypes(defs map[string]*parser.Definition, logger slog.Logger) string {
	sliceLen := len(defs) + 1
	mappedEnums := make([]string, 0, sliceLen)
	for _, k := range internal.SortMapKeysAlphabetically(defs) {
		def := internal.OverrideDefinition(defs[k])
		logger.Printf("saw enum '%s'", def.Key)

		mappedEnums = append(mappedEnums, generateEnum(def))
		logger.Printf("generated enum '%s'", def.Key)
	}
	logger.Printf("[generateEnumTypes] received %d :: mapped %d", sliceLen, len(mappedEnums))

	return strings.Join(mappedEnums, "\n\n")
}

// generateResponseTypes generates typescript types from the given responses.
func generateResponseTypes(defs map[string]*parser.Definition, logger slog.Logger) string {
	mappedDefs := make([]string, 0, len(defs)+4)
	mappedDefs = append(mappedDefs,
		constants.ResponsesImports,
		constants.GenericResponse,
		constants.SuccessResponse,
		constants.PaginatedResponse,
	)
	for _, k := range internal.SortKeysByCase(internal.SortMapKeysAlphabetically(defs)) {
		def := defs[k]
		logger.Printf("saw response '%s'", def.Key)

		if strings.HasSuffix(def.Key, "Body") {
			mappedDefs = append(mappedDefs, generateClassResponseBody(def))
		} else {
			mappedDefs = append(mappedDefs, generateClassResponse(def))
		}
		logger.Printf("generated response '%s'", def.Key)
	}
	logger.Printf("[generateResponseTypes] received %d :: mapped %d", len(defs), len(mappedDefs)-4)

	return strings.Join(mappedDefs, "\n\n")
}

// generateRequestTypes generates typescript types from the given paths.
func generateRequestTypes(defs map[string]*parser.Path, reqBodies map[string]*parser.Definition, logger slog.Logger) string {
	mappedDefs := make([]string, 0, len(defs)+1)
	mappedDefs = append(mappedDefs, constants.RequestsImports)
	for _, def := range reqBodies {
		def = internal.OverrideDefinition(def)
		def.Key = strcase.ToLowerCamel(def.Key)
		mappedDefs = append(mappedDefs, generateInterface(def))
	}
	for _, paths := range internal.MapByPkg(defs) {
		for _, path := range paths {
			logger.Printf("saw path '%s'", path.Key)
			if !internal.IsSuitedForAPIMethod(path.Parameters) {
				continue
			}
			mappedDefs = append(mappedDefs, generateClassRequest(path))
			logger.Printf("generated path '%s'", path.Key)
		}
	}
	logger.Printf("[generateRequestTypes] received %d :: mapped %d", len(defs), len(mappedDefs)-1)

	return strings.Join(mappedDefs, "\n\n")
}

// generateRequestValidationObjects generates typescript validation objects from the given
// validation properties.
func generateRequestValidationObjects(
	defs map[string][]*parser.DefinitionProperty, logger slog.Logger,
) string {
	mappedObjects := make([]string, 0, len(defs)+1)
	mappedObjects = append(mappedObjects, constants.ValidationImports)
	for k, v := range defs {
		logger.Printf("saw validation object '%s'", k)
		mappedObjects = append(mappedObjects, generateRequestClassValidationObject(k, v))
		logger.Printf("generated validation object '%s'", k)
	}
	logger.Printf("[generateRequestValidationObjects] received %d :: mapped %d", len(defs), len(mappedObjects)-1)

	return strings.Join(mappedObjects, "\n\n")
}

// generateRestClient generates the rest client code.
func generateRestClient(host, basePath string, _ slog.Logger) string {
	return fmt.Sprintf(templates.RestClient, host, basePath)
}

// generateAPIClient generates the API client code for the given spec.
func generateAPIClient(defs map[string]*parser.Path, logger slog.Logger) string {
	// The client's methods.
	mappedMethods := make([]string, 0, len(defs))
	for _, paths := range internal.MapByPkg(defs) {
		for _, path := range paths {
			logger.Printf("saw method '%s'", path.Key)
			mappedMethods = append(mappedMethods, generateAPIMethod(path))
			logger.Printf("generated method '%s'", path.Key)
		}
	}
	logger.Printf("[generateAPIClient] received %d :: mapped %d", len(defs), len(mappedMethods))

	return fmt.Sprintf(templates.APIClient, strings.Join(mappedMethods, "\n\n"))
}
