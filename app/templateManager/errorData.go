package templateManager

type ErrorData struct {
	Message string
}

// NewErrorData creates an instance of ErrorData.
func NewErrorData(msg string) *ErrorData {
	return &ErrorData{Message: msg}
}
