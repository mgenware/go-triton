package template

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/mgenware/go-packagex/httpx"
	"github.com/mgenware/go-packagex/templatex"
)

type TemplateManager struct {
	devMode bool
	dir     string

	mainView  *templatex.View
	errorView *templatex.View
}

// MustCreateTemplateManager creates an instance of TemplateManager with specified arguments. Note that this function panics when main template loading fails.
func MustCreateTemplateManager(dir string, devMode bool) *TemplateManager {
	// Set the global devMode (which affects template loading)
	templatex.SetGlobalDevMode(devMode)

	t := &TemplateManager{dir: dir}

	// Load the main template
	t.mainView = templatex.MustParseView(filepath.Join(dir, "main.html"))
	// Load the error template
	t.errorView = templatex.MustParseView(filepath.Join(dir, "error.html"))
	return t
}

func (m *TemplateManager) MustComplete(ctx context.Context, d *MainPageData, w http.ResponseWriter) {
	httpx.SetResponseContentType(w, httpx.MIMETypeHTMLUTF8)

	// TODO: Setup assets, e.g.:
	// data.Header += "<link href=\"/static/main.min.css\" rel=\"stylesheet\"/>"
	// data.Scripts += "<script src=\"/static/main.min.js\"></script>"

	err := m.mainView.Execute(w, d)
	if err != nil {
		// This panic will be recovered by panicHandler
		panic(err)
	}
}

func (m *TemplateManager) MustError(ctx context.Context, d *ErrorPageData, w http.ResponseWriter) {
	errorHTML := m.errorView.MustExecuteToString(d)
	htmlData := NewMainPageData("Error", errorHTML)
	m.MustComplete(ctx, htmlData, w)
}

// MakeTitle add a consistent suffix to your title string.
func (m *TemplateManager) MakeTitle(t string) string {
	return t + " - MyWebsite"
}

func (m *TemplateManager) NewHTMLResponse(ctx context.Context, w http.ResponseWriter) *HTMLResponse {
	return NewHTMLResponse(ctx, m, w)
}

func (m *TemplateManager) NewJSONResponse(w http.ResponseWriter) *JSONResponse {
	return NewJSONResponse(m, w)
}

func (m *TemplateManager) NewMainPageData(title, contentHTML string) *MainPageData {
	return &MainPageData{Title: title, ContentHTML: contentHTML}
}
