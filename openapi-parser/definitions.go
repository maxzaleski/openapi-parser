package openapi_parser

import "strings"

// Definition represents a type definition.
type Definition struct {
	Key         string
	Description string
	Required    []string
	Properties  []*DefinitionProperty
}

// DefinitionProperty represents a property of `Definition`.
type DefinitionProperty struct {
	// The model's name.
	Key string
	// The model's type.
	Type string
	// The model's description.
	Description string
	// The model's reference key.
	Ref string
	// Whether the property is required.
	Required bool
	// The model's validation properties.
	Validation *DefinitionPropertyValidation
	// the model's format.
	Format string
}

// DefinitionPropertyValidation represents the validation properties of a `DefinitionProperty`.
type DefinitionPropertyValidation struct {
	// The property's regex pattern to match (string).
	Pattern string
	// The property's maximum length (string).
	MaxLength int
	// The property's minimum length (string).
	MinLength int
	// The property's maximum items (slice).
	MaxItems int
	// The property's minimum items (slice).
	MinItems int
}

// parseIntoDefinitions maps swagger definitions into a new instance of `map[string]*Definition`,
func parseIntoDefinitions(rawDefs map[string]interface{}) map[string]*Definition {
	// TODO(MZ): Comment this code.
	// I'm sorry, it's late.
	defsMap := make(map[string]*Definition)
	for k, v := range rawDefs {
		vTyped := v.(map[interface{}]interface{})
		def := &Definition{
			Key: k,
		}
		if val := vTyped["title"]; val != nil {
			def.Description = val.(string)
		}
		required := make(map[string]bool)
		if val := vTyped["required"]; val != nil {
			if valTyped, ok := val.([]interface{}); ok {
				for _, prop := range valTyped {
					required[prop.(string)] = true
				}
			} else {
				required[val.(string)] = true
			}
		}
		if val := vTyped["properties"]; val != nil {
			props := make([]*DefinitionProperty, 0)
			if valTyped, ok := val.(map[interface{}]interface{}); ok {
				for propName, propVal := range valTyped {
					prop := &DefinitionProperty{
						Key:        propName.(string),
						Validation: &DefinitionPropertyValidation{},
					}
					// Checks if this property is marked as 'required'.
					if _, ok := required[propName.(string)]; ok {
						prop.Required = true
					}
					if propValTyped, ok := propVal.(map[interface{}]interface{}); ok {
						// Properties.
						if propType := propValTyped["type"]; propType != nil {
							prop.Type = propType.(string)
						}
						if propDesc := propValTyped["description"]; propDesc != nil {
							prop.Description = propDesc.(string)
						}
						if propRef := propValTyped["$ref"]; propRef != nil {
							prop.Ref = strings.Replace(propRef.(string), "#/definitions/", "", 1)
						}
						if propFormat := propValTyped["format"]; propFormat != nil {
							prop.Format = propFormat.(string)
						}
						// Validation properties.
						if propPattern := propValTyped["pattern"]; propPattern != nil {
							prop.Validation.Pattern = propPattern.(string)
						}
						if propMinLength := propValTyped["minLength"]; propMinLength != nil {
							prop.Validation.MinLength = propMinLength.(int)
						}
						if propMaxLength := propValTyped["maxLength"]; propMaxLength != nil {
							prop.Validation.MaxLength = propMaxLength.(int)
						}
						if propMinItems := propValTyped["minItems"]; propMinItems != nil {
							prop.Validation.MinItems = propMinItems.(int)
						}
						if propMaxItems := propValTyped["maxItems"]; propMaxItems != nil {
							prop.Validation.MaxItems = propMaxItems.(int)
						}
					}
					props = append(props, prop)
				}
			}
			def.Properties = props
		}
		defsMap[k] = def
	}
	return defsMap
}
