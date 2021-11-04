package gen

import (
	"log"

	"openapi-gen/gen/parser"
	"openapi-gen/gen/typescript"
)

func generateFromDocument(doc *parser.Document, extn Extension, logger *log.Logger) string {
	switch extn {
	case ".ts":
		return typescript.Generate(doc)
	default:
		return ""
	}
}
