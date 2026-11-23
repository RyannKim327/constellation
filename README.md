# Constellation

<img src="static/logo.svg" alt="Constellation logo" width="96" height="96" />

Constellation is a small web framework for Go built with the standard library.
You register endpoints, return HTML from a handler, and the framework wraps the
result in a complete page. It has no third-party Go dependencies.

- Declarative routes with `constellation.Endpoint`
- HTML helpers in `src/elements` for every common tag
- Server-rendered pages with Tailwind CSS utilities
- Markdown documentation embedded into the binary and served at `/docs`

## Requirements

- Go 1.27 or newer

## Quick start

```bash
go run .
```

Open <http://localhost:8000>.

## Navigation

### Web

| Path            | Description                              |
| --------------- | ---------------------------------------- |
| `/`             | Landing page and element examples.       |
| `/docs`         | Documentation index.                     |
| `/docs/<slug>`  | One page per Markdown file in `docs/`.    |

### User documentation

- [Getting started](docs/getting-started.md)
- [Routing](docs/routing.md)
- [API](docs/api.md)
- [Elements](docs/elements.md)
- [Rendering](docs/rendering.md)
- [Styling](docs/styling.md)

### Contributing

- [Contribution guide](contrib/README.md)
- [Architecture](contrib/architecture.md)
- [Documentation workflow](contrib/documentation.md)
- [Code conventions](contrib/conventions.md)

## Project layout

```text
main.go                  entry point, embedded assets, example page
go.mod
src/interface.go         Endpoint, WebComponent, methods, permissions
src/main.go              server startup and route registration
src/middleware/http.go   HTML page wrapper and response headers
src/elements/            HTML element helpers
src/docs/                Markdown loading and /docs pages
static/styles.css        site stylesheet
docs/                    user guides (published at /docs)
contrib/                 contributor notes (repository only)
```

## Minimal example

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

`StartServer` registers each endpoint with `http.ServeMux` as an exact
`METHOD path` pattern and renders the handler's return value through
`middleware.WebParser`.

## Styling

`static/styles.css` is embedded with `embed.FS` and passed to `StartServer`.
The page wrapper also loads Tailwind CSS from its CDN, so utility classes can be
set through a component's `Class` field. See [the styling guide](docs/styling.md)
for the full setup.

## Development checks

```bash
gofmt -l .
go vet ./...
go build ./...
```

## License

Released under the [MIT License](LICENSE.md).
