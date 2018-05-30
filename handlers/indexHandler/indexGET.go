package indexHandler

import (
	"net/http"

	"github.com/mgenware/go-web-boilerplate/app"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	md := app.Template.NewMasterData(ctx, "Home Page", "<p>Hello World</p>")
	app.Template.MustRun(md, w)
}
