package defs

// ContextKey is the key type used to fetch values from a context.
type ContextKey string

const (
	LanguageContextKey ContextKey = "lang"
	LanguageCookieKey             = "lang"
	LanguageQueryKey              = "lang"
	LanguageCSString              = "cs"
	LanguageENString              = "en"
	BodyContextKey     ContextKey = "body"
)

const (
	APIGenericError = 1
)
