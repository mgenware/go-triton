package api

import (
	"go-triton-app/app/defs"
	"net/http"

	"go-triton-app/app"
)

func jsonAPI(w http.ResponseWriter, r *http.Request) {
	resp := app.JSONResponse(w, r)
	dict := defs.BodyContext(r.Context())
	resp.MustComplete(dict)
}
