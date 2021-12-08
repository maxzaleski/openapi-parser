package gen

const basePath = "../packages"

// getAppropriateDestination returns the appropriate destination for the given language extension.
func getAppropriateDestination(extn Extension) string {
	switch extn {
	case ".ts":
		return basePath + "/boardinghub-web-sdk/src/api"
	default:
		return ""
	}
}
