# go-triton

<img src="./assets/img/triton.jpg" width="300" height="300"/>

A boilerplate template for Go web applications. Uses Go 1.11 modules.

## Table of Contents

- [Features](#features)
- [Main dependencies](#main-dependencies)
- [Usage](#usage)
- [Directory structure](#directory-structure)
  - [The `r` directory](#the--r--directory)
- [HTTP handlers](#http-handlers)
  - [Define handlers](#define-handlers)
  - [Error handling in handlers](#error-handling-in-handlers)
- [Localization](#localization)
  - [How is user language determined](#how-is-user-language-determined)
  - [Enable Localization on a specific route](#enable-localization-on-a-specific-route)
  - [Use localized strings in templates](#use-localized-strings-in-templates)
- [Logging](#logging)
- [Projects built from go-trion](#projects-built-from-go-trion)

## Features

- Config files support.
- Development and production modes (via `config.DevMode`).
- Builtin support for HTML handlers and JSON-based API handlers.
- Implemented common HTTP handlers:
  - Not found(404) handler.
  - Panic recovery handler as 500 Internal Server Error.
- Template support (auto reloads template in development mode).
- Auto serves static files in development mode.
- i18n support.
- Builtin logs to different files (also configurable).

## Main Dependencies

- `github.com/go-chi/chi`: HTTP routing.
- `github.com/mgenware/goutil`: utility functions and types such as template wrapper, MIME type definitions, etc.
- `golang.org/x/text/language`: HTTP `Accept-Language` header parsing and matching.
- `github.com/uber-go/zap`: Logging.

## Usage

Start server in development mode:

```sh
# Start with ./config/dev.json
go run main.go dev
```

Start server in production mode:

```sh
# Start with ./config/prod.json
go run main.go prod
```

The two commands above simply load a configuration file by the given name, you can also create your own config files like `./config/myName.json` and start server with it:

```sh
go run main.go myName
```

Or use the `--config` argument to specify a file:

```sh
go run main.go --config /etc/my_server/dev.json
```

## Directory structure

```
├── appdata             Application generated files, e.g. logs, git ignored
│   └── log
├── assets              Static assets, HTML/JavaScript/CSS/Image files
├── localization        Localization resources
│   └── langs               Localized strings used by your app
└── templates           Go HTML template files
├── src                 Go source directory
│   ├── app                 Core app modules, such as template manager, logger, etc.
│   ├── config              Config files
│   │   ├── dev.json
│   │   └── prod.json
│   ├── r               Routes
```

### The `r` directory

The `r`(`routes`) directory contains all routes of your application. In order to follow the best practices for package naming ([details](https://blog.golang.org/package-names)), child directories of `r` usually consist of a short name plus a letter indicating the type of the route, e.g. `sysh` for system handlers, `homep` for home page, etc.

## HTTP handlers

### Define handlers

An HTML GET handler example:

```go
// Home page template.
var homeView = app.MainPageManager.MustParseLocalizedView("home.html")

// Home page GET handler.
func HomeGET(w http.ResponseWriter, r *http.Request) handler.HTML {
	// Create an HTML response.
	resp := app.HTMLResponse(w, r)
	// Prepare home page data.
	pageData := &HomePageData{Time: time.Now().String()}
	// Generate page HTML.
	pageHTML := homeView.MustExecuteToString(resp.Lang(), pageData)
	// Create main page data, which is a core template shared by all your website pages.
	d := app.MainPageData(resp.LocalizedDictionary().Home, pageHTML)
	// Complete the response.
	return resp.MustComplete(d)
}
```

A JSON API POST handler example:

```go
// Handler for a JSON-based POST API.
func jsonAPI(w http.ResponseWriter, r *http.Request) handler.JSON {
	// Create a JSON response.
	resp := app.JSONResponse(w, r)
	// Fetch some data from the request.
	dict := defs.BodyContext(r.Context())
	// Complete the response.
	return resp.MustComplete(dict)
}
```

### Error handling in handlers

You can simply panic in handler code, it will be handled accordingly based on handler type. See [panic_handlers.go](https://github.com/mgenware/go-triton/blob/main/src/r/sysh/panic_handlers.go).

- For HTML handlers, `panic` results in the error page to be rendered, which corresponds to `error.html` template.
- For API handlers, `panic` results in a generic error response.

Alternatively, if you don't like `panic`, both `HTMLResponse` and `JSONResponse` have functions to complete the response by an error. For example:

```go
// Handler for a JSON-based POST API.
func jsonAPI(w http.ResponseWriter, r *http.Request) handler.JSON {
	// Create a JSON response.
	resp := app.JSONResponse(w, r)
	// Fetch some data from the request.
	dict := defs.BodyContext(r.Context())
	if (len(dict) == 0) {
		// Return an error response.
		return resp.MustFail(fmt.Errorf("Error: invalid input."))
	}
	// Complete the response.
	return resp.MustComplete(dict)
}
```

## Localization

### How is user language determined

- If the user explicitly specified the language ID in query string like (`/?lang=en`), we take user's input as desired language ID (this also sets the language ID in cookies).
- If not, try using saved language ID in user cookies.
- If not, determine desired language from HTTP headers.

### Enable localization on a specific route

The process of determining the desired language ID can brings some costs, so localization is disabled by default. To enable localization in a specific HTTP route, mount the `EnableContextLanguage` middleware.

```go
r := chi.NewRouter()
// lm is app localization manager
lm := app.TemplateManager.LocalizationManager
// Enable localization on home page handler
r.With(lm.EnableContextLanguage).Get("/", homep.HomeGET)
```

### Use localized strings in templates

Once localization is enabled, you can access them in HTML templates. Localized strings are stored as JSON files in `/localization/langs` with file name indicating the language ID. Go-triton comes with two example localized strings files, `en.json` for English, and `cs.json` for `Chinese Simplified`.

To reference a localized string, you need to first make your template data type derive from `template.LocalizedTemplateData`.

```go
import "go-triton-app/app/template"

// HomePageData contains the information needed for generating the home page.
type HomePageData struct {
	template.LocalizedTemplateData

	MyTemplateField1 string
	MyTemplateField2 string
}
```

Let's say our localized strings files are defined as follows:

`en.json`:

```json
{
  "home": "Home",
  "helloWorld": "Hello world!"
}
```

`cs.json`:

```json
{
  "home": "主页",
  "helloWorld": "你好，世界！"
}
```

You can now reference localized strings in templates by through the `LS` field:

```html
<div>
  <!-- Accessing localized fields -->
  <h1>{{html .LS.home}}</h1>
  <p>{{html .LS.helloWorld}}</p>
  <!-- Accessing non-localized fields -->
  <p>{{html .MyTemplateField1}}</p>
  <p>{{html .MyTemplateField2}}</p>
</div>
```

If you got `.LS` not defined error, it's probably because you didn't have your template class derive from `template.LocalizedTemplateData`.

## Logging

By default, go-triton logs to the following files:

- `error.log` errors (by calling `panic` with an `Error` or `app.Logger.Error`)
- `warning.log` warnings (by `app.Logger.Warn`)
- `info` info (by `app.Logger.Info`)
- `not_found` logs all 404 requests

These files are also saved in different directories based on configuration:

- `appdata/log/dev` development logs
- `appdata/log/prod` production logs

All settings above are configurable.

## Projects built from go-triton

- [qing](https://github.com/mgenware/qing)
