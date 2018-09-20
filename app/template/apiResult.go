package template

// APIResult contains the information about the return value of an API, which is either success or failure.
type APIResult struct {
	// Code indicates the status code of this result. 0 means success.
	Code uint `json:"code"`

	// Message represents an additional string message alongside the result. For example, in a error result, it could be the error message, and when `Code` is 0 (indicating a success result), it could be a friendly message.
	Message string `json:"message,omitempty"`

	// Data represents the requested value in a successful result.
	Data interface{} `json:"data,omitempty"`
}
