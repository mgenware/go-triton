package template

// ErrorData contains information about an error.
type ErrorData struct {
	Message string
}

// NewErrorData creates an new ErrorData object.
func NewErrorData(msg string) *ErrorData {
	return &ErrorData{Message: msg}
}
