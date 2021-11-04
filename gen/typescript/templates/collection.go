package templates

const ClientCollection = `
export class %sCollection {
	constructor(private readonly client: HTTPClient) {}

%
}
`

const ClientCollectionMethod = `
async %s(%s): Promise<%s> {
	const path = %s;
	const { data } = await this.client.%s<%s>(path%s);%s
}
`
