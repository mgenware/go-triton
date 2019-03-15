package errorp

import (
	"errors"
	"net/http"
)

// FakeErrorGET is the GET handler for "/fakeError"
func FakeErrorGET(w http.ResponseWriter, r *http.Request) {
	panic(errors.New("🐒 This is a demo error page"))
}
