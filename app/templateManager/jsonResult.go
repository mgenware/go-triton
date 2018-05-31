package templateManager

import (
	"encoding/json"
	"net/http"

	"github.com/mgenware/go-packagex/httpx"
)

type JSONResult struct {
	mgr         *TemplateManager
	writer      http.ResponseWriter
	isCompleted bool
}

func NewJSONResult(mgr *TemplateManager, wr http.ResponseWriter) *JSONResult {
	return &JSONResult{mgr: mgr, writer: wr}
}

func (j *JSONResult) MustError(msg string) {
	d := &JSONData{Message: msg, Code: 1}
	j.mustWriteData(d)
}

func (j *JSONResult) MustErrorWithObject(err error) {
	j.MustError(err.Error())
}

func (j *JSONResult) MustErrorWithCode(code uint, msg string) {
	d := &JSONData{Code: code, Message: msg}
	j.mustWriteData(d)
}

func (j *JSONResult) MustCompleteWithData(data interface{}) {
	d := &JSONData{Data: data}
	j.mustWriteData(d)
}

func (j *JSONResult) MustComplete() {
	d := &JSONData{}
	j.mustWriteData(d)
}

func (j *JSONResult) mustWriteData(d *JSONData) {
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
