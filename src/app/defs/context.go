package defs

import (
	"context"
)

// LanguageContext returns the localization language ID associated with the specified context.
func LanguageContext(ctx context.Context) string {
	val := ctx.Value(LanguageContextKey)
	if val == nil {
		return ""
	}
	result, ok := val.(string)
	if ok {
		return result
	}
	return ""
}

// BodyContext returns the localization language ID associated with the specified context.
func BodyContext(ctx context.Context) map[string]interface{} {
	val := ctx.Value(BodyContextKey)
	if val == nil {
		return nil
	}
	result, ok := val.(map[string]interface{})
	if ok {
		return result
	}
	return nil
}
