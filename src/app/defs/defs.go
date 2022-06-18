package defs

// ContextKey is the key type used to fetch values from a context.
type ContextKey string

// Predefined context keys.
const (
	LanguageContextKey ContextKey = "lang"
	BodyContextKey     ContextKey = "body"
)

const (
	LanguageCookieKey = "lang"
	LanguageQueryKey  = "lang"
	// Language code for English.
	LanguageENString = "en"
	// Language code for Chinese simplified.
	LanguageCSString = "zh-Hans"
)

const (
	APIGenericError = 1
)
