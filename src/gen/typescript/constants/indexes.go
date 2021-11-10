package constants

import (
	"strings"
)

var RootIndex = strings.TrimPrefix(`
export * from './definitions'
export * from './api-client';
`, "\n")

var DefinitionsIndex = strings.TrimPrefix(`
export * from './models';
export * from './requests';
export * from './responses';
export * from './validation'
export * from './enums';
`, "\n")
