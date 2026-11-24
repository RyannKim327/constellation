# Getting Started

Constellation is a small web framework built with Go's standard library. You declare
endpoints, return HTML from a handler, and the framework wraps it in a page.

It requires Go 1.27 or newer and has no third-party dependencies.

Developed by Ryann Kim Sesgundo with Opencode Big Pickle.

## Project layout

```text
main.go                  application entry point and embedded static files
lib/interface.go         Endpoint, WebComponent, and constants
lib/main.go              server startup
lib/middleware/http.go   HTML page wrapper
lib/elements/            HTML element helpers
lib/dev/                 development server with hot reload
static/styles.css        site stylesheet
docs/                    Markdown documentation
```

## A minimal app

```go
package main

import (
    constellation "constellation/lib"
    "constellation/lib/elements"
)

func main() {
    constellation.AddEndpoint(constellation.Endpoint{
        Path:       "/",
        Method:     constellation.GET,
        Permission: constellation.Unauthenticated,
        Title:      "Home",
        Handler: constellation.HandlerSetup{
            Web: func(r *constellation.Request) string {
                return elements.Main(constellation.WebComponent{
                    Class: "page-shell",
                    Children: []any{
                        elements.H1(constellation.WebComponent{
                            Text: "Hello, Constellation",
                        }),
                        elements.P(constellation.WebComponent{
                            Text: "Built with Go.",
                        }),
                    },
                })
            },
        },
    })

    constellation.StartServer(nil)
}
```

The `elements` helpers return HTML strings, so a `Web` handler can return one
directly. See [Elements](elements.md) for the available helpers.

Run it with `go run .` and open `http://localhost:8000`. For development with
automatic rebuilds and browser refresh, use the watcher instead:

```bash
go run ./lib/dev
```

It watches `.go`, `.css`, and `.md` files and serves the app on the same port.
Set `PORT` to use a different port.

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
