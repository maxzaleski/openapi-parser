package constants

import "strings"

var GenericResponse = strings.TrimPrefix(`
/** GenericResponse represents a generic response. */
class GenericResponse {
	/** Whether the request was successful. */
	readonly ok: boolean;

  constructor(data: any) {
		this.ok = data.ok;
	}
}`, "\n")

var SuccessResponse = strings.TrimPrefix(`
/** SuccessResponse represents a success response. */
class SuccessResponse<T> extends GenericResponse {
	/** The response data. */
	readonly data: T;

	constructor(data: any) {
    super(data);
    this.data = data.data as T;
  }
}`, "\n")

var ConstructorSuperRegisterOrganisation = strings.TrimPrefix(`{
			...data,
			data: {
				member_snapshot: new MemberSnapshot(data.data.member_snapshot),
				organisation_snapshot: new OrganisationSnapshot(data.data.organisation_snapshot),
			},
		}`, "\n")
