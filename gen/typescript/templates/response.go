package templates

import "strings"

var Response = strings.TrimPrefix(`
export class %s extends %s {
  constructor(data: any) {
    super(%s);
  }
}`, "\n")

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

var ErrorResponse = strings.TrimPrefix(`
interface %s extends %s {
%s
}`, "\n")
