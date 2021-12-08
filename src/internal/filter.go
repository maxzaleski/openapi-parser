package internal

import (
	"openapi-gen/internal/parser"
)

// FilterIntoModelsAndEnumKeys filters the given map into 2 slices of keys.
//
// @returns (m, e): m -> models, e -> enums
func FilterIntoModelsAndEnumKeys(m map[string]*parser.Definition) ([]string, []string) {
	models := make([]string, 0)
	enums := make([]string, 0)
	for k, v := range m {
		if v.Type == "enum" {
			enums = append(enums, k)
		} else {
			models = append(models, k)
		}
	}
	return models, enums
}

// FilterIntoValidationObjectMap filters the given map of paths into a map of definition
// properties suitable for validation.
func FilterIntoValidationObjectMap(m map[string]*parser.Path) map[string][]*parser.DefinitionProperty {
	validations := make(map[string][]*parser.DefinitionProperty, 0)
	for _, v := range m {
		// Check if the path is suitable.
		if !IsSuitedForAPIMethod(v.Parameters) {
			continue
		}

		for _, param := range v.Parameters {
			if param.In == "path" {
				continue
			}
			validationEligibleProps := make([]*parser.DefinitionProperty, 0, len(v.Parameters))
			for _, param = range v.Parameters {
				if param.In == "path" {
					continue
				}
				if IsPropSuitableForValidation(param.Type) {
					validationEligibleProps = append(validationEligibleProps, param)
				}
			}
			if len(validationEligibleProps) > 0 {
				validations[v.Operation] = validationEligibleProps
			}
		}
	}

	return validations
}
