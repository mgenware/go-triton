package sysh

import (
	"errors"
	"net/http"

	"go-triton-app/app"

	strf "github.com/mgenware/go-string-format"
)

// NotFoundHandler is a application wide handler for 404 errors.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Set 404 status code
	w.WriteHeader(http.StatusNotFound)
	resp := app.HTMLResponse(w, r)
	msg := strf.Format(resp.Dictionary().PPageNotFound, r.URL.String())

	if app.Config.HTTP.Log404Error {
		app.Logger.NotFound("http", r.URL.String())
	}

	// Note that pass `true` as the `expected` param so that template manager won't treat it as a 500 error.
	app.TemplateManager.MustError(r, resp.Lang(), errors.New(msg), true, w)
}
