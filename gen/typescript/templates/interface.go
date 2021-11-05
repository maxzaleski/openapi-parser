package templates

import "strings"

var Interface = strings.TrimPrefix(`
interface %s {
%s
}`, "\n")
