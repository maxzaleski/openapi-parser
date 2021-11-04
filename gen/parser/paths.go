package parser

import "strings"

// Path represents an API path.
type Path struct {
	Key        string
	HTTPVerb   string
	Parameters []*DefinitionProperty
	Operation  string
}

// parseIntoPaths maps swagger definitions into a new instance of `map[string]*Path`,
func parseIntoPaths(rawDefs map[string]interface{}) map[string]*Path {
	pathMap := make(map[string]*Path)
	for k, v := range rawDefs {
		path := &Path{
			Key: k,
		}
		if vTyped, ok := v.(map[interface{}]interface{}); ok {
			for verbKey, verbVal := range vTyped {
				path.HTTPVerb = verbKey.(string)
				if verbValTyped, ok := verbVal.(map[interface{}]interface{}); ok {
					if parameters := verbValTyped["parameters"]; parameters != nil {
						if parametersTyped, ok := parameters.([]interface{}); ok {
							params := make([]*DefinitionProperty, 0)
							for _, paramVal := range parametersTyped {
								param := &DefinitionProperty{
									Validation: &DefinitionPropertyValidation{},
								}
								if paramValTyped, ok := paramVal.(map[interface{}]interface{}); ok {
									// Properties.
									if paramName := paramValTyped["name"]; paramName != nil {
										param.Key = paramName.(string)
									}
									if paramDesc := paramValTyped["description"]; paramDesc != nil {
										param.Description = paramDesc.(string)
									}
									if paramIn := paramValTyped["in"]; paramIn != nil {
										param.In = paramIn.(string)
									}
									if paramRequired := paramValTyped["required"]; paramRequired != nil {
										param.Required = paramRequired.(bool)
									}
									if paramType := paramValTyped["type"]; paramType != nil {
										param.Type = paramType.(string)
									}
									if paramFormat := paramValTyped["format"]; paramFormat != nil {
										param.Format = paramFormat.(string)
									}
									if paramSchema := paramValTyped["schema"]; paramSchema != nil {
										if paramSchemaTyped, ok := paramSchema.(map[interface{}]interface{}); ok {
											if paramRef := paramSchemaTyped["$ref"]; paramRef != nil {
												param.Ref = strings.Replace(paramRef.(string), "#/definitions/", "", 1)
											}
											if paramType := paramSchemaTyped["type"]; paramType != nil {
												param.Type = paramType.(string)
											}
											if paramType := paramSchemaTyped["items"]; paramType != nil {
												if paramTypeTyped, ok := paramType.(map[interface{}]interface{}); ok {
													if paramSliceRef := paramTypeTyped["type"]; paramSliceRef != nil {
														param.Ref = paramSliceRef.(string)
													}
													if paramSliceRef := paramTypeTyped["$ref"]; paramSliceRef != nil {
														param.Ref = strings.Replace(paramSliceRef.(string), "#/definitions/", "", 1)
													}
												}
											}
										}
									}
									// Validation properties.
									if paramPattern := paramValTyped["pattern"]; paramPattern != nil {
										param.Validation.Pattern = paramPattern.(string)
									}
									if paramMinLength := paramValTyped["minLength"]; paramMinLength != nil {
										param.Validation.MinLength = paramMinLength.(int)
									}
									if paramMaxLength := paramValTyped["maxLength"]; paramMaxLength != nil {
										param.Validation.MaxLength = paramMaxLength.(int)
									}
									if paramMinItems := paramValTyped["minItems"]; paramMinItems != nil {
										param.Validation.MinItems = paramMinItems.(int)
									}
									if paramMaxItems := paramValTyped["maxItems"]; paramMaxItems != nil {
										param.Validation.MaxItems = paramMaxItems.(int)
									}
								}
								params = append(params, param)
							}
							path.Parameters = params
						}
					}
					if operationID := verbValTyped["operationId"]; operationID != nil {
						path.Operation = operationID.(string)
					}
				}
			}
		}
		pathMap[path.Key] = path
	}
	return pathMap
}
