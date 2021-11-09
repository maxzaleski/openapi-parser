package typescript

import (
	"fmt"
	"strings"
	"sync"

	"openapi-gen/internal/output"
	"openapi-gen/internal/parser"

	"openapi-gen/gen/typescript/constants"
	"openapi-gen/gen/typescript/templates"
	"openapi-gen/internal"
	"openapi-gen/internal/slog"
)

// generateOutput concurrently writes the generated code to the file map.
func generateOutput(doc *parser.Document, m output.FileMap, logger slog.Logger) {
	enums := make(map[string]*parser.Definition)
	internal.ShakeModelDefinitions(doc.Definitions, enums)
	internal.ShakeResponseBodyDefinitions(doc.Definitions, doc.Responses)
	validationObjectMap := internal.FilterIntoValidationObjectMap(doc.Paths)

	jobs := []output.Job{
		{
			OperationID: "rest-client",
			Generator:   "generateRESTClient",
			Args:        []interface{}{doc.Hosts, doc.BasePath, logger},
		},
		{
			OperationID: "api-client",
			Generator:   "generateAPIClient",
			Args:        []interface{}{doc.Paths, logger},
		},
		{
			OperationID: "models",
			Generator:   "generateModelTypes",
			Args:        []interface{}{doc.Definitions, logger},
		},
		{
			OperationID: "enums",
			Generator:   "generateEnumTypes",
			Args:        []interface{}{enums, logger},
		},
		{
			OperationID: "requests",
			Generator:   "generateRequestTypes",
			Args:        []interface{}{doc.Paths, logger},
		},
		{
			OperationID: "validation",
			Generator:   "generateRequestValidationObjects",
			Args:        []interface{}{validationObjectMap, logger},
		},
		{
			OperationID: "responses",
			Generator:   "generateResponseTypes",
			Args:        []interface{}{doc.Responses, logger},
		},
	}
	output.New(jobs, outputWorker, m)
}

func outputWorker(
	wg *sync.WaitGroup, jobs <-chan output.Job, m output.FileMap,
) {
	defer wg.Done()

	for job := range jobs {
		var file *output.File
		switch job.OperationID {
		case "rest-client":
			hosts := job.Args[0].([]string)
			basePath := job.Args[1].(string)
			logger := job.Args[2].(slog.Logger)
			file = &output.File{
				Name: "rest-client",
				Body: generateRestClient(hosts, basePath, logger),
			}
		case "api-client":
			paths := job.Args[0].(map[string]*parser.Path)
			logger := job.Args[1].(slog.Logger)
			file = &output.File{
				Name: "api-client",
				Body: generateAPIClient(paths, logger),
			}
		case "models":
			defs := job.Args[0].(map[string]*parser.Definition)
			logger := job.Args[1].(slog.Logger)
			file = &output.File{
				Name:      "models",
				Directory: definitionsOutDir,
				Body:      generateModelTypes(defs, logger),
			}
		case "enums":
			defs := job.Args[0].(map[string]*parser.Definition)
			logger := job.Args[1].(slog.Logger)
			file = &output.File{
				Name:      "enums",
				Directory: definitionsOutDir,
				Body:      generateEnumTypes(defs, logger),
			}
		case "requests":
			defs := job.Args[0].(map[string]*parser.Path)
			logger := job.Args[1].(slog.Logger)
			file = &output.File{
				Name:      "requests",
				Directory: definitionsOutDir,
				Body:      generateRequestTypes(defs, logger),
			}
		case "validation":
			defs := job.Args[0].(map[string][]*parser.DefinitionProperty)
			logger := job.Args[1].(slog.Logger)
			file = &output.File{
				Name:      "validation",
				Directory: definitionsOutDir,
				Body:      generateRequestValidationObjects(defs, logger),
			}
		case "responses":
			defs := job.Args[0].(map[string]*parser.Definition)
			logger := job.Args[1].(slog.Logger)
			body := generateResponseTypes(defs, logger)
			file = &output.File{
				Name:      "responses",
				Directory: definitionsOutDir,
				Body:      body,
			}
		}
		m.Add(file)
	}
}

// generateModelTypes generates typescript types from the given definitions.
func generateModelTypes(defs map[string]*parser.Definition, logger slog.Logger) string {
	mappedDefs := []string{
		constants.ModelsImports,
		constants.ExtendedDate,
	}
	for _, k := range internal.SortMapKeysAlphabetically(defs) {
		def := internal.OverrideDefinition(defs[k])
		logger.Printf("saw '%s'", def.Key)

		resultType := ""
		switch {
		case isInterface(def.Key) || strings.HasSuffix(def.Key, "Data"):
			resultType = generateInterface(def)
		case strings.Contains(def.Key, "Response"):
			resultType = generateClassResponse(def)
		case strings.HasSuffix(def.Key, "Body"):
			resultType = generateClassResponseBody(def)
		default:
			resultType = generateClass(def)
		}

		logger.Printf("generated '%s'", def.Key)
		mappedDefs = append(mappedDefs, resultType)
	}
	logger.Printf("[generateModelTypes] received %d :: mapped %d", len(defs), len(mappedDefs))

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
	logger.Printf("[generateResponseTypes] received %d :: mapped %d", len(defs), len(mappedDefs))

	return strings.Join(mappedDefs, "\n\n")
}

// generateRequestTypes generates typescript types from the given paths.
func generateRequestTypes(defs map[string]*parser.Path, logger slog.Logger) string {
	mappedDefs := make([]string, 0, len(defs)+1)
	mappedDefs = append(mappedDefs, constants.RequestsImports)
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
	logger.Printf("[generateRequestTypes] received %d :: mapped %d", len(defs), len(mappedDefs))

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
	logger.Printf("[generateRequestValidationObjects] received %d :: mapped %d", len(defs), len(mappedObjects))

	return strings.Join(mappedObjects, "\n\n")
}

// generateRestClient generates the rest client code.
func generateRestClient(hosts []string, basePath string, logger slog.Logger) string {
	return fmt.Sprintf(templates.RestClient, hosts[1], hosts[0], basePath)
}

// generateAPIClient generates the API client code for the given spec.
func generateAPIClient(defs map[string]*parser.Path, logger slog.Logger) string {
	// The client's methods.
	mappedMethods := make([]string, 0, len(defs))
	for _, paths := range internal.MapByPkg(defs) {
		for _, path := range paths {
			// Check if the path is suitable.
			// if !internal.IsSuitedForAPIMethod(path.Parameters) {
			// 	continue
			// }
			// API client's methods.
			logger.Printf("saw method '%s'", path.Key)
			mappedMethods = append(mappedMethods, generateAPIMethod(path))
			logger.Printf("generated method '%s'", path.Key)
		}
	}
	logger.Printf("[generateAPIClient] received %d :: mapped %d", len(defs), len(mappedMethods))

	return fmt.Sprintf(templates.APIClient, strings.Join(mappedMethods, "\n"))
}
