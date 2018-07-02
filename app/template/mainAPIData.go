package template

// MainAPIData contains the information used for composing a JSON response.
type MainAPIData struct {
	// Code denotes the status code of this response.
	Code uint `json:"code"`

	// Message represents an additional string message alongside the response. For example, when `Code` denotes error, it could be the error message, and when `Code` denotes success, it could be a friendly message.
	Message string `json:"msg,omitempty"`

	// Data represents the result of this response.
	Data interface{} `json:"data,omitempty"`
}
