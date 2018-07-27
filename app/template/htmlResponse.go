package template

import (
	"context"
	"net/http"
)

// HTMLResponse helps you create a HTTP response in HTML with MainPageData.
type HTMLResponse struct {
	mgr         *Manager
	writer      http.ResponseWriter
	ctx         context.Context
	isCompleted bool
}

// NewHTMLResponse creates a new HTMLResponse.
func NewHTMLResponse(ctx context.Context, mgr *Manager, wr http.ResponseWriter) *HTMLResponse {
	return &HTMLResponse{mgr: mgr, writer: wr, ctx: ctx}
}

// MustComplete finishes the response with the given MainPageData, and panics if unexpected error happens.
func (h *HTMLResponse) MustComplete(d *MainPageData) {
	if h.isCompleted {
		panic("Result has completed")
	}
	h.isCompleted = true
	h.mgr.MustComplete(h.ctx, d, h.writer)
}

// MustError finishes the response with an error message, and panics if unexpected error happens.
func (h *HTMLResponse) MustError(msg string) {
	if h.isCompleted {
		panic("Result has completed")
	}
	h.isCompleted = true
	d := NewErrorPageData(msg)
	h.mgr.MustError(h.ctx, d, h.writer)
}
