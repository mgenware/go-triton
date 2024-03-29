package r

import (
	"net/http"
	"strconv"
	"strings"

	"go-triton-app/app"
	"go-triton-app/app/handler"
	"go-triton-app/app/middleware"
	"go-triton-app/r/api"
	"go-triton-app/r/errorp"
	"go-triton-app/r/homep"
	"go-triton-app/r/sysh"

	"github.com/go-chi/chi"
	"github.com/mgenware/goutil/iox"
)

// Start starts the web router.
func Start() {
	r := chi.NewRouter()
	config := app.Config
	httpConfig := config.HTTP

	// ----------------- Middlewares -----------------
	// THE PanicMiddleware MUST BE AT THE VERY BEGINNING, OTHERWISE IT WILL NOT WORK!
	r.Use(sysh.PanicMiddleware)

	// Mount static file server
	httpStaticConfig := httpConfig.Static
	if httpStaticConfig != nil {
		url := httpStaticConfig.URL
		dir := httpStaticConfig.Dir
		app.Logger.Info("serving-assets",
			"url", url,
			"dir", dir,
		)
		fileServer(r, url, http.Dir(dir))
		if !iox.IsDirectory(dir) {
			app.Logger.Warn("serving-assets.not-found", "dir", dir)
		}
	}

	// Mount other middlewares, for example:
	// r.Use(sessionMiddleware)

	// ----------------- HTTP Routes -----------------
	lm := app.MainPageManager.LocalizationManager

	// Not found handler
	r.With(lm.EnableContextLanguage).NotFound(sysh.NotFoundHandler)

	// index handler
	r.With(lm.EnableContextLanguage).Get("/", handler.HTMLHandlerToHTTPHandler(homep.HomeGET))
	r.With(lm.EnableContextLanguage).Get("/internal_error", errorp.InternalErrorGET)

	// API handler
	r.With(middleware.ParseJSONRequest).Mount("/api", api.Router.Core)

	app.Logger.Info("server-starting", "port", httpConfig.Port)
	err := http.ListenAndServe(":"+strconv.Itoa(httpConfig.Port), r)
	if err != nil {
		app.Logger.Error("server-starting.failed", "err", err.Error())
		panic(err)
	}
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
