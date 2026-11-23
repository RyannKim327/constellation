# Code conventions

## Packages

- `constellation/src` holds the public API. Keep it free of imports from
  `elements` and `docs` to avoid dependency cycles.
- `constellation/src/elements` holds thin HTML tag helpers.
- `constellation/src/middleware` holds response helpers.
- `constellation/src/docs` holds the documentation routes and renderer.

## Elements

Name helper functions after the HTML tag they render. Keep the body to a single
`render("tag", element)` call. Add new void tags to `voidElements` in
`src/elements/render.go`.

Pass user input through `Text` or `Attributes`, never as a raw string in
`Children`. Raw strings are reserved for trusted HTML such as rendered
Markdown.

## Endpoints

Register endpoints in `main.go` before `StartServer`. Set `Title` so the page
has a useful title. Use `Unauthenticated` until permissions are enforced.

## Style and checks

Format with `gofmt`, then run:

```bash
go vet ./...
go build ./...
```

Follow existing naming and keep comments short. Update `docs/` for user-facing
changes and `contrib/` for internal changes.
