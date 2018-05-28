package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/mgenware/go-web-boilerplate/app"
	"github.com/mgenware/go-web-boilerplate/handlers/indexHandler"
	"github.com/mgenware/go-web-boilerplate/handlers/systemHandlers"

	"github.com/go-chi/chi"
)

func Start() {
	r := chi.NewRouter()
	config := app.Config
	httpConfig := config.HTTP

	// ----------------- Middlewares -----------------
	// THE PanicMiddleware MUST BE AT THE VERY BEGINNING, OTHERWISE IT WILL NOT WORK!
	if config.IsProduction {
		// *** Production only ***

		// Mount PanicMiddleware only in production, let panic crash in development
		r.Use(systemHandlers.PanicMiddleware)
	} else {
		// *** Development only ***

		// Mount static file server during development. (You may use to nginx to serve these files in production)
		httpStaticConfig := httpConfig.Static
		if httpStaticConfig != nil {
			log.Printf("Serving Assets(%v) at \"%v\"", httpStaticConfig.Pattern, httpStaticConfig.DirPath)
			fileServer(r, httpStaticConfig.Pattern, http.Dir(httpStaticConfig.DirPath))
		}
	}
	// Mount other middlewares, for example:
	// r.Use(sessionMiddleware)

	// ----------------- HTTP Routes -----------------
	// Not found handler
	r.NotFound(systemHandlers.NotFoundHandler)

	// index handler
	r.Get("/", indexHandler.IndexGET)

	log.Printf("Starting server at %v", httpConfig.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(httpConfig.Port), r))
}

// fileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
