package homep

import "go-triton-app/app/handler"

// HomePageData contains the information needed for generating the home page.
type HomePageData struct {
	handler.LocalizedTemplateData

	Time string
}
