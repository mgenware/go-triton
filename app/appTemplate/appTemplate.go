package appTemplate

import (
	"path/filepath"

	"github.com/mgenware/go-packagex/templatex"
)

type AppTemplate struct {
	devMode bool
	dir     string

	master *templatex.View
}

func NewAppTemplate(dir string, devMode bool) *AppTemplate {
	// Set the global devMode (which affects template loading)
	templatex.GlobalDevMode = devMode

	t := &AppTemplate{dir: dir}

	// Load the master template
	t.master = templatex.MustParseView(filepath.Join(dir, "master.html"))
	return t
}
