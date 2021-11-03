package gen

import (
	"log"

	"openapi-gen/gen/parser"
	"openapi-gen/gen/typescript"
)

func generateFromTemplate(doc *parser.Document, extn Extension, logger *log.Logger) string {
	return typescript.GenerateFromDefinitions(doc.Definitions)
}
