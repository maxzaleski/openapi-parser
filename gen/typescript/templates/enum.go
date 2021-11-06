package templates

import "strings"

var Enum = strings.TrimPrefix(`
enum %s {
%s
}`, "\n")

var EnumStringProperty = "\t%s = '%s',"

var EnumNumberProperty = "\t%s = %s,"
