package config

import (
	"encoding/json"
	"errors"
	"path/filepath"
)

// Config is the root configuration type for your application.
type Config struct {
	// IsProduction determines if this app is currently running in production mode.
	IsProduction bool `json:"isProduction"`

	// HTTP holds HTTP-related configuration data.
	HTTP *httpConfig `json:"http"`
	// ViewDir is the directory of templates.
	TemplatesDir string `json:"templatesDir"`
}

// ----- Internal types -----
type httpConfig struct {
	// Port is the listening port of web server.
	Port int `json:"port"`
	// Static defines how server serves static files (optional).
	Static *httpStaticConfig `json:"static"`
}
type httpStaticConfig struct {
	// Pattern is the pattern string used for registering request handler.
	Pattern string `json:"pattern"`
	// Dir is the physical directory path you want to be served.
	Dir string `json:"dir"`
}

// ReadConfig loads an ConfigType from an array of bytes.
func ReadConfig(bytes []byte) (*Config, error) {
	var config Config

	err := json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	err = config.validateAndCoerce()
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (config *Config) validateAndCoerce() error {
	httpConfig := config.HTTP
	if httpConfig == nil {
		return errors.New("Missing http config")
	}

	if httpConfig.Port == 0 {
		return errors.New("http.port must not be 0")
	}

	httpStaticConfig := httpConfig.Static
	if httpStaticConfig != nil {
		if httpStaticConfig.Pattern == "" {
			return errors.New("http.static has been defined, but http.static.pattern remains empty")
		}
		if httpStaticConfig.Dir == "" {
			return errors.New("http.static has been defined, but http.static.dir remains empty")
		}
	}

	mustCoercePath(&config.TemplatesDir)
	if httpStaticConfig != nil {
		mustCoercePath(&httpStaticConfig.Dir)
	}
	return nil
}

func mustCoercePath(p *string) {
	if p == nil {
		return
	}
	if filepath.IsAbs(*p) {
		return
	}
	res, err := filepath.Abs(*p)
	if err != nil {
		panic(err)
	}
	*p = res
}
