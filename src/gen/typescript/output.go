package typescript

import (
	"sync"

	"openapi-generator/internal"
	"openapi-generator/internal/output"
	"openapi-generator/internal/parser"
	"openapi-generator/internal/slog"
)

// generateOutput concurrently writes the generated code to the file map.
func generateOutput(doc *parser.Document, m output.FileMap, logger slog.Logger) {
	enums := make(map[string]*parser.Definition)
	reqBodies := make(map[string]*parser.Definition)
	internal.ShakeRequestBodyDefinitions(doc.Definitions, reqBodies)
	internal.ShakeModelDefinitions(doc.Definitions, enums)
	internal.ShakeResponseBodyDefinitions(doc.Definitions, doc.Responses)
	validationObjectMap := internal.FilterIntoValidationObjectMap(reqBodies, doc.Paths)

	jobs := []output.Job{
		{
			OperationID: "rest-client",
			Generator:   "generateRESTClient",
			Args:        []interface{}{doc.Host, doc.BasePath, logger},
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
			Args:        []interface{}{doc.Paths, reqBodies, logger},
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
			host := job.Args[0].(string)
			basePath := job.Args[1].(string)
			logger := job.Args[2].(slog.Logger)
			file = &output.File{
				Name: "rest-client",
				Body: generateRestClient(host, basePath, logger),
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
			reqBodies := job.Args[1].(map[string]*parser.Definition)
			logger := job.Args[2].(slog.Logger)
			file = &output.File{
				Name:      "requests",
				Directory: definitionsOutDir,
				Body:      generateRequestTypes(defs, reqBodies, logger),
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
