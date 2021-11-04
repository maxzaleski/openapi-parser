package typescript

import (
	"strings"

	"openapi-gen/gen/parser"
)

func Generate(doc *parser.Document) string {
	types := []string{
		GenerateFromDefinitions(doc.Definitions),
		GenerateFromResponses(doc.Responses),
		GenerateFromPaths(doc.Hosts, doc.BasePath, doc.Paths),
	}
	return strings.Join(types, "\n\n")
}
