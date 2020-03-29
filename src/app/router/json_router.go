package router

import (
	"go-triton-app/app/template"
	"net/http"

	"github.com/go-chi/chi"
)

// JSONHandlerFunc is an http.HandlerFunc with a return value representing the result of the handler.
type JSONHandlerFunc func(http.ResponseWriter, *http.Request) *template.JSONResponse

// JSONRouter wraps a chi.Router and provides methods for create JSON-related router.
type JSONRouter struct {
	Core chi.Router
}

// NewJSONRouter creates a JSON router.
func NewJSONRouter() *JSONRouter {
	return &JSONRouter{Core: chi.NewRouter()}
}

// Connect calls chi.Router.Connect with a customized JSONHandlerFunc.
func (r *JSONRouter) Connect(pattern string, h JSONHandlerFunc) {
	r.Core.Connect(pattern, jsonHandlerToHTTPHandler(h))
}

// Delete calls chi.Router.Delete with a customized JSONHandlerFunc.
func (r *JSONRouter) Delete(pattern string, h JSONHandlerFunc) {
	r.Core.Delete(pattern, jsonHandlerToHTTPHandler(h))
}

// Get calls chi.Router.Get with a customized JSONHandlerFunc.
func (r *JSONRouter) Get(pattern string, h JSONHandlerFunc) {
	r.Core.Get(pattern, jsonHandlerToHTTPHandler(h))
}

// Head calls chi.Router.Head with a customized JSONHandlerFunc.
func (r *JSONRouter) Head(pattern string, h JSONHandlerFunc) {
	r.Core.Head(pattern, jsonHandlerToHTTPHandler(h))
}

// Options calls chi.Router.Options with a customized JSONHandlerFunc.
func (r *JSONRouter) Options(pattern string, h JSONHandlerFunc) {
	r.Core.Options(pattern, jsonHandlerToHTTPHandler(h))
}

// Patch calls chi.Router.Patch with a customized JSONHandlerFunc.
func (r *JSONRouter) Patch(pattern string, h JSONHandlerFunc) {
	r.Core.Patch(pattern, jsonHandlerToHTTPHandler(h))
}

// Post calls chi.Router.Post with a customized JSONHandlerFunc.
func (r *JSONRouter) Post(pattern string, h JSONHandlerFunc) {
	r.Core.Post(pattern, jsonHandlerToHTTPHandler(h))
}

// Put calls chi.Router.Put with a customized JSONHandlerFunc.
func (r *JSONRouter) Put(pattern string, h JSONHandlerFunc) {
	r.Core.Put(pattern, jsonHandlerToHTTPHandler(h))
}

// Trace calls chi.Router.Trace with a customized JSONHandlerFunc.
func (r *JSONRouter) Trace(pattern string, h JSONHandlerFunc) {
	r.Core.Connect(pattern, jsonHandlerToHTTPHandler(h))
}

func jsonHandlerToHTTPHandler(h JSONHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}
