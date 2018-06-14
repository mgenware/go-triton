package index

import (
	"math/rand"
	"net/http"

	"github.com/mgenware/go-triton/app"
)

func RandGET(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp := app.Template.NewHTMLResponse(ctx, w)

	// Generate a random bool (0|1)
	i := rand.Intn(2)
	if i == 0 {
		d := app.Template.NewMainPageData(app.Template.MakeTitle("Random Result"), "<p>🙈</p>")
		resp.MustComplete(d)
	} else {
		resp.MustError("Unlucky!!!")
	}
}
