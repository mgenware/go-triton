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

func (j *JSONResult) MustWriteError(msg string) {
	d := &JSONData{Message: msg, Code: 1}
	j.mustWriteData(d)
}

func (j *JSONResult) MustWriteErrorObject(err error) {
	j.MustWriteError(err.Error())
}

func (j *JSONResult) MustWriteErrorWithCode(code uint, msg string) {
	d := &JSONData{Code: code, Message: msg}
	j.mustWriteData(d)
}

func (j *JSONResult) MustWriteObjectResult(object interface{}) {
	d := &JSONData{Result: object}
	j.mustWriteData(d)
}

func (j *JSONResult) MustWriteResult() {
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
