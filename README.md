# go-triton

<img src="./static/triton.jpg" width="300" height="300"/>

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
* [chi](https://github.com/go-chi/chi) 3.3.2: HTTP routing.
* templatex from [go-packagex](https://github.com/mgenware/go-packagex) 1.3.2: template wrapper around Go text/template.
* httpx from [go-packagex](https://github.com/mgenware/go-packagex) 1.3.2: for common MIME type constants.

## Usage
Start in development mode:
```sh
go run main.go dev
```

Start with a config file:
```sh
go run main.go --config ./configs/prod.json
```
