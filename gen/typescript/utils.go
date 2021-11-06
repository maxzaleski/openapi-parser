package typescript

import "fmt"

// toJSDoc returns the given description as a JSDoc comment.
func toJSDoc(indent, desc string) string {
	return indent + "/** " + desc + " */\n"
}

// appendValidationMessageToMethodCall appends the given message to the given call and returns
// the formatted result.
func appendValidationMessageToMethodCall(call, msg string, args ...interface{}) string {
	return fmt.Sprintf(call+", '"+msg+"')", args...)
}
