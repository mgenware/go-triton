package api

import "github.com/go-chi/chi"

var Router = chi.NewRouter()

func init() {
	Router.Post("/form_api", formAPI)
	Router.Post("/json_api", jsonAPI)
}
