package index

import (
	"math/rand"
	"net/http"

	"github.com/mgenware/go-triton/app"
)

// RandGET is the GET handler for "/rand"
func RandGET(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp := app.TemplateManager.NewHTMLResponse(ctx, w)

	// Generate a random bool (0|1)
	i := rand.Intn(2)
	if i == 0 {
		d := app.TemplateManager.NewMainPageData(app.TemplateManager.MakeTitle("Random Result"), "<p>🙈</p>")
		resp.MustComplete(d)
	} else {
		resp.MustError("🐒 This is a fake error for testing purposes only")
	}
}
