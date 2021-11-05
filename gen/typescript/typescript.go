package typescript

import (
	"strings"

	"openapi-gen/gen/parser"
	"openapi-gen/gen/typescript/constants"
)

func Generate(doc *parser.Document) string {
	out := []string{
		constants.Imports,
		GenerateFromDefinitions(doc.Definitions),
		GenerateFromResponses(doc.Responses),
		// GenerateFromPaths(doc.Hosts, doc.BasePath, doc.Paths),
	}
	return strings.Join(out, "\n\n")
}
