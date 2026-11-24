# Architecture

Constellation is a thin layer over Go's `net/http`. The core never imports the
element helpers or the docs pages; those packages import the core instead.

## Boot sequence

```text
main()
  ├── docs.Register(DocsFS)       register /docs and /docs/<slug>
  ├── AddEndpoint(...)            register application endpoints
  ├── StaticFS.ReadFile(css)      read embedded stylesheet
  └── StartServer(css)
        ├── SetStylesheet(css)    store CSS for the page wrapper
        ├── ServeMux.HandleFunc   one exact "METHOD path" pattern per endpoint
        └── ListenAndServe        listen on :8000
```

## Endpoint registry

`src/interface.go` defines `Endpoint` and a package-level `Endpoints` slice.
`AddEndpoint` normalizes empty paths to `/`, empty methods to `GET`, and empty
permissions to `Unauthenticated` before appending. Permission values are
enforced by `middleware.AccessController` before the endpoint handler runs.

## Request flow

`src/main.go` matches the request with `http.ServeMux`, runs
`middleware.AccessController`, calls the endpoint's handler with the request cast
to `*Request`, and passes the returned HTML string to `middleware.WebParser`.
The wrapper writes the full HTML document, the response headers, the title, and
the embedded CSS.

## Element rendering

`src/elements/render.go` contains the shared renderer:

- `Class` and `Attributes` are written as HTML attributes, sorted by name.
- `Text` and attribute values are escaped with `html.EscapeString`.
- Strings in `Children` are inserted as raw HTML.
- Void tags are rendered without a closing tag.

Each helper in the other files is a thin `render("tag", element)` call.

## Documentation pages

`main.go` embeds `docs/` with `go:embed docs` and passes the `embed.FS` to
`docs.Register`. `src/docs/web.go` reads each Markdown file once at startup,
derives its title and summary, and registers a normal endpoint for it.
`src/docs/markdown.go` converts the trusted Markdown to HTML with a small
standard-library renderer. Nothing is read from disk while serving requests.
