# Elements

The `lib/elements` package provides one function per HTML tag. Every function
takes a `WebComponent` and returns an HTML string.

## WebComponent

```go
type WebComponent struct {
    Class      string
    Attributes map[string]string
    Children   []any
    Text       string
}
```

| Field        | Description                                           |
| ------------ | ----------------------------------------------------- |
| `Class`      | Value for the `class` attribute.                      |
| `Attributes` | Other attributes such as `href`, `id`, or `type`.     |
| `Children`   | Nested elements or raw HTML strings.                  |
| `Text`       | Escaped text content.                                 |

## Example

```go
page := elements.Div(constellation.WebComponent{
    Class: "page-shell",
    Children: []any{
        elements.H1(constellation.WebComponent{
            Text: "Welcome",
        }),
        elements.A(constellation.WebComponent{
            Class:      "text-link",
            Attributes: map[string]string{"href": "/docs"},
            Text:       "Read the docs",
        }),
    },
})
```

## Escaping and children

`Text` and attribute values are HTML-escaped. Strings placed directly in
`Children` are treated as raw HTML, which is useful for Markdown output but
should not contain untrusted input. `fmt.Stringer` values are escaped.

Void tags such as `br`, `hr`, `img`, `input`, and `meta` are rendered without a
closing tag. Attributes are sorted by name, so output is stable.

## Available helpers

| File          | Helpers                                      |
| ------------- | -------------------------------------------- |
| `structure.go` | `Div`, `Header`, `Main`, `Nav`, `Section`, `Footer`, ... |
| `headers.go`  | `H1` through `H6`                            |
| `text.go`     | `A`, `P`, `Span`, `Strong`, `Code`, `Pre`, ... |
| `lists.go`    | `Ul`, `Ol`, `Li`, `Dl`, `Dt`, `Dd`            |
| `forms.go`    | `Form`, `Input`, `Label`, `Button`, `Select`, ... |
| `tables.go`   | `Table`, `Thead`, `Tbody`, `Tr`, `Th`, `Td`, ... |
| `media.go`    | `Img`, `Audio`, `Video`, `Iframe`, ...        |
| `document.go` | `Link`, `Meta`, `Script`, `Style`, ...        |
