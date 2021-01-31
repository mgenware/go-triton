package homep

import (
	"net/http"
	"time"

	"go-triton-app/app"
	"go-triton-app/app/handler"
)

var indexView = app.MainPageManager.MustParseLocalizedView("home.html")

// HomeGET is the HTTP handler for root URL.
func HomeGET(w http.ResponseWriter, r *http.Request) handler.HTML {
	resp := app.HTMLResponse(w, r)
	pageData := &HomePageData{Time: time.Now().String()}
	pageHTML := indexView.MustExecuteToString(resp.Lang(), pageData)

	d := app.MainPageData(resp.LocalizedDictionary().Home, pageHTML)
	return resp.MustComplete(d)
}
