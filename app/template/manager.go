package template

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/mgenware/go-packagex/httpx"
	"github.com/mgenware/go-packagex/templatex"
	"github.com/mgenware/go-triton/app/template/localization"
)

// Manager provides common functions to generate HTML strings.
type Manager struct {
	devMode bool
	dir     string

	masterView          *LocalizedView
	errorView           *LocalizedView
	LocalizationManager *localization.Manager
}

// MustCreateManager creates an instance of TemplateManager with specified arguments. Note that this function panics when main template loading fails.
func MustCreateManager(
	dir string,
	devMode bool,
	i18nDir string,
	defaultLang string,
) *Manager {
	// Set the global devMode (which affects template loading)
	templatex.SetGlobalDevMode(devMode)

	// Create the localization manager used by localized template views
	localizationManager, err := localization.NewManagerFromDirectory(i18nDir, defaultLang)
	if err != nil {
		panic(err)
	}

	t := &Manager{dir: dir, LocalizationManager: localizationManager}

	// Load the master template
	t.masterView = t.MustParseLocalizedView("master.html")
	// Load the error template
	t.errorView = t.MustParseLocalizedView("error.html")

	return t
}

// MustComplete executes the main view template with the specified data and panics if error occurs.
func (m *Manager) MustComplete(ctx context.Context, d *MasterPageData, w http.ResponseWriter) {
	httpx.SetResponseContentType(w, httpx.MIMETypeHTMLUTF8)

	// Setup additional assets, e.g.:
	// data.Header += "<link href=\"/static/main.min.css\" rel=\"stylesheet\"/>"
	// data.Scripts += "<script src=\"/static/main.min.js\"></script>"

	m.masterView.MustExecute(ctx, w, d)
}

// MustError executes the main view template with the specified data and panics if error occurs.
func (m *Manager) MustError(ctx context.Context, d *ErrorPageData, w http.ResponseWriter) {
	errorHTML := m.errorView.MustExecuteToString(ctx, d)
	htmlData := NewMasterPageData("Error", errorHTML)
	m.MustComplete(ctx, htmlData, w)
}

// NewTitle adds a consistent suffix to the specified title.
func (m *Manager) NewTitle(t string) string {
	return t + " - Go-Triton"
}

// NewLocalizedTitle calls NewTitle with a localized title associated with the specified key.
func (m *Manager) NewLocalizedTitle(ctx context.Context, key string) string {
	ls := m.LocalizationManager.ValueForKey(ctx, key)
	return m.NewTitle(ls)
}

// NewHTMLResponse wraps a call to NewHTMLResponse.
func (m *Manager) NewHTMLResponse(ctx context.Context, w http.ResponseWriter) *HTMLResponse {
	return NewHTMLResponse(ctx, m, w)
}

// NewJSONResponse wraps a call to NewJSONResponse.
func (m *Manager) NewJSONResponse(w http.ResponseWriter) *JSONResponse {
	return NewJSONResponse(m, w)
}

// NewMasterPageData wraps a call to MasterPageData.
func (m *Manager) NewMasterPageData(title, contentHTML string) *MasterPageData {
	return NewMasterPageData(title, contentHTML)
}

// MustParseLocalizedView creates a new LocalizedView with the given relative path.
func (m *Manager) MustParseLocalizedView(relativePath string) *LocalizedView {
	file := filepath.Join(m.dir, relativePath)
	view := templatex.MustParseView(file)
	return &LocalizedView{view: view, localizationManager: m.LocalizationManager}
}

// MustParseView creates a new View with the given relative path.
func (m *Manager) MustParseView(relativePath string) *templatex.View {
	file := filepath.Join(m.dir, relativePath)
	return templatex.MustParseView(file)
}
