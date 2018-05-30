package index

import (
	"math/rand"
	"net/http"

	"github.com/mgenware/go-web-boilerplate/app"
)

func RandGET(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result := app.Template.NewHTMLResult(ctx, w)

	// Generate a random bool (0|1)
	i := rand.Intn(2)
	if i == 0 {
		d := app.Template.NewHTMLData(app.Template.MakeTitle("Random Result"), "<p>🙈</p>")
		result.MustComplete(d)
	} else {
		result.MustError("Unlucky!!!")
	}
}
