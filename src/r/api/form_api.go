package api

import (
	"net/http"

	"go-triton-app/app"
)

func formAPI(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	resp := app.JSONResponse(w, r)
	// Note that you are throwing `err.Error()` not the `err` itself, it's because we think this (form parsing error) is an user error, not a fatal error of our app, so throwing a string instead of an error makes it an user error and will not be logged.
	panic(err.Error())
	resp.MustComplete(r.Form)
}
