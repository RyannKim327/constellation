# Architecture

Constellation is a thin layer over [Gin](https://github.com/gin-gonic/gin). It is
split into two responsibilities:

- `config/` — the **framework**: bootstrap, middleware, route registration,
  rendering utilities, and the server lifecycle.
- `src/` — the **application**: the entrypoint, your route definitions, and
  static assets.

## Boot sequence

```
main()                        src/main.go
  └── config.Constellation(routes)   main.go
        ├── mw.CorsSetup()            config/middleware/server-controller.go
        │     ├── CORS middleware
        │     └── response header middleware
        ├── for each route:
        │     └── mw.Register(route)  # normalize + attach middleware
        └── mw.StartServer()          # static files, logging, 404, Run()
```

1. **Route collection.** `src/main.go` builds a `[]mw.Route` slice and passes it
   to `config.Constellation`.
2. **CORS setup.** `CorsSetup` registers CORS (origins from `ALLOWED_ORIGINS`) and
   a small response-header middleware.
3. **Registration.** `Register` normalizes each route and attaches the correct
   middleware chain based on method and permission.
4. **Server start.** `StartServer` serves static files, installs a request logger
   and a custom 404 handler, and calls `app.Run`.

## The global engine

`server-controller.go` holds a package-level `gin.Default()` engine:

```go
var app = gin.Default()
```

All registration happens against this single engine. `gin.Default()` already
includes Gin's Logger and Recovery middleware.

## Handler model

Constellation handlers return data instead of writing it directly:

```go
type HandlerSetup func(ctx *gin.Context) (string, map[string]any)
```

The wrapper `_routesHandler` inspects the result:

- `(nil, nil)` — the handler wrote the response itself; nothing else happens.
- otherwise — the body and params are passed to `utils.Render`, which produces a
  full HTML document.

This lets JSON endpoints and server-rendered pages share one handler type.

## Ordering note

Gin copies a route's middleware chain at registration time. Middleware added with
`app.Use(...)` **after** a route is registered does not apply to that route.
Constellation registers CORS and the developer header before routes, so they
apply everywhere; the request logger added in `StartServer` is only seen by
routes registered after it. Keep this in mind if you add middleware.

## Generated documentation

Markdown files in `docs/` are the source of truth for the documentation pages.
`tools/md2go` converts each file into a Go constant under `src/routes/` and
writes a manifest of all documents. `src/routes/docs.go` turns that manifest into
regular routes at `/docs/<slug>`. No Markdown is read from disk at request time.

Regenerate the files after editing any document:

```bash
go generate ./src/routes
```
