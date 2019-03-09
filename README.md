# go-triton

<img src="./assets/img/triton.jpg" width="300" height="300"/>

A boilerplate template for Go web applications. Uses Go 1.11 modules.

* Configuration file support.
* Development/production mode (via `config.DevMode`).
* Implemented common HTTP handlers:
  * Not found(404) handler.
  * Panic recovery handler as 500 Internal Server Error.
* Template support (auto reloads template in development mode).
* Auto serves static files in development mode.
* i18n support.

## Main Dependencies
* `github.com/go-chi/chi`: HTTP routing. 
* `github.com/mgenware/go-packagex`: for common helpers like template wrapper, MIME type definitions, etc.
* `golang.org/x/text/language`: HTTP `Accept-Language` header parsing and matching.
* `github.com/sirupsen/logrus`: Logging.

## Usage
Start in development mode:
```sh
# Start with ./config/dev.json
go run main.go dev
```

Start in production mode:
```sh
# Start with ./config/prod.json
go run main.go prod
```

The two commands simply load a configuration file by the given name, you can also create your own config file like `./config/myName.json` and start the app with it:
```sh
go run main.go myName
```

Or use the `--config` argument to specify a file:
```sh
go run main.go --config /etc/my_server/dev.json
```

## Directory Structure
```
├── appdata             Application generated files, e.g. logs, git ignored
│   └── log
├── assets              Static assets, HTML/JavaScript/CSS/Image files
├── localization        Localization resources
│   └── langs               Localized strings used by your app
├── src                 Go source directory
│   ├── app                 Core app modules, such as template manager, logger, etc.
│   ├── config              Config files
│   │   ├── dev.json
│   │   └── prod.json
│   ├── handlers        Web handlers
└── templates           Go HTML template files 
```

## Projects built from go-trion
* [qing](https://github.com/mgenware/qing)