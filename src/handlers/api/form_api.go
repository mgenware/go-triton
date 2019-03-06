package api

import (
	"net/http"

	"go-triton-app/app"
)

func formAPI(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	resp := app.JSONResponse(w, r)
	if err != nil {
		resp.MustFailWithMessage(err.Error())
	} else {
		resp.MustComplete(r.Form)
	}
}
