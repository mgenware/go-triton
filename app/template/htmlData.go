package template

type HTMLData struct {
	Title       string
	ContentHTML string
	Header      string
	Scripts     string
}

func NewHTMLData(title, contentHTML string) *HTMLData {
	return &HTMLData{Title: title, ContentHTML: contentHTML}
}
