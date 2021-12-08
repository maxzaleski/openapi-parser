package gen

import (
	"os"

	"openapi-gen/gen/typescript"
	"openapi-gen/internal/output"
	"openapi-gen/internal/parser"
	"openapi-gen/internal/slog"
)

// New generates code for the given OpenAPI spec based on the given language extension.
func New(b []byte, extn Extension) error {
	logger := slog.NewLogger("[gen]")

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
	if err = output.CreateFiles(outDir, files, extn.String(), logger); err != nil {
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
