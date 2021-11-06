package typescript

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"openapi-gen/gen/parser"
	"openapi-gen/gen/typescript/templates"
	"openapi-gen/internal"
)

// generateRequestClassValidationObject generates a request class' validation object from the
// given definition.
func generateRequestClassValidationObject(def *parser.Path) string {
	mappedProps := make([]string, 0, len(def.Parameters))
	for _, prop := range internal.SortProperties(def.Parameters) {
		mappedProps = append(mappedProps, generateRequestValidationProperty("\t", prop))
	}

	return fmt.Sprintf(templates.RequestValidation,
		strcase.ToCamel(def.Operation),
		strings.Join(mappedProps, "\n"),
	)
}

// generateRequestValidationProperty generates a validation object property from the given
// definition.
func generateRequestValidationProperty(initialIndent string, prop *parser.DefinitionProperty) string {
	// TODO: nested objects like Address.

	result := initialIndent + strcase.ToLowerCamel(prop.Key)
	indent := "\n\t" + initialIndent

	switch prop.Type {
	case "string":
		result += ": yupString()"
		if pattern := prop.Validation.Pattern; pattern != "" {
			result += indent + appendValidationMessageToMethodCall(`.matches(/%v/`,
				internal.SetValidationMessageFromPattern(pattern),
				pattern,
			)
		}
		if maxLength := prop.Validation.MaxLength; maxLength != 0 {
			result += indent + appendValidationMessageToMethodCall(".max(%[1]d",
				internal.ValidationMessageMaxLength,
				maxLength,
			)
		}
		if minLength := prop.Validation.MinLength; minLength != 0 {
			result += indent + appendValidationMessageToMethodCall(".min(%[1]d",
				internal.ValidationMessageMinLength,
				minLength,
			)
		}
		if required := prop.Required; required {
			result += indent + fmt.Sprintf(".required('%s')", internal.ValidationMessageRequired)
		}
		switch prop.Key {
		case "email":
			result += indent + fmt.Sprintf(".email('%s')", internal.ValidationMessageEmail)
		case "avatar_url", "cover_image_url":
			result += indent + fmt.Sprintf(".url('%s')", internal.ValidationMessageURL)
		}
	case "integer":
		result += ": yupNumber()"
		if required := prop.Required; required {
			result += indent + fmt.Sprintf(".required('%s')", internal.ValidationMessageRequired)
		}
		if min := prop.Validation.Min; min != 0 {
			result += indent + appendValidationMessageToMethodCall(".min(%[1]d",
				internal.ValidationMessageMin,
				min,
			)
		}
		if max := prop.Validation.Max; max != 0 {
			result += indent + appendValidationMessageToMethodCall(".max(%[1]d",
				internal.ValidationMessageMax,
				max,
			)
		}
	case "array":
		result += ": yupArray()"
		if minItems := prop.Validation.MinItems; minItems != 0 {
			result += indent + appendValidationMessageToMethodCall(".max(%[1]d",
				internal.ValidationMessageMinItems,
				minItems,
			)
		}
		if maxItems := prop.Validation.MaxItems; maxItems != 0 {
			result += indent + appendValidationMessageToMethodCall(".max(%[1]d",
				internal.ValidationMessageMaxItems,
				maxItems,
			)
		}
	}
	return result + ","
}
