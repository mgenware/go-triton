package template

import (
	"context"
	"go-triton-app/app/defs"
	"net/http"
)

// BaseResponse provides basic properties shared by both HTMLResponse and JSONResponse.
type BaseResponse struct {
	req  *http.Request
	ctx  context.Context
	mgr  *Manager
	lang string
}

func newBaseResponse(r *http.Request, mgr *Manager) BaseResponse {
	ctx := r.Context()
	c := BaseResponse{
		req:  r,
		ctx:  ctx,
		lang: defs.LanguageContext(ctx),
		mgr:  mgr,
	}

	return c
}

func (b *BaseResponse) Request() *http.Request {
	return b.req
}

func (b *BaseResponse) Context() context.Context {
	return b.ctx
}

func (b *BaseResponse) Lang() string {
	return b.lang
}

// LocalizedString calls TemplateManager.LocalizedString.
func (b *BaseResponse) LocalizedString(key string) string {
	return b.mgr.LocalizedString(b.lang, key)
}

// FormatLocalizedString calls TemplateManager.FormatLocalizedString.
func (b *BaseResponse) FormatLocalizedString(key string, a ...interface{}) string {
	return b.mgr.FormatLocalizedString(b.lang, key, a...)
}

// PageTitle calls TemplateManager.PageTitle.
func (b *BaseResponse) PageTitle(s string) string {
	return b.mgr.PageTitle(b.lang, s)
}

// LocalizedPageTitle calls TemplateManager.LocalizedPageTitle.
func (b *BaseResponse) LocalizedPageTitle(key string) string {
	return b.mgr.LocalizedPageTitle(b.lang, key)
}

// SetStatus sets the underlying HTTP status code.
func (b *BaseResponse) SetStatus(key string) string {
	return b.mgr.LocalizedPageTitle(b.lang, key)
}
