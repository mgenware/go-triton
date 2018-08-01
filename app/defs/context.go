package defs

import (
	"context"
)

// GetContextLanguage returns the localization language ID associated with the specified context.
func GetContextLanguage(ctx context.Context) string {
	return stringFromContext(ctx, LanguageKey)
}

func stringFromContext(ctx context.Context, key string) string {
	val := ctx.Value(key)
	if val == nil {
		return ""
	}
	return val.(string)
}
