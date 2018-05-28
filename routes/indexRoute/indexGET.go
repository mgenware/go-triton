package indexRoute

import (
	"fmt"
	"net/http"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>go-web-boilerplate</h1><p>It's working</p>")
}
