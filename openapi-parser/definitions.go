package openapi_parser

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

// Definition represents a type definition.
type Definition struct {
	// The model's key.
	Key string
	// The model's type.
	Type string
	// The model's description.
	Description string
	// The model's properties.
	Properties []*DefinitionProperty
	// The model's enum entries.
	EnumEntries []string
}

// DefinitionProperty represents a property of `Definition`.
type DefinitionProperty struct {
	// The property's name.
	Key string
	// The property's type.
	Type string
	// The property's description.
	Description string
	// The property's reference key.
	Ref string
	// Whether the property is required.
	Required bool
	// The property's validation properties.
	Validation *DefinitionPropertyValidation
	// The property's format.
	Format string
}

// enumToMap represents an enumeration to map into a `Definition`.
type enumToMap struct {
	Key     string
	Entries []string
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

var descriptionRegex = regexp.MustCompile(`^(.*\.)`)

// parseIntoDefinitions maps swagger definitions into a new instance of `map[string]*Definition`,
func parseIntoDefinitions(rawDefs map[string]interface{}) map[string]*Definition {
	// TODO(MZ): Comment this code.
	// I'm sorry, it's late.
	defsMap := make(map[string]*Definition)
	enumsToMap := make([]*enumToMap, 0)
	for k, v := range rawDefs {
		def := &Definition{
			Key:  k,
			Type: "object",
		}
		vTyped := v.(map[interface{}]interface{})
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
							matches := descriptionRegex.FindStringSubmatch(propDesc.(string))
							if len(matches) >= 2 {
								prop.Description = matches[1]
							}
						}
						if propRef := propValTyped["$ref"]; propRef != nil {
							// Strip away spec's prefix.
							prop.Ref = strings.Replace(propRef.(string), "#/definitions/", "", 1)
						}
						if propFormat := propValTyped["format"]; propFormat != nil {
							prop.Format = propFormat.(string)
						}
						if propSchema := propValTyped["schema"]; propSchema != nil {
							if propSchemaTyped, ok := propSchema.(map[interface{}]interface{}); ok {
								prop.Type = propSchemaTyped["type"].(string)
							}
						}
						// Enumerations should be extracted into their own definitions,
						// and properties should reference them.
						if propEnum := propValTyped["enum"]; propEnum != nil {
							enumEntries := make([]string, 0)
							if propEnumTyped, ok := propEnum.([]interface{}); ok {
								for _, entry := range propEnumTyped {
									enumEntries = append(enumEntries, entry.(string))
								}
								key := strcase.ToCamel(prop.Key)
								prop.Type = "" // Reset to follow ref.
								prop.Ref = key
								enumsToMap = append(enumsToMap, &enumToMap{
									Key:     key,
									Entries: enumEntries,
								})
							}
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
		// Assign definition.
		defsMap[k] = def

		// Map enumeration definitions.
		for _, enum := range enumsToMap {
			// Check that this definition hasn't already been defined.
			if _, ok := defsMap[enum.Key]; !ok {
				defsMap[enum.Key] = &Definition{
					Key:         enum.Key,
					EnumEntries: enum.Entries,
					Type:        "enum",
				}
			}
		}
	}
	return defsMap
}
