package app

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go-triton-app/app/config"
	"go-triton-app/app/logx"
	"go-triton-app/app/template"
)

// Config is the application configuration loaded.
var Config *config.Config

// Logger is the main logger for this app.
var Logger *logx.Logger

// TemplateManager is a app-wide instance of template.Manager.
var TemplateManager *template.Manager

// HTMLResponse returns common objects used to compose an HTML response.
func HTMLResponse(w http.ResponseWriter, r *http.Request) *template.HTMLResponse {
	ctx := r.Context()
	tm := TemplateManager
	resp := template.NewHTMLResponse(ctx, tm, w)

	return resp
}

// JSONResponse returns common objects used to compose an HTML response.
func JSONResponse(w http.ResponseWriter, r *http.Request) *template.JSONResponse {
	ctx := r.Context()
	tm := TemplateManager
	resp := template.NewJSONResponse(ctx, tm, w, Config.Debug)
	return resp
}

// MasterPageData wraps a call to MasterPageData.
func MasterPageData(title, contentHTML string) *template.MasterPageData {
	return template.NewMasterPageData(title, contentHTML)
}

func init() {
	mustSetupConfig()
	mustSetupLogger()
	mustSetupTemplates(Config)
}

func mustSetupConfig() {
	// Parse command-line arguments
	var configPath string
	flag.StringVar(&configPath, "config", "", "path of application config file")
	flag.Parse()

	if configPath == "" {
		// If --config is not specified, check if user has an extra argument like "go run main.go dev", which we consider it as --config "./config/dev.json"
		userArgs := os.Args[1:]
		if len(userArgs) >= 1 {
			configPath = fmt.Sprintf("./config/%v.json", userArgs[0])
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

	log.Printf("✅ Loaded config at \"%v\"", configPath)
	if config.DevMode() {
		log.Printf("⚠️ Application running in dev mode")
	}
	Config = config
}

func mustSetupLogger() {
	if Config == nil {
		panic("Config must be set before mustSetupLogger")
	}
	logger, err := logx.NewLogger(Config.Log.Dir, Config.DevMode())
	if err != nil {
		panic(err)
	}
	Logger = logger
}

func mustSetupTemplates(config *config.Config) {
	templatesConfig := config.Templates
	localizationConfig := config.Localization

	TemplateManager = template.MustCreateManager(templatesConfig.Dir, localizationConfig.Dir, localizationConfig.DefaultLang, Logger, config.Debug)
}
