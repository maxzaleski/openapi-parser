package internal

// ShortenPathParam returns a shortened version of the given path parameter.
func ShortenPathParam(param string) string {
	switch param {
	case "accommodation_id":
		return "accom_id"
	case "organisation_id":
		return "org_id"
	case "member_id":
		return "mbr_id"
	case "location_id":
		return "loc_id"
	case "address_id":
		return "addr_id"
	case "group_id":
		return "grp_id"
	default:
		return param
	}
}
