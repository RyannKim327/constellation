# Styling

The site stylesheet lives in `static/styles.css`. It is embedded into the
binary at build time, so the server has no runtime file dependency for CSS.

## Embedding static files with `embed.FS`

Static assets are embedded at compile time. The `all:` prefix includes every
file under `static/`, including dotfiles.

```go
import "embed"

//go:embed all:static
var StaticFS embed.FS

func main() {
    css, err := StaticFS.ReadFile("static/styles.css")
    if err != nil {
        panic(err)
    }

    constellation.StartServer(css)
}
```

`StartServer` calls `middleware.SetStylesheet`, and `middleware.WebParser`
writes that CSS in a `<style type="text/tailwindcss">` block in every page, so
the stylesheet can use Tailwind features. Only `styles.css` is injected
automatically; other files under `static/` are embedded but are not served yet.

## Tailwind CSS

The page wrapper loads the Tailwind CSS v4 browser build from
`https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4`. The browser build is meant
for development; compile a stylesheet for production. Add any Tailwind utility
class through the `Class` field of a `WebComponent`:

```go
elements.Div(constellation.WebComponent{
    Class: "mx-auto max-w-[1120px] p-8 text-slate-200 md:grid md:grid-cols-2 md:gap-8",
    Children: []any{
        elements.H1(constellation.WebComponent{
            Class: "text-3xl font-bold text-orange-400",
            Text:  "Hello",
        }),
    },
})
```

Define design tokens with `@theme` in `static/styles.css`:

```css
@theme {
  --color-primary: #fb923c;
}
```

This creates utilities such as `bg-primary`, `text-primary`, and
`border-primary`, and keeps `var(--color-primary)` available in custom CSS.
Add more `--color-*`, `--font-*`, or `--breakpoint-*` variables to extend the
theme.

Use `Class` for layout and spacing. Edit `static/styles.css` for shared
component styles such as `.page-shell` and `.markdown`. Restart the server, or
run `go run ./lib/dev` to reload automatically.

## Existing classes

| Class           | Purpose                                   |
| --------------- | ----------------------------------------- |
| `page-shell`    | Centers the page and sets its width.      |
| `page-nav`      | Top navigation row.                       |
| `page-header`   | Title block for inner pages.              |
| `docs-grid`     | Grid of documentation cards.              |
| `docs-card`     | A card linking to one guide.              |
| `docs-layout`   | Sidebar plus article layout.              |
| `docs-sidebar`  | Sticky list of documents.                 |
| `markdown`      | Styles headings, code, lists, and tables. |
| `page-footer`   | Footer row.                               |
| `status-page`   | Centered layout for 404 and 403 pages.    |
| `status-code`   | Large gradient status number.             |
| `status-actions`| Row of status-page buttons.               |
| `status-path`   | Requested path shown on status pages.     |

Colors are defined as theme variables in `@theme`. Edit `styles.css` and restart
the server, or use `go run ./lib/dev` for automatic reloads.
