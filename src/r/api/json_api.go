package api

import (
	"go-triton-app/app/defs"
	"go-triton-app/app/handler"
	"net/http"

	"go-triton-app/app"
)

func jsonAPI(w http.ResponseWriter, r *http.Request) handler.JSON {
	dict := defs.BodyContext(r.Context())
	return app.JSONResponse(w, r).MustComplete(dict)
}
