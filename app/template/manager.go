package template

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/mgenware/go-packagex/httpx"
	"github.com/mgenware/go-packagex/templatex"
)

// Manager provides common operations on generating HTML output.
type Manager struct {
	devMode bool
	dir     string

	mainView  *templatex.View
	errorView *templatex.View
}

// MustCreateManager creates an instance of TemplateManager with specified arguments. Note that this function panics when main template loading fails.
func MustCreateManager(dir string, devMode bool) *Manager {
	// Set the global devMode (which affects template loading)
	templatex.SetGlobalDevMode(devMode)

	t := &Manager{dir: dir}

	// Load the main template
	t.mainView = templatex.MustParseView(filepath.Join(dir, "master.html"))
	// Load the error template
	t.errorView = templatex.MustParseView(filepath.Join(dir, "error.html"))
	return t
}

// MustComplete executes the main view template with the specified data and panics if error occurs.
func (m *Manager) MustComplete(ctx context.Context, d *MainPageData, w http.ResponseWriter) {
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

// MustError executes the main view template with the specified data and panics if error occurs.
func (m *Manager) MustError(ctx context.Context, d *ErrorPageData, w http.ResponseWriter) {
	errorHTML := m.errorView.MustExecuteToString(d)
	htmlData := NewMainPageData("Error", errorHTML)
	m.MustComplete(ctx, htmlData, w)
}

// MakeTitle add a consistent suffix to your title string.
func (m *Manager) MakeTitle(t string) string {
	return t + " - MyWebsite"
}

// NewHTMLResponse wraps a call to NewHTMLResponse.
func (m *Manager) NewHTMLResponse(ctx context.Context, w http.ResponseWriter) *HTMLResponse {
	return NewHTMLResponse(ctx, m, w)
}

// NewJSONResponse wraps a call to NewJSONResponse.
func (m *Manager) NewJSONResponse(w http.ResponseWriter) *JSONResponse {
	return NewJSONResponse(m, w)
}

// NewMainPageData wraps a call to MainPageData.
func (m *Manager) NewMainPageData(title, contentHTML string) *MainPageData {
	return NewMainPageData(title, contentHTML)
}
