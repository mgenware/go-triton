# go-triton

<img src="./static/img/triton.jpg" width="300" height="300"/>

A boilerplate template for Go web applications.

* Configuration file support.
* `IsProduction` flag.
* Implemented common HTTP handlers:
  * Not found(404) handler.
  * Panic recovery handler as 500 Internal Server Error.
* Template support (reloads templates when `IsProduction` flag is `false`).
* Serves static files in development mode.
* i18n support.

## Dependencies
* `github.com/go-chi/chi` 3.3.3: HTTP routing. 
* `github.com/mgenware/go-packagex` 2.0.0: for common helpers like template wrapper, MIME type definitions, etc.
* `golang.org/x/text/language` 0.3.0 : HTTP `Accept-Language` header parsing and matching.

## Usage
Install dependencies:
```sh
dep ensure
```

Start in development mode:
```sh
go run main.go dev
```

Start with a config file:
```sh
go run main.go --config ./configs/prod.json
```
