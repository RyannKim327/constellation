# Code conventions

## Packages

- `constellation/lib` holds the public API. Keep it free of imports from
  `elements` and `docs` to avoid dependency cycles.
- `constellation/lib/elements` holds thin HTML tag helpers.
- `constellation/lib/middleware` holds response helpers.
- `constellation/lib/status` holds the built-in 404 and 403 page handlers. Register them from `main.go` with `SetErrorPages` so `lib` stays free of `elements` imports.
- `constellation/lib/docs` holds the documentation routes and renderer.

## Elements

Name helper functions after the HTML tag they render. Keep the body to a single
`render("tag", element)` call. Add new void tags to `voidElements` in
`lib/elements/render.go`.

Pass user input through `Text` or `Attributes`, never as a raw string in
`Children`. Raw strings are reserved for trusted HTML such as rendered
Markdown.

## Endpoints

Register endpoints in `main.go` before `StartServer`. Set `Title` so the page
has a useful title. Use `Unauthenticated` for public endpoints; private endpoints
need checks configured in `lib/middleware/access-controller.go`.

## Style and checks

Format with `gofmt`, then run:

```bash
go vet ./...
go build ./...
```

Follow existing naming and keep comments short. Update `docs/` for user-facing
changes and `contrib/` for internal changes.
