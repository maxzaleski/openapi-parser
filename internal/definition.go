package internal

import (
	"strings"

	"openapi-gen/gen/parser"
)

// OverrideDefinition overrides the definition of a type.
func OverrideDefinition(def *parser.Definition) *parser.Definition {
	// Override definition.
	switch def.Key {
	case "Role":
		def.Key = "MemberRole"
		def.Description = "Role represents a member role."
	case "EntityType":
		def.Key = "ViewEntityType"
		def.Description = "ViewEntityType represents a view entity type."
	case "ErrorCode":
		def.Description = "ErrorCode represents an error code."
	case "ErrorType":
		def.Description = "ErrorType represents an error type."
	case "RelationshipWithMember":
		def.Description = "RelationshipWithMember represents a host-member relationship."
	}
	// Override properties.
	for _, prop := range def.Properties {
		overrideDefinitionProperty(def.Key, prop)
	}

	return def
}

// OverrideResponseDefinition overrides the definition of a response.
func OverrideResponseDefinition(def *parser.Definition) *parser.Definition {
	switch def.Key {
	case "RegisterOrganisationResponse":
		def.Returns = def.Ref + "Data"
	}

	return def
}

func overrideDefinitionProperty(pKey string, prop *parser.DefinitionProperty) {
	switch {
	case prop.Ref == "Role":
		prop.Ref = "MemberRole"
	case prop.Ref == "EntityType":
		prop.Ref = "ViewEntityType"
	case prop.Ref != "":
		switch prop.Type {
		case "ImageFallback":
			prop.Type = "Colour"
		}
	case prop.Key == "country_code" && IsStandardType(pKey):
		prop.Key = "country"
		prop.Type = "Country"
		prop.Description = "The entity's country information."
	case strings.HasSuffix(prop.Key, "_at") || strings.Contains(pKey, "ListMembers"):
		prop.Type = "ExtendedDate"
	}
}
