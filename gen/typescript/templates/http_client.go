package templates

const HTTPClient = `
const API_HOST =
  process.env.NODE_ENV === 'production'
    ? 'https://%s'
    : process.env.API_HOST || 'http://%s';
const API_BASE_PATH = API_HOST + '%s';

enum HTTP_METHOD {
  GET = 'GET',
  POST = 'POST',
  PUT = 'PUT',
  DELETE = 'DELETE',
}

/** HTTPClient represents an HTTP client. */
class HTTPClient {
  /** The user's JWT. */
	private _token: string;

  constructor(token: string) {
		this._token = token;
	}

	/** setToken updates the local value. */
  setToken(t: string): void {
    this._token = t;
  }

	/** get executes a GET request. */
  async get<P = any, R = any>(url: string, payload?: P): Promise<R> {
    return this.do<P, R>(HTTP_METHOD.GET, url, payload);
  }

	/** post executes a POST request. */
  async post<P = any, R = any>(url: string, payload?: P): Promise<R> {
    return this.do<P, R>(HTTP_METHOD.POST, url, payload);
  }

	/** put executes a PUT request. */
  async put<P = any, R = any>(url: string, payload?: P): Promise<R> {
    return this.do<P, R>(HTTP_METHOD.PUT, url, payload);
  }

	/** delete executes a DELETE request. */
  async delete<P = any, R = any>(url: string, payload?: P): Promise<R> {
    return this.do<P, R>(HTTP_METHOD.DELETE, url, payload);
  }

	/** do executes a request. */
  private async do<P = any, R = any>(
    method: HTTP_METHOD,
    path: string,
    payload?: P
  ): Promise<R> {
    if (!this._token) throw new APITokenError();
    const headers: Record<string, string> = { Authorization: this._token };

    let body = undefined;
    if (payload) {
      headers['Content-Type'] = 'application/json';
      try {
        body = JSON.stringify(payload);
      } catch (err) {
        throw 'failed to stringify request payload :' + err;
      }
    }

    const url = API_BASE_PATH + path;
    const resp = await fetch(url, {
      method,
      headers,
      body,
    });
		if (process.env.NODE_ENV != "production")
			console.log(method, path, resp.status, 'payload:', payload)

		const respData = await resp.json();
		if (respData.ok) return (respData.data || {}) as R
		else {
		  const err = respData.error as APIError;
			if (err.error_type == ErrorType.AUTHENTICATION) window.reload()
			else throw new APIFetchError(err.message)
		}
  }
}

/** APIFetchError represents an error that occurred during fetching from the API. */
export class APIFetchError extends Error {
  constructor(msg: string) {
    super('api-client: ' + msg);
    this.name = 'APIFetchError';
  }
}

/** APITokenError represents an empty JWT. */
export class APITokenError extends Error {
  constructor() {
    super('api-client: cannot execute API request: missing token');
    this.name = 'APITokenError';
  }
}`
