package api

import (
	"net/http"

	"go-triton-app/app"
)

func formAPI(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// Note that we are throwing `err.Error()` not the `err` itself, it's because we think
		// this (form parsing error) is an user error, not a fatal error of our app, so throwing
		// a string instead of an error makes it an user error, which doesn't need to be logged.
		panic(err.Error())
	}
	resp := app.JSONResponse(w, r)
	resp.MustComplete(r.Form)
}
