package template

import (
	"encoding/json"
	"net/http"

	"github.com/mgenware/go-packagex/httpx"
)

// JSONResponse helps you create a HTTP response in JSON with MainAPIData.
type JSONResponse struct {
	mgr         *TemplateManager
	writer      http.ResponseWriter
	isCompleted bool
}

// NewJSONResponse creates a new JSONResponse.
func NewJSONResponse(mgr *TemplateManager, wr http.ResponseWriter) *JSONResponse {
	return &JSONResponse{mgr: mgr, writer: wr}
}

// MustError finish the response with an error message, and panics if unexpected error happens.
func (j *JSONResponse) MustError(msg string) {
	d := &MainAPIData{Message: msg, Code: 1}
	j.mustWriteData(d)
}

// MustErrorWithObject finish the response with an error object, and panics if unexpected error happens.
func (j *JSONResponse) MustErrorWithObject(err error) {
	j.MustError(err.Error())
}

// MustErrorWithCode finish the response with an error code and a message, and panics if unexpected error happens.
func (j *JSONResponse) MustErrorWithCode(code uint, msg string) {
	d := &MainAPIData{Code: code, Message: msg}
	j.mustWriteData(d)
}

// MustCompleteWithData finish the response with the given data, and panics if unexpected error happens.
func (j *JSONResponse) MustCompleteWithData(data interface{}) {
	d := &MainAPIData{Data: data}
	j.mustWriteData(d)
}

// MustComplete finish the response with empty data, and panics if unexpected error happens.
func (j *JSONResponse) MustComplete() {
	d := &MainAPIData{}
	j.mustWriteData(d)
}

func (j *JSONResponse) mustWriteData(d *MainAPIData) {
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
