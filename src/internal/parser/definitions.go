package parser

import (
	"github.com/iancoleman/strcase"
	"regexp"
	"strconv"
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
	// The model's enum entries (enum only).
	EnumEntries []string
	// The model's returned entity (response only).
	Returns string
	// The model's extended definition's reference key.
	Ref string
	// Whether the model is a dynamic query request.
	DynamicQuery *DynamicQuery
}

// DefinitionProperty represents a property of `Definition`.
type DefinitionProperty struct {
	// The property's name.
	Key string
	// The property's type.
	Type string
	// The property's description.
	Description string
	// The property's definition reference key.
	Ref string
	// Whether the property is required.
	Required bool
	// The property's validation properties.
	Validation *DefinitionPropertyValidation
	// The property's format.
	Format string
	// The parameter's destination.
	In string
}

// DynamicQuery represents a dynamic query request.
type DynamicQuery struct {
	// Whether the model is a dynamic query request.
	OK bool
	// The query's characteristics' definition keys.
	CharacteristicKeys []string
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
	// The property's maximum (integer).
	Max int
	// The property's minimum (integer).
	Min int
	// The property's maximum items (slice).
	MaxItems int
	// The property's minimum items (slice).
	MinItems int
}

var dynamicQueryTruthMath = map[string]bool{
	"ListMembersRequestBody": true,
}

var (
	requestRegex = regexp.MustCompile(`((Create|Get|List|Update)[aA-zZ]+)?Request[Body]?`)
)

// parseIntoDefinitions maps swagger definitions into a new instance of `map[string]*Definition`,
func parseIntoDefinitions(rawDefs map[string]interface{}) map[string]*Definition {
	defMap := make(map[string]*Definition)
	enumsToMap := make([]*enumToMap, 0)
	for k, v := range rawDefs {
		def := &Definition{
			Key:  k,
			Type: "object",
		}
		vTyped := v.(Record)
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
		if properties := vTyped["properties"]; properties != nil {
			props := make([]*DefinitionProperty, 0)
			if valTyped, ok := properties.(Record); ok {
				for propName, propVal := range valTyped {
					prop := &DefinitionProperty{
						Key:        propName.(string),
						Validation: &DefinitionPropertyValidation{},
					}
					// Checks if this property is marked as 'required'.
					if _, ok := required[propName.(string)]; ok {
						prop.Required = true
					}
					if propValTyped, ok := propVal.(Record); ok {
						// Properties.
						if propType := propValTyped["type"]; propType != nil {
							prop.Type = propType.(string)
						}
						if propDesc := propValTyped["description"]; propDesc != nil {
							prop.Description = extractDescription(propDesc.(string))
						}
						if propRef := propValTyped["$ref"]; propRef != nil {
							// Strip away spec's prefix.
							prop.Ref = toRef(propRef.(string))
							// Account for dynamic query characteristics.
							// Note: ref will always be set with a dynamic query.
							if dq := def.DynamicQuery; dq != nil && dq.OK {
								dq.CharacteristicKeys = append(dq.CharacteristicKeys, prop.Ref)
							}
						}
						if propFormat := propValTyped["format"]; propFormat != nil {
							prop.Format = propFormat.(string)
						}
						if propSchema := propValTyped["schema"]; propSchema != nil {
							if propSchemaTyped, ok := propSchema.(Record); ok {
								prop.Type = propSchemaTyped["type"].(string)
							}
						}
						// Slices will have their own reference.
						if propItems := propValTyped["items"]; propItems != nil {
							if propItemsTyped, ok := propItems.(Record); ok {
								if propItemsRef := propItemsTyped["$ref"]; propItemsRef != nil {
									prop.Ref = toRef(propItemsRef.(string))
								}
								if propItemsType := propItemsTyped["type"]; propItemsType != nil {
									prop.Ref = toRef(propItemsType.(string))
								}
							}
						}
						// Enumerations should be extracted into their own definitions,
						// and properties should reference them.
						if propEnum := propValTyped["enum"]; propEnum != nil {
							enumEntries := make([]string, 0)
							if propEnumTyped, ok := propEnum.([]interface{}); ok {
								for _, entry := range propEnumTyped {
									entryTyped, ok := entry.(string)
									if !ok {
										entryTyped = strconv.Itoa(entry.(int))
									}
									enumEntries = append(enumEntries, entryTyped)
								}

								key := strcase.ToCamel(prop.Key)
								// Keys equal to "image_fallback" and "colour" will produce the same enum,
								// but we are only interested in the misc.Colour enum.
								if key != "ImageFallbackColourIdx" {
									prop.Type = "" // Reset to follow ref.
									prop.Ref = key
									enumsToMap = append(enumsToMap, &enumToMap{
										Key:     key,
										Entries: enumEntries,
									})
								}
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
						if propMax := propValTyped["maximum"]; propMax != nil {
							prop.Validation.Max = propMax.(int)
						}
						if propMin := propValTyped["minimum"]; propMin != nil {
							prop.Validation.Min = propMin.(int)
						}
					}
					props = append(props, prop)
				}
			}
			def.Properties = props
		}
		// Assign definition.
		defMap[k] = def

		// Map enumeration definitions.
		for _, enum := range enumsToMap {
			// Check that this definition hasn't already been defined.
			if _, ok := defMap[enum.Key]; !ok {
				defMap[enum.Key] = &Definition{
					Key:         enum.Key,
					EnumEntries: enum.Entries,
					Type:        "enum",
				}
			}
		}
	}
	return defMap
}
