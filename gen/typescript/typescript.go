package typescript

import (
	"openapi-gen/internal/output"
	"openapi-gen/internal/parser"

	"openapi-gen/gen/typescript/constants"
	"openapi-gen/internal/slog"
)

const definitionsOutDir = "definitions/"

// Generate generates the typescript files for the given spec.
func Generate(doc *parser.Document, logger slog.Logger) []*output.File {
	logger.SetPrefix("[Generate] ")

	fm := output.NewFileMap()
	generateOutput(doc, fm, logger)

	// .
	// ├── definitions
	// │   ├── index.ts
	// │   └── m.ts
	// │   └── e.ts
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
