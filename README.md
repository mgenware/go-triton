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
* Builtin logging to different files based on settings or sources.

### Table of Contents
* [Main Dependencies](#main-dependencies)
* [Usage](#usage)
* [Directory Structure](#directory-structure)
	+ [The `r` Directory](#the--r--directory)
* [Error handling in HTTP handlers](#error-handling-in-http-handlers)
* [Localization](#localization)
	+ [How is user language determined](#how-is-user-language-determined)
	+ [Enabling Localization on a specific route](#enabling-localization-on-a-specific-route)
	+ [Using localized strings in templates](#using-localized-strings-in-templates)
* [Logging](#logging)
* [Projects built from go-trion](#projects-built-from-go-trion)

## Main Dependencies
* `github.com/go-chi/chi`: HTTP routing. 
* `github.com/mgenware/go-packagex`: for common helpers like template wrapper, MIME type definitions, etc.
* `golang.org/x/text/language`: HTTP `Accept-Language` header parsing and matching.
* `github.com/uber-go/zap`: Logging.

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
└── templates           Go HTML template files 
├── src                 Go source directory
│   ├── app                 Core app modules, such as template manager, logger, etc.
│   ├── config              Config files
│   │   ├── dev.json
│   │   └── prod.json
│   ├── r               Routes
```

### The `r` Directory
The `r`(`routes`) directory contains all routes of your application, and because it is commonly used so we shortens in to `r`, and in order to follow the best practices for package naming ([details](https://blog.golang.org/package-names)), child directories of `r` usually consist of a short name plus a letter indicating the type of the route, e.g. `sysh` for system handlers, `homep` for home page stuff, etc.

## Error handling in HTTP handlers
Two styles of error handling are supported, "panic" style and "return" style.

* Panic style, handles errors by panicking.
  * Pros: no "double writing" issue when you forgot to `return` after writing content to response.
  * Cons: may look a bit weird to some devs as `panic` is supposed to crash the process.

* Return style, handles error and result in a similar way.
  * Pros: no `panic` in handler code.
  * Cons: handler is screwed when you forgot to `return` after writing content to response.

```go
// "panic" style
func formAPI1(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		// `panic` with a string to indicate an expected error (or user error)
		// Expected errors are not logged and served with normal 200 HTTP code.
		panic("The argument \"id\" cannot be empty")
	}
	result, err := systemCall()
	if err != nil {
		// `panic` with an error to indicate an unexpected error (or app error)
		// Unexpected errors are considered fatal and happened when something
		//  went wrong in your code or system. They are logged and served
		//  with 500 (Internal Server Error) HTTP code.
		panic(err)
	}
	resp := app.JSONResponse(w, r)
	resp.MustComplete(result)
}

// "return" style
func formAPI2(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	resp := app.JSONResponse(w, r)
	if id == "" {
		// `panic` with a string to indicate an expected error (or user error)
		// Expected errors are not logged and served with normal 200 HTTP code.
		resp.MustFailWithUserError("The argument \"id\" cannot be empty")
		// DON'T FORGET THE `return`
		return
	}
	result, err := systemCall()
	if err != nil {
		// For unexpected errors (or app errors), call `MustFail`.
		resp.MustFail(err)
		// DON'T FORGET THE `return`
		return
	}
	resp.MustComplete(result)
}
```

## Localization
### How is user language determined
* If user explicitly specify the language ID in query string like (`/?lang=en`), then we take user's input as desired language ID (this also sets the language ID to user cookies).
* If not, try using saved language ID in user cookies.
* If not, determine desired language from HTTP headers.

### Enabling Localization on a specific route
As shown above, the process of determining the desired language ID can definitely brings some cost, so localization is disabled by default. To support localization in a specific HTTP route, we need to mount the `EnableContextLanguage` middleware.

```go
r := chi.NewRouter()
// lm is app localization manager
lm := app.TemplateManager.LocalizationManager
// Enable localization on home page handler
r.With(lm.EnableContextLanguage).Get("/", homep.HomeGET)
```

### Using localized strings in templates
One localization is enabled, we can access them in HTML templates. Localized strings are stored as JSON files in `/localization/langs` with file name indicating the language ID. Go-triton comes with two example localized strings files, `en.json` for English, and `cs.json` for `Chinese Simplified`.

To reference a locaized string, you need to first make your template data type derive from `template.LocalizedTemplateData`. e.g. the `home_page_data.go` in project:

```go
import "go-triton-app/app/template"

// HomePageData contains the information needed for generating the home page.
type HomePageData struct {
	template.LocalizedTemplateData

	MyTemplateField1 string
	MyTemplateField2 string
}
```

Let's say our localized strings files are defined like this:

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

You can now reference localized strings in template by accessing the `LS` field:

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

If you got `.LS` not defined then it's probably because you didn't have you template class derive from `template.LocalizedTemplateData`.

## Logging
By default, go-triton logs to the following files:
* `error.log` errors (by calling `panic` with an `Error` or `app.Logger.Error`)
* `warning.log` warnings (by `app.Logger.Warn`)
* `info` info (by `app.Logger.Info`)
* `not_found` logs all 404 requests.

This files are also saved in different directories based on configuration:
* `appdata/log/dev` development logs
* `appdata/log/prod` production logs

All above settings are configurable.

## Projects built from go-trion
* [qing](https://github.com/mgenware/qing)
