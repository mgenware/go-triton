package templateManager

import (
	"context"
	"net/http"
)

type HTMLResult struct {
	mgr         *TemplateManager
	writer      http.ResponseWriter
	ctx         context.Context
	isCompleted bool
}

func NewHTMLResult(ctx context.Context, mgr *TemplateManager, wr http.ResponseWriter) *HTMLResult {
	return &HTMLResult{mgr: mgr, writer: wr, ctx: ctx}
}

func (h *HTMLResult) MustComplete(d *HTMLData) {
	if h.isCompleted {
		panic("Result has completed")
	}
	h.isCompleted = true
	h.mgr.MustComplete(h.ctx, d, h.writer)
}

func (h *HTMLResult) MustError(msg string) {
	if h.isCompleted {
		panic("Result has completed")
	}
	h.isCompleted = true
	d := NewErrorData(msg)
	h.mgr.MustError(h.ctx, d, h.writer)
}
