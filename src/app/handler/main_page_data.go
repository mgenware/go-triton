package handler

// MainPageData holds the data needed in main page template.
type MainPageData struct {
	LocalizedTemplateData

	Title       string
	ContentHTML string
	Header      string
	Scripts     string
}

// NewMainPageData creates a new MainPageData.
func NewMainPageData(title, contentHTML string) *MainPageData {
	return &MainPageData{Title: title, ContentHTML: contentHTML}
}
