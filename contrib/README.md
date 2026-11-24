# Contributing

These notes are for people working on Constellation itself. Guides for people
using the framework live in [docs/](../docs/) and are published at `/docs`.

## Setup

You need Go 1.27 or newer. The project uses only the standard library.

```bash
go run .
```

The site is served at `http://localhost:8000`. For hot reload during
development, run `go run ./lib/dev` instead.

## Layout

| Path                  | Responsibility                                      |
| --------------------- | --------------------------------------------------- |
| `main.go`             | Embeds assets, registers endpoints, starts server.  |
| `lib/interface.go`    | Public types, methods, permissions, endpoint registry. |
| `lib/main.go`         | Server startup and route registration.              |
| `lib/middleware/`     | HTML page wrapper and response headers.             |
| `lib/elements/`       | One helper per HTML tag.                            |
| `lib/docs/`           | Markdown loading and the `/docs` pages.             |
| `lib/dev/`            | Development server with hot reload.                 |
| `src/api/`            | Example API endpoint.                               |
| `src/web/`            | Example web endpoint.                               |
| `static/styles.css`   | Site stylesheet, embedded at build time.            |
| `docs/`               | User guides (Markdown source for `/docs`).          |
| `contrib/`            | Contributor notes (not published on the site).      |

## Before submitting

```bash
gofmt -w .
go vet ./...
go build ./...
```

Keep changes small and focused. Update the relevant Markdown in `docs/` or
`contrib/` when behavior changes.

## Further reading

- [Architecture](architecture.md)
- [Documentation workflow](documentation.md)
- [Code conventions](conventions.md)
