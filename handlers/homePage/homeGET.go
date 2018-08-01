package homePage

import (
	"net/http"
	"time"

	"github.com/mgenware/go-triton/app"
)

var indexView = app.TemplateManager.MustParseLocalizedView("home.html")

// HomeGET is the GET handler for "/".
func HomeGET(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tm := app.TemplateManager

	resp := tm.NewHTMLResponse(ctx, w)

	pageData := &HomePageData{Time: time.Now().String()}
	pageHTML := indexView.MustExecuteToString(ctx, pageData)

	d := tm.NewMasterPageData(tm.MakeTitle(tm.LocalizationManager.ValueForKey(ctx, "home")), pageHTML)
	resp.MustComplete(d)
}
