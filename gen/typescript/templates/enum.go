package templates

import "strings"

var Enum = strings.TrimPrefix(`
export enum %s {
%s
}`, "\n")

var EnumStringProperty = "\t%s = '%s',"

var EnumNumberProperty = "\t%s = %s,"
