package internal

import (
	"strings"

	"openapi-gen/internal/parser"
)

// OverrideDefinition overrides the definition of a type.
func OverrideDefinition(def *parser.Definition) *parser.Definition {
	// Override definition.
	switch def.Key {
	case "Role":
		def.Key = "MemberRole"
		def.Description = "MemberRole represents a member role."
	case "EntityType":
		def.Key = "ViewEntityType"
		def.Description = "ViewEntityType represents a view entity type."
	case "ErrorCode":
		def.Description = "ErrorCode represents an error code."
	case "ErrorType":
		def.Description = "ErrorType represents an error type."
	case "RelationshipWithMember":
		def.Description = "RelationshipWithMember represents a host-member relationship."
	case "Colour":
		def.Description = "Colour represents a recognised colour."
	}
	// Override properties.
	for _, prop := range def.Properties {
		overrideDefinitionProperty(def.Key, prop)
	}

	return def
}

// ShakeModelDefinitions shakes definitions by moving those with the "enum" type.
func ShakeModelDefinitions(m map[string]*parser.Definition, out map[string]*parser.Definition) {
	for k, v := range m {
		if v.Type == "enum" {
			out[k] = v
			delete(m, k)
		}
	}
}

// ShakeResponseBodyDefinitions shakes definitions by moving those with keys containing
// "ResponseBody".
func ShakeResponseBodyDefinitions(m map[string]*parser.Definition, out map[string]*parser.Definition) {
	for k, v := range m {
		if strings.Contains(k, "ResponseBody") {
			out[k] = v
			delete(m, k)
		}
	}
}

// overrideDefinitionProperty overrides the definition of a property.
func overrideDefinitionProperty(pKey string, prop *parser.DefinitionProperty) {
	switch {
	case pKey == "ListMembersFilterRole" && prop.Key == "value":
		prop.Ref = "MemberRole"
	case prop.Ref == "Role":
		prop.Ref = "MemberRole"
	case prop.Ref == "EntityType":
		prop.Ref = "ViewEntityType"
	case prop.Ref != "":
		switch prop.Type {
		case "ImageFallback":
			prop.Type = "Colour"
		}
	case prop.Key == "changed_by_self":
		prop.Description = "Whether the whereabouts were updated by the current user."
	case prop.Key == "address":
		switch {
		case strings.Contains(pKey, "Accommodation"):
			prop.Description = "The accommodation's address."
		case strings.Contains(pKey, "Organisation"):
			prop.Description = "The organisation's address."
		case strings.Contains(pKey, "Household"):
			prop.Description = "The household's address."
		}
	case prop.Key == "country_code" && IsStandardType(pKey):
		prop.Key = "country"
		prop.Type = "Country"
		prop.Description = "The entity's country information."
	case strings.HasSuffix(prop.Key, "_at") && !strings.Contains(pKey, "ListMembers"):
		prop.Type = "ExtendedDate"
	}
}
