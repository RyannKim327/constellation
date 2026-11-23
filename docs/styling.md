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
writes that CSS in a plain `<style>` block in every page. The custom classes
work even if the Tailwind CDN is unavailable. Only `styles.css` is injected
automatically; other files under `static/` are embedded but are not served yet.

## Tailwind CSS

The page wrapper loads `https://cdn.tailwindcss.com?plugins=typography`. Add any
Tailwind utility class through the `Class` field of a `WebComponent`:

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

Use `Class` for layout and spacing. Edit `static/styles.css` for shared
component styles such as `.page-shell` and `.markdown`, then restart the server
to apply changes.

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

Colors are defined as CSS variables in `:root`. Edit `styles.css` and restart
the server to apply changes.
