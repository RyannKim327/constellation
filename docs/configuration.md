# Configuration

Constellation is configured through environment variables, with defaults that
work out of the box.

## Environment variables

| Variable          | Default                                       | Used by      | Description                       |
| ----------------- | --------------------------------------------- | ------------ | --------------------------------- |
| `PORT`            | `8000`                                        | `StartServer`| Port the HTTP server listens on.  |
| `ALLOWED_ORIGINS` | `http://localhost:3000,http://localhost:5173` | `CorsSetup`  | Comma-separated CORS allow-list.  |

### Examples

```bash
PORT=3000 go run ./src/index.go
```

```bash
ALLOWED_ORIGINS="http://localhost:5173,https://example.com" go run ./src/index.go
```

## Port resolution

`StartServer(port ...int)` resolves the port in this order:

1. `PORT` environment variable (if it parses as an integer)
2. The first variadic argument, if provided
3. `8000`

```go
mw.StartServer()      // env PORT, else 8000
mw.StartServer(9000)  // env PORT, else 9000
```

## Static files

Static assets are served from `src/static` at the `/static` URL prefix. The path
is fixed in `StartServer`.

## Live reload

`.air.toml` configures [Air](https://github.com/air-verse/air):

- Builds `./src/index.go` into `./tmp/main`.
- Watches `.go`, `.css`, `.js`, and `.html` files.
- Excludes `assets`, `tmp`, `vendor`, `testdata`, and `_test.go` files.

Run it with:

```bash
air
```
