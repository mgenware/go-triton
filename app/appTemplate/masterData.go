package appTemplate

type MasterData struct {
	Title       string
	ContentHTML string
	Header      string
	Scripts     string
}

func NewMasterData(title, contentHTML string) *MasterData {
	return &MasterData{Title: title, ContentHTML: contentHTML}
}
