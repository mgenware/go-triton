package index

import (
	"net/http"

	"github.com/mgenware/go-packagex/templatex"

	"github.com/mgenware/go-triton/app"
)

var indexView = templatex.MustParseViewFromDirectory(app.Config.TemplatesDir, "index.html")
var indexTitle = "Home Page"

// IndexGET is the GET handler for "/".
func IndexGET(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := app.TemplateManager.NewHTMLResponse(ctx, w)

	pageData := &PageData{PageName: indexTitle, FeedsHTML: "<p>Hello World</p>"}
	pageHTML := indexView.MustExecuteToString(pageData)

	d := app.TemplateManager.NewMainPageData(app.TemplateManager.MakeTitle(indexTitle), pageHTML)
	resp.MustComplete(d)
}
