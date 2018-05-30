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
}

// MustCreateTemplateManager creates an instance of TemplateManager with specified arguments. Note that this function panics when master template loading fails.
func MustCreateTemplateManager(dir string, devMode bool) *TemplateManager {
	// Set the global devMode (which affects template loading)
	templatex.SetGlobalDevMode(devMode)

	t := &TemplateManager{dir: dir}

	// Load the master template
	t.masterView = templatex.MustParseView(filepath.Join(dir, "master.html"))
	return t
}

// NewMasterData creates a new MasterData
func (m *TemplateManager) NewMasterData(ctx context.Context, title, contentHTML string) *MasterData {
	data := NewMasterData(title, contentHTML)

	// TODO: retrieve user info from context

	return data
}

func (m *TemplateManager) MustRun(data *MasterData, w http.ResponseWriter) {
	httpx.SetResponseContentType(w, httpx.MIMETypeHTMLUTF8)

	// TODO: Setup assets, e.g.:
	// data.Header += "<link href=\"/static/main.min.css\" rel=\"stylesheet\"/>"
	// data.Scripts += "<script src=\"/static/main.min.js\"></script>"

	err := m.masterView.Execute(w, data)
	if err != nil {
		// This panic will be recovered by panicHandler
		panic(err)
	}
}

// MakeTitle add a consistent suffix to your title string.
func (m *TemplateManager) MakeTitle(t string) string {
	return t + " - MyWebsite"
}
