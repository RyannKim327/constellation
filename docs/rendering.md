# Rendering

Constellation renders HTML with a minimal placeholder-substitution helper rather
than Go's `html/template`. It is intentionally small: `{key}` tokens in a string
are replaced with values from a map.

## `utils.HTParser` — raw substitution

```go
func HTParser(body string, params map[string]any) string
```

Replaces every `{key}` with `fmt.Sprint(value)` **without escaping**. Use it for
trusted, pre-built HTML fragments.

```go
list := "<li>Ada</li><li>Grace</li>"
html := utils.HTParser(`<ul>{items}</ul>`, map[string]any{"items": list})
```

> **Security:** `HTParser` does not escape values. Never pass untrusted user
> input through it. Use the handler's parameter return value (rendered by
> `Render`) for anything user-controlled.

## `utils.Render` — full document

```go
func Render(ctx *gin.Context, body string, params map[string]any)
```

Wraps `body` in a complete HTML document (charset, viewport, title, and the
`/static/styles.css` stylesheet), escapes the parameters via Go's
`html.EscapeString`, and writes the result with `text/html`.

Handler return values flow here automatically:

```go
func handler(ctx *gin.Context) (string, map[string]any) {
    return `<h1>Hello {name}</h1>`, map[string]any{"name": ctx.Query("name")}
}
```

Because `{name}` is supplied as a parameter, it is HTML-escaped.

## Escaped vs raw

| Helper            | Escapes? | Use for                                  |
| ----------------- | -------- | ---------------------------------------- |
| `utils.HTParser`  | no       | Trusted HTML fragments you built in code |
| `utils.Render`    | yes      | User-supplied data in the parameter map  |

The example page in `src/routes/index.go` uses both: it builds list items and
injects them raw with `HTParser`, then returns `name` as a rendered parameter so
it is escaped.

## Static assets

Files under `src/static` are served at `/static`:

```html
<link rel="stylesheet" href="/static/styles.css" />
```

Add images, fonts, or scripts there and reference them with the `/static/` prefix.
