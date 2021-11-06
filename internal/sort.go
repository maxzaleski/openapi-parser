package internal

import (
	"reflect"
	"sort"
	"strings"

	"openapi-gen/gen/parser"
)

// SortMapKeysAlphabetically returns an alphabetically-sorted slice of keys from a map.
func SortMapKeysAlphabetically(m interface{}) []string {
	keys := reflect.ValueOf(m).MapKeys()
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].String() < keys[j].String()
	})
	return sliceToString(keys)
}

func sliceToString(values []reflect.Value) []string {
	result := make([]string, 0, len(values))
	for _, v := range values {
		result = append(result, v.String())
	}
	return result
}

// MapByPkg returns the map with keys matching the different core packages.
func MapByPkg(paths map[string]*parser.Path) map[string][]*parser.Path {
	result := make(map[string][]*parser.Path)
	for k, v := range paths {
		key := strings.Split(k, ".")[0]
		if slice, ok := result[key]; ok {
			result[key] = append(slice, v)
		} else {
			result[key] = []*parser.Path{v}
		}
	}
	return result
}

// SortProperties will sort the given properties by:
// - placing entity Identifiers at the top
// - placing timestamps at the bottom
// - otherwise, alphabetically
func SortProperties(props []*parser.DefinitionProperty) []*parser.DefinitionProperty {
	sort.Slice(props, func(i, j int) bool {
		// Always place "id" property first.
		if props[i].Key == "id" {
			return true
		}
		// Always place timestamps last.
		if strings.HasSuffix(props[i].Key, "_at") {
			return false
		}
		// Otherwise, sort alphabetically.
		return props[i].Key < props[j].Key
	})
	return props
}
