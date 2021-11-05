package templates

import "strings"

var Response = `
export class %s extends %s {
  constructor(data: any) {
    super(%s);
  }
}`

var ResponseBody = strings.TrimPrefix(`
class %s extends %s {
  constructor(data: any) {
    super(data);
  }
}`, "\n")

var ResponseErrorBody = strings.TrimPrefix(`
class %s extends %s {
	/** The error. */
	readonly error: APIError;

  constructor(data: any) {
    super(data);
		this.error = new APIError(data.error);
  }
}`, "\n")

var ErrorResponse = `
interface %s extends %s {
%s
}`
