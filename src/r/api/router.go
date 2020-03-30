package api

import "go-triton-app/app/handler"

var Router = handler.NewJSONRouter()

func init() {
	Router.Post("/form_api", formAPI)
	Router.Post("/json_api", jsonAPI)
}
