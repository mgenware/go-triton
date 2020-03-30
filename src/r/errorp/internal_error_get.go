package errorp

import (
	"errors"
	"net/http"
)

// InternalErrorGET triggers a server internal error reponse by `panic` in handler.
func InternalErrorGET(w http.ResponseWriter, r *http.Request) {
	panic(errors.New("🐒 This is a demo error page"))
}
