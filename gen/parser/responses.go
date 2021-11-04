package parser

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

// Response represents an API response.
type Response struct {
	// The response's key.
	Key string
	// The response's description.
	Description string
	// The response's properties.
	Properties []*ResponseProperty
	// The response's extended definition's reference key.
	Ref string
	// The response's returned entity.
	Returns string
}

// ResponseProperty represents a property of `Response`.
type ResponseProperty struct {
	// The property's key.
	Key string
	// The property's type.
	Type string
	// The property's description.
	Description string
	// The property's definition reference key.
	Ref string
}

var entityRegex = regexp.MustCompile(`([aA-zZ]+[a-z])(Create|Get|List)([aA-zZ]+)?Response`)

// parseIntoResponses maps swagger definitions into a new instance of `map[string]*Response`,
func parseIntoResponses(rawDefs map[string]interface{}) map[string]*Response {
	respMap := make(map[string]*Response)
	for k, v := range rawDefs {
		// This method will be deprecated soon.
		// TODO(MZ): Re: api refactor - groups.memberships/${member_id}
		if k == "memberListGroupsResponse" {
			continue
		}
		if k == "membersList" {

		}
		resp := &Response{
			Key: strcase.ToCamel(k),
		}
		if vTyped, ok := v.(map[interface{}]interface{}); ok {
			matched := entityRegex.FindStringSubmatch(resp.Key)
			if len(matched) > 0 {
				switch matched[2] {
				case "Create":
					resp.Returns = matched[1]
				case "Get":
					resp.Returns = matched[1] + matched[3]
				case "List":
					resp.Returns = matched[1][:len(matched[1])-1] + "[]"
				}
			}
			if headers := vTyped["headers"]; headers != nil {
				if headersTyped, ok := headers.(map[interface{}]interface{}); ok {
					// Since responses can only be one of the following:
					//
					// 1. GenericResponse (ok)
					// 2. SuccessResponse (GenericResponse, data)
					// 3. PaginationResponse (SuccessResponse, pagination)
					// 4. ErrorResponse (GenericResponse, error)
					//
					// We map each definition to its parent through `Response.Ref`,
					// thereby marking a reference to be extended.
					respRef, remainingKeys := isExtendedInstance(resp.Key, headersTyped)
					if respRef != "" {
						resp.Ref = respRef
					}
					// Keys which aren't part of the extended definition.
					props := make([]*ResponseProperty, 0, len(remainingKeys))
					for _, propKey := range remainingKeys {
						prop := &ResponseProperty{
							Key: propKey,
						}
						if propVal, ok := headersTyped[propKey]; ok {
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
							props = append(props, prop)
						}
					}
					resp.Properties = props
				}
			}
			if schema := vTyped["schema"]; schema != nil {
				if schemaTyped, ok := schema.(map[interface{}]interface{}); ok {
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
func isExtendedInstance(key string, rawProps map[interface{}]interface{}) (string, []string) {
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
