package handler

import (
	"database/sql"
	"go-triton-app/app/cfg"
	"go-triton-app/app/logx"
	"log"
	"net/http"
	"path/filepath"

	"go-triton-app/app/handler/localization"

	"github.com/mgenware/goutil/httpx"
	"github.com/mgenware/goutil/templatex"
)

// MainPageManager provides common functions to generate HTML strings.
type MainPageManager struct {
	dir    string
	logger *logx.Logger
	config *cfg.Config

	reloadViewsOnRefresh bool
	log404Error          bool

	mainView            *LocalizedView
	errorView           *LocalizedView
	LocalizationManager *localization.Manager
}

// MustCreateMainPageManager creates an instance of MainPageManager with the specified arguments. Note that this function panics when main template fails to load.
func MustCreateMainPageManager(
	dir string,
	i18nDir string,
	defaultLang string,
	logger *logx.Logger,
	config *cfg.Config,
) *MainPageManager {
	reloadViewsOnRefresh := config.Debug != nil && config.Debug.ReloadViewsOnRefresh
	if reloadViewsOnRefresh {
		log.Print("⚠️ View dev mode is on")
	}

	// Create the localization manager used by localized template views
	localizationManager, err := localization.NewManagerFromDirectory(i18nDir, defaultLang)
	if err != nil {
		panic(err)
	}

	t := &MainPageManager{
		dir:                  dir,
		LocalizationManager:  localizationManager,
		config:               config,
		logger:               logger,
		reloadViewsOnRefresh: reloadViewsOnRefresh,
		log404Error:          config.HTTP.Log404Error,
	}

	// Load the main template.
	t.mainView = t.MustParseLocalizedView("main.html")
	// Load the error template.
	t.errorView = t.MustParseLocalizedView("error.html")

	return t
}

// MustCompleteWithContent finished the response with the given HTML content.
func (m *MainPageManager) MustCompleteWithContent(content []byte, w http.ResponseWriter) {
	httpx.SetResponseContentType(w, httpx.MIMETypeHTMLUTF8)
	_, err := w.Write(content)
	if err != nil {
		panic(err)
	}
}

// MustComplete executes the main view template with the specified data and panics if error occurs.
func (m *MainPageManager) MustComplete(r *http.Request, lang string, d *MainPageData, w http.ResponseWriter) {
	if d == nil {
		panic("Unexpected empty `MainPageData` in `MustComplete`")
	}
	httpx.SetResponseContentType(w, httpx.MIMETypeHTMLUTF8)

	// Setup additional assets, e.g.:
	// data.Header += "<link href=\"/static/main.min.css\" rel=\"stylesheet\"/>"
	// data.Scripts += "<script src=\"/static/main.min.js\"></script>"

	// Add site title.
	d.Title = m.PageTitle(lang, d.Title)

	m.mainView.MustExecute(lang, w, d)
}

// MustError executes the main view template with the specified data and panics if error occurs.
func (m *MainPageManager) MustError(r *http.Request, lang string, err error, expected bool, w http.ResponseWriter) {
	d := &ErrorPageData{Message: err.Error()}
	// Handle unexpected errors.
	if !expected {
		if err == sql.ErrNoRows {
			// Consider `sql.ErrNoRows` as 404 not found error.
			w.WriteHeader(http.StatusNotFound)
			// Set `expected` to `true`.
			expected = true

			d.Message = m.LocalizedDictionary(lang).ResourceNotFound
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
	htmlData := NewMainPageData(m.LocalizedDictionary(lang).ErrorOccurred, errorHTML)
	m.MustComplete(r, lang, htmlData, w)
}

// PageTitle returns the given string followed by the localized site name.
func (m *MainPageManager) PageTitle(lang, s string) string {
	return s + " - " + m.LocalizationManager.Dictionary(lang).SiteName
}

// MustParseLocalizedView creates a new LocalizedView with the given relative path.
func (m *MainPageManager) MustParseLocalizedView(relativePath string) *LocalizedView {
	file := filepath.Join(m.dir, relativePath)
	view := templatex.MustParseView(file, m.reloadViewsOnRefresh)
	return &LocalizedView{view: view, localizationManager: m.LocalizationManager}
}

// MustParseView creates a new View with the given relative path.
func (m *MainPageManager) MustParseView(relativePath string) *templatex.View {
	file := filepath.Join(m.dir, relativePath)
	return templatex.MustParseView(file, m.reloadViewsOnRefresh)
}

// LocalizedDictionary returns a localized dictionary with the specified language ID.
func (m *MainPageManager) LocalizedDictionary(lang string) *localization.Dictionary {
	return m.LocalizationManager.Dictionary(lang)
}
