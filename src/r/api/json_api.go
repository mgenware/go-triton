package api

import (
	"go-triton-app/app/defs"
	"go-triton-app/app/template"
	"net/http"

	"go-triton-app/app"
)

func jsonAPI(w http.ResponseWriter, r *http.Request) *template.JSONResponse {
	dict := defs.BodyContext(r.Context())
	return app.JSONResponse(w, r).MustComplete(dict)
}
