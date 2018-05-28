package systemRoute

import (
	"fmt"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	msg := "The resource you requested \"" + r.URL.String() + "\" does not exist."

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, msg)
}
