package system

import (
	"net/http"

	"github.com/mgenware/go-triton/app"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	resp := app.HTMLResponse(w, r)
	msg := resp.FormatLocalizedString("pPageNotFound", r.URL.String())
	resp.MustFailWithMessage(msg)
}
