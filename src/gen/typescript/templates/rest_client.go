package templates

import "strings"

var RestClient = strings.TrimPrefix(`
import { APIError, ErrorType } from './definitions';

const API_HOST =
  process.env.NODE_ENV === 'production'
    ? 'https://%s'
    : process.env.API_HOST || 'http://localhost:1000';
const API_BASE_PATH = API_HOST + '%s';

enum HTTP_METHOD {
  GET = 'GET',
  POST = 'POST',
  PUT = 'PUT',
  DELETE = 'DELETE',
}

/** RestClient represents an HTTP client. */
export class RestClient {
  /** The user's JWT. */
	private _token: string;
	/** The default retry count. */
	private readonly _defaultRetryCount = 3;

  constructor(token: string) {
		this._token = token;

		// Verify that we are in a browser environment.
		// This class could later be expanded to support Node.js.
		if (typeof window === 'undefined') {
      throw new RuntimeEnvironmentError();
    }
	}

	/** setToken updates the local value. */
	setToken(t: string): void {
    this._token = t;
  }

	/** get executes a GET request. */
  async get<P = any>(url: string): Promise<any> {
    return await this.do<P>(HTTP_METHOD.GET, url, undefined, this._defaultRetryCount);
  }

	/** post executes a POST request. */
  async post<P = any>(url: string, payload?: P): Promise<any> {
    return await this.do<P>(HTTP_METHOD.POST, url, payload);
  }

	/** put executes a PUT request. */
  async put<P = any>(url: string, payload?: P): Promise<any> {
    return await this.do<P>(HTTP_METHOD.PUT, url, payload);
  }

	/** delete executes a DELETE request. */
  async delete<P = any>(url: string, payload?: P): Promise<any> {
    return await this.do<P>(HTTP_METHOD.DELETE, url, payload, this._defaultRetryCount);
  }

	/** do executes a request. */
  private async do<P = any>(
    method: HTTP_METHOD,
    path: string,
    payload?: P,
	  retries: number = 0,
  ): Promise<any> {
    if (!this._token) throw new MissingTokenError();
    const headers: Record<string, string> = { Authorization: 'Bearer ' + this._token };

    let body = '';
    if (payload) {
      headers['Content-Type'] = 'application/json';
      body = JSON.stringify(payload);
    }

    const resp = await fetch(API_BASE_PATH + path, {
      method,
      headers,
      body,
    });
    if (process.env.NODE_ENV != "production")
		console.debug(method, path, resp.status, { payload, retries });

		const respData = await resp.json();
		if (respData.ok) return respData;
		else {
			const err = respData.error as APIError;
			if (err.error_type === ErrorType.AUTHENTICATION) window.reload()
			if (retries > 0) return this.do<P>(method, path, payload, retries - 1)
			else throw new FetchError(JSON.stringify(err, undefined, 1))
		}
  }
}

/**
 * RuntimeEnvironmentError represents an error that occurred because the runtime environment is not
 * a browser.
 */
export class RuntimeEnvironmentError extends Error {
  constructor() {
    super('RestClient: must be used in a browser environment.');
    this.name = 'RuntimeEnvironmentError';
  }
}

/** FetchError represents an error that occurred during fetching from the API. */
export class FetchError extends Error {
  constructor(msg: string) {
    super('RestClient: ' + msg);
    this.name = 'FetchError';
  }
}

/** MissingTokenError represents an empty JWT. */
export class MissingTokenError extends Error {
  constructor() {
    super('RestClient: cannot execute API request: missing token');
    this.name = 'MissingTokenError';
  }
}`, "\n")
