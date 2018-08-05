package template

import (
	"encoding/json"
	"net/http"

	"github.com/mgenware/go-packagex/httpx"
	"github.com/mgenware/go-triton/app/defs"
)

// JSONResponse helps you create a HTTP response in JSON.
type JSONResponse struct {
	mgr         *Manager
	writer      http.ResponseWriter
	isCompleted bool
}

// NewJSONResponse creates a new JSONResponse.
func NewJSONResponse(mgr *Manager, wr http.ResponseWriter) *JSONResponse {
	return &JSONResponse{mgr: mgr, writer: wr}
}

// MustFailWithCodeAndMessage finishes the response with the specified code and message, and panics if unexpected error happens.
func (j *JSONResponse) MustFailWithCodeAndMessage(code uint, msg string) {
	d := &APIResult{Message: msg, Code: code}
	j.mustWriteData(d)
}

// MustFailWithMessage finishes the response with the specified error message, and panics if unexpected error happens.
func (j *JSONResponse) MustFailWithMessage(msg string) {
	j.MustFailWithCodeAndMessage(defs.APIGenericError, msg)
}

// MustFailWithCode finishes the response with the specified error code, and panics if unexpected error happens.
func (j *JSONResponse) MustFailWithCode(code uint) {
	j.MustFailWithCodeAndMessage(code, "")
}

// MustFail finishes the response with an error object, and panics if unexpected error happens.
func (j *JSONResponse) MustFail(err error) {
	j.MustFailWithMessage(err.Error())
}

// MustComplete finishes the response with the given data, and panics if unexpected error happens.
func (j *JSONResponse) MustComplete(data interface{}) {
	d := &APIResult{Data: data}
	j.mustWriteData(d)
}

func (j *JSONResponse) mustWriteData(d *APIResult) {
	if j.isCompleted {
		panic("Result has completed")
	}
	j.isCompleted = true

	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	httpx.SetResponseContentType(j.writer, httpx.MIMETypeJSONUTF8)
	j.writer.Write(bytes)
}
