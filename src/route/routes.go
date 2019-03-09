package route

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mgenware/go-packagex/iox"

	"go-triton-app/app"
	"go-triton-app/app/logx"
	"go-triton-app/app/middleware"
	"go-triton-app/route/api"
	"go-triton-app/route/errorp"
	"go-triton-app/route/homep"
	"go-triton-app/route/sysh"

	"github.com/go-chi/chi"
)

// Start starts the web router.
func Start() {
	r := chi.NewRouter()
	config := app.Config
	httpConfig := config.HTTP

	// ----------------- Middlewares -----------------
	// THE PanicMiddleware MUST BE AT THE VERY BEGINNING, OTHERWISE IT WILL NOT WORK!
	if !config.DevMode() {
		// *** Production only ***

		// Mount PanicMiddleware only in production, let panic crash in development
		r.Use(sysh.PanicMiddleware)
	}

	// Mount static file server
	httpStaticConfig := httpConfig.Static
	if httpStaticConfig != nil {
		url := httpStaticConfig.URL
		dir := httpStaticConfig.Dir
		app.Logger.LogInfo("Serving Assets", logx.D{
			"url": url,
			"dir": dir,
		})
		fileServer(r, url, http.Dir(dir))
		if !iox.IsDirectory(dir) {
			app.Logger.LogWarning("Assets directory doesn't exist", logx.D{"dir": dir})
		}
	}

	// Mount other middlewares, for example:
	// r.Use(sessionMiddleware)

	// ----------------- HTTP Routes -----------------
	lm := app.TemplateManager.LocalizationManager

	// Not found handler
	r.With(lm.EnableContextLanguage).NotFound(sysh.NotFoundHandler)

	// index handler
	r.With(lm.EnableContextLanguage).Get("/", homep.HomeGET)
	r.With(lm.EnableContextLanguage).Get("/fake_error", errorp.FakeErrorGET)

	// API handler
	r.With(middleware.ParseJSONRequest).Mount("/api", api.Router)

	app.Logger.LogInfo("Server starting", logx.D{"port": httpConfig.Port})
	err := http.ListenAndServe(":"+strconv.Itoa(httpConfig.Port), r)
	if err != nil {
		app.Logger.LogError("Server failed to start", logx.D{"err": err.Error()})
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