package homep

import (
	"net/http"
	"time"

	"go-triton-app/app"
	"go-triton-app/app/handler"
)

// Home page template.
var homeView = app.MainPageManager.MustParseLocalizedView("home.html")

// Home page GET handler.
func HomeGET(w http.ResponseWriter, r *http.Request) handler.HTML {
	// Create an HTML response.
	resp := app.HTMLResponse(w, r)
	// Prepare home page data.
	pageData := &HomePageData{Time: time.Now().String()}
	// Generate page HTML.
	pageHTML := homeView.MustExecuteToString(resp.Lang(), pageData)
	// Create main page data, which is a core template shared by all your website pages.
	d := app.MainPageData(resp.LocalizedDictionary().Home, pageHTML)
	// Complete the response.
	return resp.MustComplete(d)
}
