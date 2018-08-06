package template

import (
	"context"
	"net/http"
)

// HTMLResponse helps you create a HTTP response in HTML with MasterPageData.
type HTMLResponse struct {
	BaseResponse

	writer      http.ResponseWriter
	isCompleted bool
}

// NewHTMLResponse creates a new HTMLResponse.
func NewHTMLResponse(ctx context.Context, mgr *Manager, wr http.ResponseWriter) *HTMLResponse {
	return &HTMLResponse{
		BaseResponse: newBaseResponse(ctx, mgr),
		writer:       wr,
	}
}

// MustComplete finishes the response with the given MasterPageData, and panics if unexpected error happens.
func (h *HTMLResponse) MustComplete(d *MasterPageData) {
	if h.isCompleted {
		panic("Result has completed")
	}
	h.isCompleted = true
	h.mgr.MustComplete(h.lang, d, h.writer)
}

// MustFailWithMessage finishes the response with an error message, and panics if unexpected error happens.
func (h *HTMLResponse) MustFailWithMessage(msg string) {
	if h.isCompleted {
		panic("Result has completed")
	}
	h.isCompleted = true
	d := NewErrorPageData(msg)
	h.mgr.MustError(h.lang, d, h.writer)
}

// MustFail calls MustFailWithMessage with the given error object.
func (h *HTMLResponse) MustFail(err error) {
	h.MustFailWithMessage(err.Error())
}
