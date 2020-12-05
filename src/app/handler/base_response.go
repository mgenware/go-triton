package handler

import (
	"context"
	"go-triton-app/app/defs"
	"net/http"
)

// BaseResponse provides basic properties shared by both HTMLResponse and JSONResponse.
type BaseResponse struct {
	req  *http.Request
	ctx  context.Context
	lang string
}

func newBaseResponse(r *http.Request) BaseResponse {
	ctx := r.Context()
	c := BaseResponse{
		req:  r,
		ctx:  ctx,
		lang: defs.LanguageContext(ctx),
	}

	return c
}

// Request returns underlying http.Request.
func (b *BaseResponse) Request() *http.Request {
	return b.req
}

// Context returns context.Context associated with current request.
func (b *BaseResponse) Context() context.Context {
	return b.ctx
}

// Lang returns current language ID.
func (b *BaseResponse) Lang() string {
	return b.lang
}
