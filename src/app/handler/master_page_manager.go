package handler

import (
	"database/sql"
	"go-triton-app/app/cfg"
	"go-triton-app/app/logx"
	"log"
	"net/http"
	"path/filepath"

	"go-triton-app/app/handler/localization"

	"github.com/mgenware/go-packagex/v5/httpx"
	"github.com/mgenware/go-packagex/v5/templatex"
)

// MasterPageManager provides common functions to generate HTML strings.
type MasterPageManager struct {
	dir    string
	logger *logx.Logger
	config *cfg.Config

	reloadViewsOnRefresh bool
	log404Error          bool

	masterView          *LocalizedView
	errorView           *LocalizedView
	LocalizationManager *localization.Manager
}

// MustCreateMasterPageManager creates an instance of MasterPageManager with the specified arguments. Note that this function panics when master template fails to load.
func MustCreateMasterPageManager(
	dir string,
	i18nDir string,
	defaultLang string,
	logger *logx.Logger,
	config *cfg.Config,
) *Manager {
	reloadViewsOnRefresh := config.Debug != nil && config.Debug.ReloadViewsOnRefresh
	if reloadViewsOnRefresh {
		log.Print("⚠️ View dev mode is on")
	}

	// Create the localization manager used by localized template views
	localizationManager, err := localization.NewManagerFromDirectory(i18nDir, defaultLang)
	if err != nil {
		panic(err)
	}

	t := &Manager{
		dir:                  dir,
		LocalizationManager:  localizationManager,
		config:               config,
		logger:               logger,
		reloadViewsOnRefresh: reloadViewsOnRefresh,
		log404Error:          config.HTTP.Log404Error,
	}

	// Load the master template.
	t.masterView = t.MustParseLocalizedView("master.html")
	// Load the error template.
	t.errorView = t.MustParseLocalizedView("error.html")

	return t
}

// MustCompleteWithContent finished the response with the given HTML content.
func (m *Manager) MustCompleteWithContent(content []byte, w http.ResponseWriter) {
	httpx.SetResponseContentType(w, httpx.MIMETypeHTMLUTF8)
	w.Write(content)
}

// MustComplete executes the main view template with the specified data and panics if error occurs.
func (m *Manager) MustComplete(r *http.Request, lang string, d *MasterPageData, w http.ResponseWriter) {
	if d == nil {
		panic("Unexpected empty `MasterPageData` in `MustComplete`")
	}
	httpx.SetResponseContentType(w, httpx.MIMETypeHTMLUTF8)

	// Setup additional assets, e.g.:
	// data.Header += "<link href=\"/static/main.min.css\" rel=\"stylesheet\"/>"
	// data.Scripts += "<script src=\"/static/main.min.js\"></script>"

	// Add site title.
	d.Title = m.PageTitle(lang, d.Title)

	m.masterView.MustExecute(lang, w, d)
}

// MustError executes the main view template with the specified data and panics if error occurs.
func (m *Manager) MustError(r *http.Request, lang string, err error, expected bool, w http.ResponseWriter) {
	d := &ErrorPageData{Message: err.Error()}
	// Handle unexpected errors.
	if !expected {
		if err == sql.ErrNoRows {
			// Consider `sql.ErrNoRows` as 404 not found error.
			w.WriteHeader(http.StatusNotFound)
			// Set `expected` to `true`.
			expected = true

			d.Message = m.Dictionary(lang).ResourceNotFound
			if m.config.HTTP.Log404Error {
				m.logger.NotFound("sql", r.URL.String())
			}
		} else {
			// At this point, this should be a 500 server internal error.
			w.WriteHeader(http.StatusInternalServerError)
			m.logger.Error("fatal-error", "msg", d.Message)
		}
	}
	errorHTML := m.errorView.MustExecuteToString(lang, d)
	htmlData := NewMasterPageData(m.Dictionary(lang).ErrorOccurred, errorHTML)
	m.MustComplete(r, lang, htmlData, w)
}

// PageTitle returns the given string followed by the localized site name.
func (m *Manager) PageTitle(lang, s string) string {
	return s + " - " + m.LocalizationManager.Dictionary(lang).SiteName
}

// MustParseLocalizedView creates a new LocalizedView with the given relative path.
func (m *Manager) MustParseLocalizedView(relativePath string) *LocalizedView {
	file := filepath.Join(m.dir, relativePath)
	view := templatex.MustParseView(file, m.reloadViewsOnRefresh)
	return &LocalizedView{view: view, localizationManager: m.LocalizationManager}
}

// MustParseView creates a new View with the given relative path.
func (m *Manager) MustParseView(relativePath string) *templatex.View {
	file := filepath.Join(m.dir, relativePath)
	return templatex.MustParseView(file, m.reloadViewsOnRefresh)
}

// Dictionary returns a localized dictionary with the specified language ID.
func (m *Manager) Dictionary(lang string) *localization.Dictionary {
	return m.LocalizationManager.Dictionary(lang)
}
