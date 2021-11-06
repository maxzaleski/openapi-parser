package templates

import "strings"

const ObjectProperty = "\treadonly %s%s: %s;"

var Interface = strings.TrimPrefix(`
interface %s {
%s
}`, "\n")

var Class = strings.TrimPrefix(`
export class %s {
%s

	constructor(data: any) {
%s
	}%s
}`, "\n")
