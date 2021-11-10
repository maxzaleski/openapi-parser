package templates

import "strings"

var RestClient = strings.TrimPrefix(`
import { APIError, ErrorType } from './definitions';

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

/** RestClient represents an HTTP client. */
export class RestClient {
  /** The user's JWT. */
	private _token: string;
	/** The default retry count. */
	private readonly _defaultRetryCount: number = 3;

  constructor(token: string) {
		this._token = token;
	}

	/** setToken updates the local value. */
  setToken(t: string): void {
    this._token = t;

		// Verify that we are in a browser environment.
		// This class could later be expanded to support Node.js.
		if (typeof window === 'undefined') {
      throw new Error('RestClient must be used in a browser environment.');
    }
  }

	/** get executes a GET request. */
  async get<P = any, R = any>(url: string, payload?: P): Promise<R> {
    return this.do<P, R>(HTTP_METHOD.GET, url, payload, this._defaultRetryCount);
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
    return this.do<P, R>(HTTP_METHOD.DELETE, url, payload, this._defaultRetryCount);
  }

	/** do executes a request. */
  private async do<P = any, R = any>(
    method: HTTP_METHOD,
    path: string,
    payload?: P,
	  retries: number = 0,
  ): Promise<R> {
    if (!this._token) throw new MissingTokenError();
    const headers: Record<string, string> = { Authorization: 'Bearer ' + this._token };

    let body: string;
    if (payload && payload.toString() === '[object Object]') {
      headers['Content-Type'] = 'application/json';
      body = JSON.stringify(payload);
    }

    const resp = await fetch(API_BASE_PATH + path, {
      method,
      headers,
      body,
    });
		if (process.env.NODE_ENV != "production")
			console.debug(method, path, resp.status, {
				payload,
				retries,
			});

		const respData = await resp.json();
		if (respData.ok) return respData as R;
		else {
		  const err = respData.error as APIError;
			if (err.error_type === ErrorType.AUTHENTICATION) window.reload()
			if (retries > 0) return this.do<P, R>(method, path, payload, retries - 1)
      else throw new FetchError(err.message)
		}
  }
}

/** FetchError represents an error that occurred during fetching from the API. */
export class FetchError extends Error {
  constructor(msg: string) {
    super('rest-client: ' + msg);
    this.name = 'FetchError';
  }
}

/** MissingTokenError represents an empty JWT. */
export class MissingTokenError extends Error {
  constructor() {
    super('rest-client: cannot execute API request: missing token');
    this.name = 'MissingTokenError';
  }
}`, "\n")
