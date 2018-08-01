package homePage

import "github.com/mgenware/go-triton/app/template"

// HomePageData contains the information needed for generating the home page.
type HomePageData struct {
	template.LocalizedTemplateData

	Time string
}
