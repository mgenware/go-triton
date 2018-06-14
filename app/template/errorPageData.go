package template

// ErrorPageData contains information about an error.
type ErrorPageData struct {
	Message string
}

// NewErrorPageDataData creates an new ErrorPageData object.
func NewErrorPageData(msg string) *ErrorPageData {
	return &ErrorPageData{Message: msg}
}
