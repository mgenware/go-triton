package api

import "go-triton-app/app/router"

var Router = router.NewJSONRouter()

func init() {
	Router.Post("/form_api", formAPI)
	Router.Post("/json_api", jsonAPI)
}
