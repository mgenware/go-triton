package template

// ErrorPageData contains information about an error.
type ErrorPageData struct {
	LocalizedTemplateData
	Message string
}

// NewErrorPageData creates a new ErrorPageData.
func NewErrorPageData(msg string) *ErrorPageData {
	return &ErrorPageData{Message: msg}
}
