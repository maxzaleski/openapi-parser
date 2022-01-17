package constants

import "strings"

var DynamicQueryFilterGeneric = strings.TrimPrefix(`
export interface DynamicQueryFilter<T> {
	// Whether the filter is enabled.
	enabled: boolean;
	// The filter's value.
	value: T;
}
`, "\n")
