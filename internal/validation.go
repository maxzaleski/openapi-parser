package internal

const (
	ValidationMessageRequired  = "This field is required."
	ValidationMessageMaxLength = "This field allows a maximum of %[1]d characters."
	ValidationMessageMinLength = "This field requires a minimum of %[1]d characters."
	ValidationMessageMin       = "This field requires a minimum of %[1]d."
	ValidationMessageMax       = "This field allows a maximum of %[1]d."
	ValidationMessageMinItems  = "This field requires a minimum of %[1]d item."
	ValidationMessageMaxItems  = "This field allows a maximum of %[1]d item(s)."
	ValidationMessageEmail     = "This field must be a valid email address."
	ValidationMessageURL       = "This field must be a valid URL."
)

// SetValidationMessageFromPattern returns a validation message for the given pattern.
func SetValidationMessageFromPattern(pattern string) string {
	switch pattern {
	case `^[aA-zZ]+[aA-zZ\\s]+$`:
		return "This field must not contain any numbers or special characters."
	case `^\+?\d+$`:
		return "This field must be a valid phone number under the form +1234567890."
	default:
		return ""
	}
}
