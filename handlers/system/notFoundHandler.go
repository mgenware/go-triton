package system

import (
	"net/http"

	"github.com/mgenware/go-triton/app"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tm := app.TemplateManager
	ctx := r.Context()
	msg := tm.FormatLocalizedString(ctx, "pPageNotFound", r.URL.String())

	resp := tm.NewHTMLResponse(ctx, w)
	resp.MustFailWithMessage(msg)
}
