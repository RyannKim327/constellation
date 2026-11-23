# Rendering

A handler returns the body of the page as a string. The framework supplies the
HTML document around it.

## Handler signature

```go
type HandlerSetup struct {
    Web HandlerWeb // func(r *Request) string
    Api HandlerApi // func(r *Request, w Response)
}
```

Set `Web` to return an HTML body. Return a single element or several joined
strings:

```go
Handler: constellation.HandlerSetup{
    Web: func(r *constellation.Request) string {
        return elements.Main(constellation.WebComponent{
            Class: "page-shell",
            Children: []any{
                elements.H1(constellation.WebComponent{Text: "Hello"}),
                elements.P(constellation.WebComponent{Text: "Built with Go."}),
            },
        })
    },
},
```

## The page wrapper

`middleware.WebParser` writes the response with a `text/html` content type and
these headers:

| Header           | Value                          |
| ---------------- | ------------------------------ |
| `X-Powered-By`   | `Constellation`                |
| `Developed-Under`| `Master Piece of Paper`        |
| `Developed-By`   | `Ryann Kim Sesgundo`           |

It outputs the document type, `lang="en"`, the page title, the embedded CSS,
and the handler body inside a `<main>` element. An empty title becomes
`Constellation`.

## Raw HTML

Handlers return strings, so raw HTML works too. Use it only with trusted or
escaped content; the element helpers escape text and attributes for you.
