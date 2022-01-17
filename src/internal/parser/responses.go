package parser

import (
	"regexp"

	"github.com/iancoleman/strcase"
)

var responseRegex = regexp.MustCompile(`([aA-zZ]+[a-z])(Create|Get|List|Update)([aA-zZ]+)?Response`)

var returnsSnapshotTruthMap = map[string]bool{
	"AccommodationsListResponse": true,
	"GroupsListResponse":         true,
}

// parseIntoResponses maps swagger definitions into a new instance of `map[string]*Definition`,
func parseIntoResponses(rawDefs map[string]interface{}) map[string]*Definition {
	respMap := make(map[string]*Definition)
	for k, v := range rawDefs {
		resp := &Definition{
			Key:        strcase.ToCamel(k),
			Properties: make([]*DefinitionProperty, 0),
		}
		if vTyped, ok := v.(Record); ok {
			matches := responseRegex.FindStringSubmatch(resp.Key)
			if len(matches) > 0 {
				switch matches[2] {
				case "Create":
					if t := matches[1]; t != "GroupMembers" && t != "AccommodationResidents" {
						resp.Returns = matches[1]
					}
				case "Get":
					resp.Returns = matches[1] + matches[3]
				case "List":
					withSnapshot := ""
					if val, ok := returnsSnapshotTruthMap[resp.Key]; ok && val {
						withSnapshot = "Snapshot"
					}
					returns := matches[1][:len(matches[1])-1]
					if resp.Key == "MembersListResponse" {
						returns = "Dynamic" + returns // DynamicMember
					}
					resp.Returns = returns + withSnapshot + "[]"
				case "Update":
					if matches[1] != "AccommodationResidents" {
						resp.Returns = "string"
					}
				}
			}
			if headers := vTyped["headers"]; headers != nil {
				if headersTyped, ok := headers.(Record); ok {
					// Keys which aren't part of the extended definition.
					for propKey, propVal := range headersTyped {
						prop := &DefinitionProperty{
							Key: propKey.(string),
						}
						if propValTyped, ok := propVal.(Record); ok {
							if propType := propValTyped["type"]; propType != nil {
								prop.Type = propType.(string)
							}
							if propDesc := propValTyped["description"]; propDesc != nil {
								matches := descriptionRegex.FindStringSubmatch(propDesc.(string))
								if len(matches) >= 2 {
									prop.Description = matches[1]
								}
							}
						}
						if schema := vTyped["schema"]; schema != nil {
							if schemaTyped, ok := schema.(Record); ok {
								if propRef := schemaTyped["$ref"]; propRef != nil {
									prop.Ref = toRef(propRef.(string))
								}
							}
						}
						resp.Properties = append(resp.Properties, prop)
					}
				}
			}
			if schema := vTyped["schema"]; schema != nil {
				if schemaTyped, ok := schema.(Record); ok {
					if propRef := schemaTyped["$ref"]; propRef != nil {
						resp.Ref = toRef(propRef.(string))
					}
				}
			}
		}
		respMap[resp.Key] = resp
	}
	return respMap
}
