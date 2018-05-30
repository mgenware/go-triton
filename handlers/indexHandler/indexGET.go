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
	result := app.Template.NewHTMLResult(ctx, w)

	indexData := &IndexData{PageName: indexTitle, FeedsHTML: "<p>Hello World</p>"}
	indexDataHTML := indexView.MustExecuteToString(indexData)

	d := app.Template.NewHTMLData(app.Template.MakeTitle(indexTitle), indexDataHTML)
	result.MustComplete(d)
}
