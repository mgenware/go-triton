package handler

import (
	"io"

	"go-triton-app/app/handler/localization"

	"github.com/mgenware/goutil/templatex"
)

// LocalizedView wraps a templatex.View, providing localization support.
type LocalizedView struct {
	localizationManager *localization.Manager
	view                *templatex.View
}

func (v *LocalizedView) MustExecuteToString(lang string, data ILocalizedTemplateData) string {
	return v.view.MustExecuteToString(v.coerceTemplateData(data, lang))
}

func (v *LocalizedView) MustExecute(lang string, wr io.Writer, data ILocalizedTemplateData) {
	v.view.MustExecute(wr, v.coerceTemplateData(data, lang))
}

func (v *LocalizedView) coerceTemplateData(data ILocalizedTemplateData, lang string) interface{} {
	dict := v.localizationManager.Dictionary(lang)
	data.SetLS(dict)
	return data
}
