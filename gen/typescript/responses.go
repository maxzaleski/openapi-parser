package typescript

import (
	"fmt"
	"strings"

	"openapi-gen/gen/parser"
	"openapi-gen/internal/utils"
)

// GenerateFromResponses generates typescript types from the given responses.
func GenerateFromResponses(resps map[string]*parser.Response) string {
	mappedDefs := make([]string, 0, len(resps))
	for _, k := range utils.MapIntoSortedKeys(resps) {
		mappedDefs = append(mappedDefs, generateResponse(resps[k]))
	}
	return strings.Join(mappedDefs, "\n\n")
}

func generateResponse(resp *parser.Response) string {
	result := "interface %s%s"

	respExtendFlag := resp.Ref
	if respExtendFlag != "" {
		respExtendFlag = " extends " + respExtendFlag
	}
	result += "{"

	mappedProps := make([]string, 0, len(resp.Properties))
	for _, prop := range resp.Properties {
		mappedProps = append(mappedProps, generateResponseProperty(prop))
	}
	if len(mappedProps) != 0 {
		result += fmt.Sprintf("\n%s\n}", strings.Join(mappedProps, "\n"))
	} else {
		result += "}"
	}
	return fmt.Sprintf(result, resp.Key, respExtendFlag)
}

func generateResponseProperty(prop *parser.ResponseProperty) string {
	result := ""

	propDesc := prop.Description
	if propDesc == "" {
		// When a property is referenced as another, swagger-go will omit the comment.
		switch prop.Key {
		case "whereabouts":
			propDesc = "The member's last signed-in location."
		}
	}
	if propDesc != "" {
		result += "\t// " + propDesc + "\n"
	}

	propType := prop.Type
	if propRef := prop.Ref; propRef != "" {
		propType = propRef
	}
	switch prop.Type {
	case "integer":
		propType = "number"
	case "array":
		propType += "[]"
	case "":
		// This cases catches all remaining, but we only want to target the ones we know are empty.
		if prop.Ref == "" {
			propType = "T"
		}
	}

	result += fmt.Sprintf("\t%s: %s;", prop.Key, propType)
	return result
}
