package handler

import "go-triton-app/app/handler/localization"

// ILocalizedTemplateData is the base type for all localized page data types.
type ILocalizedTemplateData interface {
	SetLS(value *localization.Dictionary)
}

// LocalizedTemplateData implements ILocalizedTemplateData.
type LocalizedTemplateData struct {
	LS *localization.Dictionary
}

func (td *LocalizedTemplateData) SetLS(dict *localization.Dictionary) {
	td.LS = dict
}
