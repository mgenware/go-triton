package errorPage

import (
	"net/http"

	"github.com/mgenware/go-triton/app"
)

// FakeErrorGET is the GET handler for "/fakeError"
func FakeErrorGET(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := app.TemplateManager.NewHTMLResponse(ctx, w)

	resp.MustError("🐒 This is a demo error page!")
}
