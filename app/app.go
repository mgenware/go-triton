package app

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/mgenware/go-triton/app/config"
	"github.com/mgenware/go-triton/app/template"
)

// Config is the application configuration loaded.
var Config *config.Config

// TemplateManager is a app-wide instance of template.Manager.
var TemplateManager *template.Manager

// HTMLResponse returns common objects used to compose an HTML response.
func HTMLResponse(w http.ResponseWriter, r *http.Request) (context.Context, *template.Manager, *template.HTMLResponse) {
	ctx := r.Context()
	tm := TemplateManager
	resp := tm.NewHTMLResponse(ctx, w)

	return ctx, tm, resp
}

// JSONResponse returns common objects used to compose an HTML response.
func JSONResponse(w http.ResponseWriter, r *http.Request) (context.Context, *template.Manager, *template.JSONResponse) {
	ctx := r.Context()
	tm := TemplateManager
	resp := tm.NewJSONResponse(w)

	return ctx, tm, resp
}

func init() {
	mustSetupConfig()
	mustSetupTemplates(Config)
}

func mustSetupConfig() {
	// Parse command-line arguments
	var configPath string
	flag.StringVar(&configPath, "config", "", "path of application config file")
	flag.Parse()

	if configPath == "" {
		// If --config is not specified, check if user runs "go run main.go dev" which will read ./configs/dev.json as config file
		userArgs := os.Args[1:]
		if len(userArgs) == 1 && userArgs[0] == "dev" {
			configPath = "./configs/dev.json"
		} else {
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	// Read config file
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	config, err := config.ReadConfig(configBytes)
	if err != nil {
		panic(err)
	}

	log.Printf("Loaded config at \"%v\"", configPath)
	if config.IsProduction {
		log.Printf("[Application runs in production!]")
	}
	Config = config
}

func mustSetupTemplates(c *config.Config) {
	templatesConfig := c.Templates
	localizationConfig := c.Localization

	TemplateManager = template.MustCreateManager(templatesConfig.RootDir, !c.IsProduction, localizationConfig.RootDir, localizationConfig.DefaultLang)
}
