package gen

import (
	"os"

	"openapi-generator/gen/typescript"
	"openapi-generator/internal/output"
	"openapi-generator/internal/parser"
	"openapi-generator/internal/slog"
)

// New generates code for the given OpenAPI spec based on the given language extension.
func New(b []byte, version string, extn Extension) error {
	logger := slog.NewLogger("")
	logger.Println("DEBUG=1: logs enabled")
	logger.Println(`
README:
---
Concurrency is at play, the 'saw' and 'generated' keywords should be logged in pairs.
However, they are NOT expected to be outputted after one another.
Please double check the entirety of the logs before reaching a conclusion.
-------------------------------------------------------------------------------------------------|
`)

	doc, err := parser.NewDocument(b)
	if err != nil {
		logger.Println("parser.NewDocument:", err)
		return err
	}

	// Create directory if it doesn't exist.
	outDir := getAppropriateDestination(extn)
	if err = os.MkdirAll(outDir, 0755); err != nil {
		logger.Println("os.MkdirAll:", err)
		return err
	}

	files := (func() []*output.File {
		switch extn {
		case ".ts":
			return typescript.Generate(doc, logger)
		default:
			return nil
		}
	})()
	if err = output.CreateFiles(version, outDir, files, extn.String(), logger); err != nil {
		logger.Println(err)
		return err
	}

	return nil
}

// Extension represents a file extension.
type Extension string

func (e Extension) String() string {
	return string(e)
}

const (
	ExtensionTypescript Extension = ".ts"
)
