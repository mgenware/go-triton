package app

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/mgenware/go-web-boilerplate/app/config"
	"github.com/mgenware/go-web-boilerplate/app/template"
)

// Config is the application configuration loaded.
var Config *config.Config

// Template is the template manager.
var Template *template.TemplateManager

func init() {
	setupConfigOrPanic()
	setupTemplateOrPanic(Config)
}

// ------- Private functions -------
func setupConfigOrPanic() {
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

func setupTemplateOrPanic(config *config.Config) {
	t := template.MustCreateTemplateManager(config.TemplatesDir, !config.IsProduction)
	Template = t
}
