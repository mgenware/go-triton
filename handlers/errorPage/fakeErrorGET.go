package errorPage

import (
	"net/http"

	"github.com/mgenware/go-triton/app"
)

// FakeErrorGET is the GET handler for "/fakeError"
func FakeErrorGET(w http.ResponseWriter, r *http.Request) {
	_, _, resp := app.HTMLResponse(w, r)

	resp.MustError("🐒 This is a demo error page!")
}
