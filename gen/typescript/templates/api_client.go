package templates

const APIClient = `
/** APIClient represents the BoardingHub API interface. */
class APIClient {
	/** The HTTP client. */
	private readonly _client: HTTPClient
	/** The user's JWT'. */
	private _token: string

	constructor() {
		this._client = new HTTPClient('');
		this._token = '';
	}
%s

	/** setToken updates the local value. */
	setToken(value: string): void {
		this._token = value
	}
}

/** BoardingHubAPI represents the API client instance. */
export const BoardingHubAPI = new APIClient();`

const APIClientMethod = `
	async %s(%s): Promise<%s> {
		const path = %s;
	  return await this._client.%s<%s>(%s)
	}`
