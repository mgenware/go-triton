package defs

import (
	"context"
)

// ContextLanguage returns the localization language ID associated with the specified context.
func ContextLanguage(ctx context.Context) string {
	return stringFromContext(ctx, LanguageContextKey)
}

func stringFromContext(ctx context.Context, key ContextKey) string {
	val := ctx.Value(key)
	if val == nil {
		return ""
	}
	return val.(string)
}
