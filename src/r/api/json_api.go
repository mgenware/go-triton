package api

import (
	"go-triton-app/app/defs"
	"go-triton-app/app/handler"
	"net/http"

	"go-triton-app/app"
)

// Handler for a JSON-based POST API.
func jsonAPI(w http.ResponseWriter, r *http.Request) handler.JSON {
	// Create a JSON response.
	resp := app.JSONResponse(w, r)
	// Fetch some data from the request.
	dict := defs.BodyContext(r.Context())
	// Complete the response.
	return resp.MustComplete(dict)
}
