package templates

const Request = `
interface %sRequest {
%s
}`

const RequestValidation = `
const %sRequestValidation = yupObject({
%s
})`
