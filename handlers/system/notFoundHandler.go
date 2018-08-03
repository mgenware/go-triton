package system

import (
	"fmt"
	"net/http"

	"github.com/mgenware/go-triton/app"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tm := app.TemplateManager
	ctx := r.Context()
	msg := fmt.Sprintf(tm.LocalizationManager.ValueForKey(ctx, "pPageNotFound"), r.URL.String())

	resp := tm.NewHTMLResponse(ctx, w)
	resp.MustFailWithMessage(msg)
}
