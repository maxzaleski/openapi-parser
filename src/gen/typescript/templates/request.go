package templates

import "strings"

var Request = strings.Trim(`
export interface %sRequest {
%s
}`, "\n")

var RequestValidation = strings.Trim(`
const %sRequestValidation = yupObject({
%s
})`, "\n")
