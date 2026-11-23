# Getting Started

Constellation is a small web framework built with Go's standard library. You declare
endpoints, return HTML from a handler, and the framework wraps it in a page.

It requires Go 1.27 or newer and has no third-party dependencies.

## Project layout

```text
main.go                  application entry point and embedded static files
src/interface.go         Endpoint, WebComponent, and constants
src/main.go              server startup
src/middleware/http.go   HTML page wrapper
src/elements/            HTML element helpers
static/styles.css        site stylesheet
docs/                    Markdown documentation
```

## A minimal app

```go
package main

import constellation "constellation/src"

func main() {
    constellation.AddEndpoint(constellation.Endpoint{
        Path:       "/",
        Method:     constellation.GET,
        Permission: constellation.Unauthenticated,
        Title:      "Home",
        Handler: constellation.HandlerSetup{
            Web: func(r *constellation.Request) string {
                return "<h1>Hello, Constellation</h1>"
            },
        },
    })

    constellation.StartServer(nil)
}
```

Run it with `go run .` and open `http://localhost:8000`.

## Adding styles

Embed a CSS file and pass its bytes to `StartServer`:

```go
//go:embed static/styles.css
var stylesheet []byte

func main() {
    // ... register endpoints ...
    constellation.StartServer(stylesheet)
}
```

See [Styling](styling.md) for details.
