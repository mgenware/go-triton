# go-triton

<img src="./triton.jpg" width="300" height="300"/>

A boilerplate template for Go web applications.

* Configuration file support.
* `IsProduction` flag.
* Implemented common HTTP handlers:
  * Not found(404) handler.
  * Panic recovery handler as 500 Internal Server Error.
* Template support (reloads templates when `IsProduction` flag is `false`).
* Serves static files in development mode.

## Dependencies
* [chi](https://github.com/go-chi/chi): HTTP routing.
* templatex from [go-packagex](https://github.com/mgenware/go-packagex): template wrapper around Go text/template.
* httpx from [go-packagex](https://github.com/mgenware/go-packagex): for common MIME type constants.

## Usage
Start in development mode:
```sh
go run main.go dev
```

Start with a config file:
```sh
go run main.go --config ./configs/prod.json
```
