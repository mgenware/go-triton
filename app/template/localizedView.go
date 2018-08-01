package template

import (
	"context"
	"io"

	"github.com/mgenware/go-triton/app/defs"
	"github.com/mgenware/go-triton/app/template/localization"

	"github.com/mgenware/go-packagex/templatex"
)

type LocalizedView struct {
	localizationManager *localization.Manager
	view                *templatex.View
}

func (v *LocalizedView) MustExecuteToString(ctx context.Context, data ILocalizedTemplateData) string {
	lang := defs.ContextLanguage(ctx)
	return v.view.MustExecuteToString(v.coerceTemplateData(data, lang))
}

func (v *LocalizedView) MustExecute(ctx context.Context, wr io.Writer, data ILocalizedTemplateData) {
	lang := defs.ContextLanguage(ctx)
	v.view.MustExecute(wr, v.coerceTemplateData(data, lang))
}

func (v *LocalizedView) coerceTemplateData(data ILocalizedTemplateData, lang string) interface{} {
	dic := v.localizationManager.DictionaryForLanguage(lang)
	data.SetLS(dic.Map)
	return data
}
