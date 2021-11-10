package templates

import "strings"

var Response = strings.TrimPrefix(`
export class %s extends %s {
  constructor(data: any) {
    super(%s);
  }
}`, "\n")

var ResponseBody = strings.TrimPrefix(`
export class %s extends %s {
  constructor(data: any) {
    super(data);
  }
}`, "\n")

var ResponseErrorBody = strings.TrimPrefix(`
export class %s extends %s {
	/** The error. */
	readonly error: m.APIError;

  constructor(data: any) {
    super(data);
		this.error = new m.APIError(data.error);
  }
}`, "\n")
