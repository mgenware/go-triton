package indexHandler

import (
	"net/http"

	"github.com/mgenware/go-packagex/templatex"

	"github.com/mgenware/go-web-boilerplate/app"
)

var indexView = templatex.MustParseViewFromDirectory(app.Config.TemplatesDir, "index.html")
var indexTitle = "Home Page"

func IndexGET(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := &IndexData{PageName: indexTitle, FeedsHTML: "<p>Hello World</p>"}
	dataHTML := indexView.MustExecuteToString(data)

	md := app.Template.NewMasterData(ctx, indexTitle, dataHTML)
	app.Template.MustRun(md, w)
}
