# Constellation

A small, convention-over-configuration web framework for Go, built on top of
[Gin](https://github.com/gin-gonic/gin). Constellation focuses on declarative
routing, server-rendered HTML with a tiny template helper, and a sensible set of
default middleware.

- **Declarative routes** — describe an endpoint with a `middleware.Route` value.
- **Server-rendered pages** — compose HTML with `{placeholder}` templating.
- **Batteries included** — CORS, bot/scraper detection, cookie auth, request
  logging, custom 404 page, and static file serving.
- **Zero config to start** — sensible defaults, with env-var overrides.
- **Live reload** — ships with an [Air](https://github.com/air-verse/air) config.

---

## Requirements

- Go **1.27+**
- (optional) [Air](https://github.com/air-verse/air) for live reload

## Getting started

```bash
git clone <repository-url>
cd nextgo
go mod download
go run ./src/index.go
```

Then open <http://localhost:8000>.

### Development with live reload

```bash
air
```

Air builds `./src/index.go`, runs the binary, and rebuilds on changes to `.go`,
`.css`, `.js`, and `.html` files. Build artifacts go to `tmp/`.

## Project structure

```
.
├── .air.toml                 # Air live-reload configuration
├── go.mod
├── go.sum
├── config/
│   ├── main.go               # Constellation() bootstrap
│   ├── middleware/
│   │   ├── access-controller.go   # Cookie-based access guard
│   │   ├── scraper-detector.go    # Bot / scraper detection
│   │   ├── route-controller.go    # Route type
│   │   └── server-controller.go   # CORS, route registration, server start
│   └── utils/
│       ├── render.go         # HTML templating and page rendering
│       └── statics.go        # HTTP method + permission constants, handler type
└── src/
    ├── index.go              # Entrypoint and route registry
    ├── backend/api.go        # Example JSON API route
    ├── routes/index.go       # Example server-rendered page
    └── static/
        └── styles.css        # Stylesheet served at /static/styles.css
```

## How it works

The entrypoint registers every route and hands them to the framework:

```go
// src/index.go
var _routes = []mw.Route{
    backend.API,
    routes.Index,
}

func main() {
    c.Constellation(_routes)
}
```

`Constellation` sets up CORS, registers each route, and starts the server.

### Defining a route

A route is a `middleware.Route`:

```go
var Hello = middleware.Route{
    Method: utils.METHOD_GET,   // optional, defaults to GET
    Path:   "/hello",           // lowercased automatically
    Handler: func(ctx *gin.Context) (string, map[string]any) {
        return `<h1>Hello {name}</h1>`, map[string]any{"name": "World"}
    },
}
```

Then add it to `_routes` in `src/index.go`.

### Handler contract

Every handler has the signature:

```go
func(ctx *gin.Context) (string, map[string]any)
```

- Return an **HTML body** and a **parameter map** and Constellation renders a full
  page for you (with escaping applied to the parameters).
- Return `(nil, nil)` after writing the response yourself (e.g. `ctx.JSON(...)`)
  and no rendering happens.

```go
func api(ctx *gin.Context) (string, map[string]any) {
    ctx.JSON(200, gin.H{"message": "Hello"})
    return "", nil
}
```

### Rendering HTML

`utils.HTParser` performs raw `{placeholder}` substitution (no escaping — use it
for trusted, pre-built fragments). `utils.Render` wraps a body in the HTML
document shell and escapes the supplied parameters.

See [docs/rendering.md](docs/rendering.md) for details.

## Configuration

| Variable          | Default                                          | Description                          |
| ----------------- | ------------------------------------------------ | ------------------------------------ |
| `PORT`            | `8000`                                           | Port the server listens on.          |
| `ALLOWED_ORIGINS` | `http://localhost:3000,http://localhost:5173`    | Comma-separated CORS allow-list.     |

The port can also be passed programmatically via `StartServer(port)`; `PORT`
takes precedence, then the argument, then `8000`.

## Built-in routes

| Method | Path   | Description                        |
| ------ | ------ | --------------------------------- |
| `GET`  | `/`    | Server-rendered example page.     |
| `GET`  | `/api` | Example JSON API response.        |
| `GET`  | `/static/*` | Static files from `src/static`. |

## Documentation

- [Architecture overview](docs/architecture.md)
- [Routing](docs/routing.md)
- [Middleware](docs/middleware.md)
- [Rendering](docs/rendering.md)
- [Configuration](docs/configuration.md)

## Known limitations

- `src/routes/index.go` is a demo: it calls an external API and panics on error.
  Replace it with real data fetching and error handling.
- Bot detection in `BotProtection` is a simple `User-Agent` heuristic; harden it
  before relying on it in production.
- Middleware is attached at route-registration time in Gin. CORS and the
  developer header are registered before routes; the request logger added in
  `StartServer` only applies to handlers registered after it.
- `PERMISSION_ADMIN` is defined but not enforced yet.

## License

No license has been specified for this project.
