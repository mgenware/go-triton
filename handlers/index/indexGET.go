package index

import (
	"net/http"

	"github.com/mgenware/go-packagex/templatex"

	"github.com/mgenware/go-triton/app"
)

var indexView = templatex.MustParseViewFromDirectory(app.Config.TemplatesDir, "index.html")
var indexTitle = "Home Page"

func IndexGET(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := app.Template.NewHTMLResponse(ctx, w)

	indexData := &IndexData{PageName: indexTitle, FeedsHTML: "<p>Hello World</p>"}
	indexDataHTML := indexView.MustExecuteToString(indexData)

	d := app.Template.NewMainPageData(app.Template.MakeTitle(indexTitle), indexDataHTML)
	resp.MustComplete(d)
}
