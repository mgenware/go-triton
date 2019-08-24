package defs

import (
	"context"
)

// LanguageContext returns the localization language ID associated with the specified context.
func LanguageContext(ctx context.Context) string {
	result, _ := ctx.Value(LanguageContextKey).(string)
	return result
}

// BodyContext returns the request payload (e.g. parsed JSON contents) associated with the specified context.
func BodyContext(ctx context.Context) map[string]interface{} {
	result, _ := ctx.Value(BodyContextKey).(map[string]interface{})
	return result
}
