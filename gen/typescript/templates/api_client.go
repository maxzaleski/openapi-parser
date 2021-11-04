package templates

const APIClient = `
/** APIClient represents the BoardingHub API interface. */
class APIClient {
	/** The HTTP client. */
	private readonly _httpClient: HTTPClient
	/** The user's JWT'. */
	private _token: string

	constructor() {
		this._httpClient = new HTTPClient('');
		this._token = '';
	}

%

	/** setToken updates the local value. */
	setToken(value: string): void {
		this._token = value
	}
}

/** BoardingHubAPI represents the API client instance. */
export const BoardingHubAPI = new APIClient();`

const ClientCollectionMethod = `
async %s(%s): Promise<%s> {
	const path = %s;
	const { data } = await this.client.%s<%s>(path%s);%s
}
`