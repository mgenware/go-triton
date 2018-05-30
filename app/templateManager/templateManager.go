package templateManager

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

	masterView *templatex.View
	errorView  *templatex.View
}

// MustCreateTemplateManager creates an instance of TemplateManager with specified arguments. Note that this function panics when master template loading fails.
func MustCreateTemplateManager(dir string, devMode bool) *TemplateManager {
	// Set the global devMode (which affects template loading)
	templatex.SetGlobalDevMode(devMode)

	t := &TemplateManager{dir: dir}

	// Load the master template
	t.masterView = templatex.MustParseView(filepath.Join(dir, "master.html"))
	// Load the error template
	t.errorView = templatex.MustParseView(filepath.Join(dir, "error.html"))
	return t
}

// NewMasterData creates a new instance of HTMLData.
func (m *TemplateManager) NewHTMLData(title, contentHTML string) *HTMLData {
	return NewHTMLData(title, contentHTML)
}

// NewErrorData creates a new instance of ErrorData.
func (m *TemplateManager) NewErrorData(msg string) *ErrorData {
	return NewErrorData(msg)
}

func (m *TemplateManager) MustComplete(ctx context.Context, d *HTMLData, w http.ResponseWriter) {
	httpx.SetResponseContentType(w, httpx.MIMETypeHTMLUTF8)

	// TODO: Setup assets, e.g.:
	// data.Header += "<link href=\"/static/main.min.css\" rel=\"stylesheet\"/>"
	// data.Scripts += "<script src=\"/static/main.min.js\"></script>"

	err := m.masterView.Execute(w, d)
	if err != nil {
		// This panic will be recovered by panicHandler
		panic(err)
	}
}

func (m *TemplateManager) MustError(ctx context.Context, d *ErrorData, w http.ResponseWriter) {
	errorHTML := m.errorView.MustExecuteToString(d)
	htmlData := m.NewHTMLData("Error", errorHTML)
	m.MustComplete(ctx, htmlData, w)
}

// MakeTitle add a consistent suffix to your title string.
func (m *TemplateManager) MakeTitle(t string) string {
	return t + " - MyWebsite"
}

func (m *TemplateManager) NewHTMLResult(ctx context.Context, w http.ResponseWriter) *HTMLResult {
	return NewHTMLResult(ctx, m, w)
}

func (m *TemplateManager) NewJSONResult(w http.ResponseWriter) *JSONResult {
	return NewJSONResult(m, w)
}
