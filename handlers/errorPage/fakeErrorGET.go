package errorPage

import (
	"errors"
	"net/http"

	"github.com/mgenware/go-triton/app"
)

// FakeErrorGET is the GET handler for "/fakeError"
func FakeErrorGET(w http.ResponseWriter, r *http.Request) {
	_, _, resp := app.HTMLResponse(w, r)
	err := errors.New("🐒 This is a demo error page")

	resp.MustErrorWithObject(err)
}
