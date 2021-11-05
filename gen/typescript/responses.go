package typescript

import (
	"fmt"
	"strings"

	"openapi-gen/gen/parser"
	"openapi-gen/gen/typescript/constants"
	"openapi-gen/gen/typescript/templates"
	"openapi-gen/internal/utils"
)

// GenerateFromResponses generates typescript types from the given responses.
func GenerateFromResponses(resps map[string]*parser.Definition) string {
	mappedDefs := make([]string, 0, len(resps))
	for _, k := range utils.MapIntoSortedKeys(resps) {
		mappedDefs = append(mappedDefs, generateResponse(resps[k]))
	}
	return strings.Join(mappedDefs, "\n\n")
}

func generateResponse(def *parser.Definition) string {
	extends := def.Ref
	extendsType := def.Returns
	superData := "data"
	switch {
	case def.Key == "RegisterOrganisationResponse":
		extendsType = def.Ref + "Data"
		superData = strings.TrimSuffix(constants.ConstructorSuperRegisterOrganisation, "\n")
	case def.Returns != "":
		if strings.HasSuffix(def.Returns, "[]") {
			superData = fmt.Sprintf("{ ...data, data: data.data.map(e => new %s(e))}",
				strings.TrimSuffix(def.Returns, "[]"))
		} else {
			superData = fmt.Sprintf("{ ...data, data: new %s(data.data)}", def.Returns)
		}
	}
	if extendsType != "" && !strings.Contains(strings.ToLower(def.Key), "error") {
		extends += "<" + extendsType + ">"
	}
	return fmt.Sprintf(strings.TrimPrefix(templates.Response, "\n"),
		def.Key,
		extends,
		superData,
	)
}
