package api

import (
	"go-triton-app/app/defs"
	"net/http"

	"go-triton-app/app"
)

func jsonAPI(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	resp := app.JSONResponse(w, r)
	if err != nil {
		resp.MustFailWithMessage(err.Error())
	} else {
		dict := defs.BodyContext(r.Context())
		resp.MustComplete(dict)
	}
}
