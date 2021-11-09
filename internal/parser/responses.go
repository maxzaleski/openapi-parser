package parser

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

var entityRegex = regexp.MustCompile(`([aA-zZ]+[a-z])(Create|Get|List|Update)([aA-zZ]+)?Response`)

type Record = map[interface{}]interface{}

// parseIntoResponses maps swagger definitions into a new instance of `map[string]*Definition`,
func parseIntoResponses(rawDefs map[string]interface{}) map[string]*Definition {
	respMap := make(map[string]*Definition)
	for k, v := range rawDefs {
		// This method will be deprecated soon.
		// TODO(MZ): Re: api refactor - groups.memberships/{member_id}
		if k == "memberListGroupsResponse" {
			continue
		}
		resp := &Definition{
			Key:        strcase.ToCamel(k),
			Properties: make([]*DefinitionProperty, 0),
		}
		if vTyped, ok := v.(Record); ok {
			matched := entityRegex.FindStringSubmatch(resp.Key)
			if len(matched) > 0 {
				switch matched[2] {
				case "Create":
					if t := matched[1]; t != "GroupMembers" && t != "AccommodationResidents" {
						resp.Returns = matched[1]
					}
				case "Get":
					resp.Returns = matched[1] + matched[3]
				case "List":
					resp.Returns = matched[1][:len(matched[1])-1] + "[]"
				case "Update":
					if matched[3] == "Whereabouts" {
						resp.Returns = "MemberWhereabouts"
					}
				}
			}
			if headers := vTyped["headers"]; headers != nil {
				if headersTyped, ok := headers.(map[interface{}]interface{}); ok {
					// Keys which aren't part of the extended definition.
					for propKey, propVal := range headersTyped {
						prop := &DefinitionProperty{
							Key: propKey.(string),
						}
						if propValTyped, ok := propVal.(map[interface{}]interface{}); ok {
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
							if schemaTyped, ok := schema.(map[interface{}]interface{}); ok {
								if propRef := schemaTyped["$ref"]; propRef != nil {
									prop.Ref = strings.Replace(propRef.(string), "#/definitions/", "", 1)
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
						resp.Ref = strings.Replace(propRef.(string), "#/definitions/", "", 1)
					}
				}
			}
		}
		respMap[resp.Key] = resp
	}
	return respMap
}

// isExtendedInstance validates whether the response is extended from a parent type.
//
// If it is found to be an extension, will return the definition's key.
// Always returns a slice of unclassified properties.
func isExtendedInstance(key string, rawProps Record) (string, []string) {
	// Map the incoming map's keys into a new slice.
	keys := make([]string, 0, len(rawProps))
	for k := range rawProps {
		keys = append(keys, k.(string))
	}
	// Exit early if the type is in fact `SuccessResponse`.
	if key == "SuccessResponse" {
		return "", keys
	}
	// Keys to be added to the response's properties.
	remainingKeys := make([]string, 0, len(rawProps))

	// Props of `SuccessResponse` which extends `GenericResponse`.
	var (
		hasOK   bool
		hasData bool
	)
	for _, k := range keys {
		switch k {
		case "ok":
			hasOK = true
		case "data":
			hasData = true
		default:
			remainingKeys = append(remainingKeys, k)
		}
	}

	var ref string
	if hasOK && hasData {
		ref = "SuccessResponse"
	} else if hasOK {
		ref = "GenericResponse"
	}
	return ref, remainingKeys
}
