package gen

import (
	"reflect"
	"sort"
)

func mapIntoSortedKeys(m interface{}) []string {
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
