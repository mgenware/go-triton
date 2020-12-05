package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-triton-app/app/cfg"
	"go-triton-app/app/handler"
	"go-triton-app/app/logx"
)

// Config is the application configuration loaded.
var Config *cfg.Config

// Logger is the main logger for this app.
var Logger *logx.Logger

// MasterPageManager is a app-wide instance of handler.MasterPageManager.
var MasterPageManager *handler.MasterPageManager

// HTMLResponse returns common objects used to compose an HTML response.
func HTMLResponse(w http.ResponseWriter, r *http.Request) *handler.HTMLResponse {
	tm := MasterPageManager
	resp := handler.NewHTMLResponse(r, tm, w)
	return resp
}

// JSONResponse returns common objects used to compose an HTML response.
func JSONResponse(w http.ResponseWriter, r *http.Request) *handler.JSONResponse {
	tm := MasterPageManager
	resp := handler.NewJSONResponse(r, tm, w)
	return resp
}

// PanicIfErr panics if the given `err` is not nil.
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// MasterPageData wraps a call to MasterPageData.
func MasterPageData(title, contentHTML string) *handler.MasterPageData {
	return handler.NewMasterPageData(title, contentHTML)
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
	config, err := cfg.ReadConfig(configPath)
	if err != nil {
		panic(fmt.Errorf("Error reading config file, %v", err))
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

func mustSetupTemplates(config *cfg.Config) {
	templatesConfig := config.Templates
	localizationConfig := config.Localization

	MasterPageManager = handler.MustCreateManager(templatesConfig.Dir, localizationConfig.Dir, localizationConfig.DefaultLang, Logger, config)
}
