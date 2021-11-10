package typescript

import (
	"fmt"
	"regexp"

	"github.com/iancoleman/strcase"

	"openapi-gen/internal/parser"

	"openapi-gen/gen/typescript/templates"
	"openapi-gen/internal"
)

// routePathParamRegex is a regexp that extracts path parameters from a route path.
var routePathParamRegex = regexp.MustCompile(`{([a-z_]+)}`)

// generateAPIMethod generates a typescript class method from the given definition.
func generateAPIMethod(def *parser.Path) string {
	template := templates.APIClientMethod
	// The method's description.
	if desc := def.Description; desc != "" {
		template = toJSDoc("\t", desc) + template
	}
	// The method's name but capitalised under camel case.
	operationAsCamel := strcase.ToCamel(def.Operation)

	// The method's arguments.
	methodArgs := ""
	// The method's path.
	methodPath := "'" + routePathParamRegex.ReplaceAllString(def.Key, "") + "'"
	// The method's path parameter.
	routePathParam := ""
	regexResult := routePathParamRegex.FindStringSubmatch(def.Key)
	if len(regexResult) == 2 {
		routePathParam = internal.ShortenPathParam(regexResult[1])

		methodPath += " + " + routePathParam
		methodArgs = routePathParam + ": string"
	}

	var (
		flagPayload bool
	)
	// Check if a payload is required; We assume a payload if the conditions are met.
	if len(def.Parameters) > 1 || len(def.Parameters) == 1 && def.Parameters[0].In == "body" {
		flagPayload = true

		if routePathParam != "" {
			methodArgs += ", "
		}
		methodArgs += "payload: deoutput." + operationAsCamel + "Request"
	}

	// Method's rest client call.
	methodRestFunction := def.HTTPVerb
	if methodRestFunction == "list" {
		methodRestFunction = "get"
	}
	// Assign the rest client method's generics.
	methodRestFunctionGenerics := "void, deoutput." + operationAsCamel + "Response"
	// Assign the rest client method's arguments.
	methodRestFunctionArgs := "path"
	if flagPayload {
		methodRestFunctionArgs += ", payload"
		methodRestFunctionGenerics = fmt.Sprintf("deoutput.%[1]sRequest, deoutput.%[1]sResponse", operationAsCamel)
	}

	return fmt.Sprintf(template,
		def.Operation,
		methodArgs,
		"deoutput."+operationAsCamel+"Response",
		methodPath,
		methodRestFunction,
		methodRestFunctionGenerics,
		methodRestFunctionArgs,
	)
}
