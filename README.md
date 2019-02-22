# go-triton

<img src="./static/img/triton.jpg" width="300" height="300"/>

A boilerplate template for Go web applications. Uses Go 1.11 modules.

* Configuration file support.
* Development/production mode (via `config.DevMode`).
* Implemented common HTTP handlers:
  * Not found(404) handler.
  * Panic recovery handler as 500 Internal Server Error.
* Template support (auto reloads template in development mode).
* Auto serves static files in development mode.
* i18n support.

## Dependencies
* `github.com/go-chi/chi`: HTTP routing. 
* `github.com/mgenware/go-packagex`: for common helpers like template wrapper, MIME type definitions, etc.
* `golang.org/x/text/language`: HTTP `Accept-Language` header parsing and matching.

## Usage
Start in development mode:
```sh
go run main.go dev
```

Start with a config file:
```sh
go run main.go --config ./configs/prod.json
```
