package app

import (
	"encoding/json"
	"errors"
)

// AppConfig is the root configuration type for your application.
type AppConfig struct {
	// IsProduction determines if this app is currently running in production mode.
	IsProduction bool `json:"isProduction"`

	// HTTP holds HTTP-related configuration data.
	HTTP *httpConfig `json:"http"`
}

type httpConfig struct {
	// Port is the listening port of web server.
	Port int `json:"port"`
	// Static defines how server serves static files (optional).
	Static *httpStaticConfig `json:"static"`
}
type httpStaticConfig struct {
	// Pattern is the pattern string used for registering request handler.
	Pattern string `json: route`
	// DirPath is the physical directory path you want to be served.
	DirPath string `json: dirPath`
}

// LoadAppConfig loads an AppConfig from an array of bytes.
func LoadAppConfig(bytes []byte) (*AppConfig, error) {
	var config AppConfig

	err := json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	err = config.validate()
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (config *AppConfig) validate() error {
	httpConfig := config.HTTP
	if httpConfig == nil {
		return errors.New("Missing http config")
	}

	if httpConfig.Port == 0 {
		return errors.New("http.port must not be 0")
	}

	httpStaticConfig := httpConfig.Static
	if httpStaticConfig != nil {
		if httpStaticConfig.Route == "" {
			return errors.New("http.static has been defined, but http.static.route remains empty")
		}
		if httpStaticConfig.DirPath == "" {
			return errors.New("http.static has been defined, but http.static.dirPath remains empty")
		}
	}

	return nil
}
