package typescript

import (
	"openapi-generator/internal/output"
	"openapi-generator/internal/parser"

	"openapi-generator/gen/typescript/constants"
	"openapi-generator/internal/slog"
)

const definitionsOutDir = "definitions/"

// Generate generates the typescript files for the given spec.
func Generate(doc *parser.Document, logger slog.Logger) []*output.File {
	logger.SetPrefix("[typescript.Generate] ")

	fm := output.NewFileMap()
	generateOutput(doc, fm, logger)

	// ../packages/
	// ├── definitions
	// │   ├── index.ts
	// │   └── models.ts
	// │   └── enums.ts
	// │   └── responses.ts
	// │   └── requests.ts
	// │   └── validation.ts
	// │   └── countries.ts
	// └── index.ts
	// └── rest-client.ts
	// └── api-client.ts
	files := make([]*output.File, 0, 10)
	files = append(files,
		&output.File{
			Name: "index",
			Body: constants.RootIndex,
		},
		&output.File{
			Name:      "index",
			Directory: definitionsOutDir,
			Body:      constants.DefinitionsIndex,
		},
		&output.File{
			Name:      "countries",
			Directory: definitionsOutDir,
			Body:      constants.Countries,
		},
	)
	for _, k := range fm.Range() {
		files = append(files, fm.Get(k))
	}

	return files
}
