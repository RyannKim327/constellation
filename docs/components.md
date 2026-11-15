# Components

Reusable chunks of HTML live in `src/components/` as plain Go functions and
variables that return a `string`. There is no component runtime: a component is
just code that builds markup, so you can call it from any route or from another
component.

The canonical example is `src/components/header.go`, which renders the shared
navigation bar and highlights the current page.

## Where components live

```
src/components/
└── header.go   # package components
```

Every file in the folder uses `package components`. Routes import it with the
application's module path:

```go
import "github.com/RyannKim327/constellation/src/components"
```

## Two ways to use components

A component is either a **variable** or a **function**:

- **Variable** — no parameters. Assign the markup directly and reference it by
  name.
- **Function** — takes parameters, so call it with the data it needs.

### Variable: no parameters

Use a `var` when the markup is always the same. Store the raw HTML string:

```go
// src/components/footer.go
package components

var Footer = `<footer>Developed by Ryann Kim Sesgundo</footer>`
```

Reference the variable (no parentheses) in the page's parameter map:

```go
return utils.HTParser(`{footer}`, utils.PARAMS{
    "footer": components.Footer,
}), nil
```

The page token `{footer}` is replaced with the value of `components.Footer`.

### Function: with parameters

Use a `func` when the markup depends on input. It returns the built string:

```go
// src/components/header.go
func Header(currentPage string) string {
    /* ... builds the nav using currentPage ... */
}
```

Call the function (with parentheses) and pass the result:

```go
return utils.HTParser(`{navigation}`, utils.PARAMS{
    "navigation": components.Header(ctx.Request.URL.Path),
}), nil
```

The rule of thumb: **use `var` if you don't pass any parameters, use `func` if
you do.**

## The `Header` component

`Header` takes the current URL path, builds one link per entry, and marks the
matching link as active:

```go
type navs struct {
    Name string
    Url  string
}

func Header(currentPage string) string {
    _endpoints := []navs{
        {Url: "about", Name: "About"},
        {Url: "docs", Name: "Documentation"},
        {Url: "samples", Name: "Samples"},
        {Url: "api", Name: "API"},
    }

    if strings.HasPrefix(currentPage, "/") {
        currentPage = currentPage[1:]
    }

    navigation := ""
    for _, endpoint := range _endpoints {
        if endpoint.Url == currentPage {
            navigation += utils.HTParser(`
                <a class="border-b-2 border-b-white border-b-solid px-2" href="/{url}">{name}</a>
            `, utils.PARAMS{"url": endpoint.Url, "name": endpoint.Name})
        } else {
            navigation += utils.HTParser(`
                <a class="px-2" href="/{url}">{name}</a>
            `, utils.PARAMS{"url": endpoint.Url, "name": endpoint.Name})
        }
    }

    return utils.HTParser(`
        <nav class="flex flex-col items-start justify-between gap-4 pt-4 pb-8 text-sm text-slate-400 md:flex-row md:items-center [&_a:hover]:text-slate-50">
            <a href="/">← Back to main</a>
            <div class="flex flex-wrap gap-5">
                {navigation}
            </div>
        </nav>
    `, utils.PARAMS{"navigation": navigation})
}
```

Two things are worth noting:

- The `{key}` tokens are filled with `utils.HTParser` and a `utils.PARAMS` map,
  the same helper routes use (see [rendering](rendering.md)).
- `Header` returns a `string`, so its output can itself be inserted into another
  `HTParser` call as a parameter. That is how components compose.

## Calling a component from a route

Get the current path from Gin, pass it to the component, and inject the returned
HTML through a parameter. This is exactly what `src/routes/about.go` does:

```go
func about(ctx *gin.Context) (string, utils.PARAMS) {
    currentPage := ctx.Request.URL.Path

    return utils.HTParser(`
        <div class="mx-auto max-w-[1120px] pt-8 pb-20">
            {navigation}
            <h1>About Constellation</h1>
        </div>
    `, utils.PARAMS{
        "navigation": components.Header(currentPage),
    }), nil
}
```

The steps are always the same:

1. Import `components`.
2. Call the component, passing whatever data it needs.
3. Put the result in the parameter map of the page's `HTParser` call under a
   token such as `{navigation}`.

If a handler wants to add data of its own, keep adding keys to the same map:

```go
return utils.HTParser(body, utils.PARAMS{
    "navigation": components.Header(currentPage),
    "year":       2026,
}), nil
```

## Passing data to a component

Give the function the parameters it needs and keep the markup inside the
component. A small card component:

```go
func Card(title, summary string) string {
    return utils.HTParser(`
        <div class="rounded border border-white/10 p-4">
            <h3 class="m-0">{title}</h3>
            <p class="m-0 text-slate-400">{summary}</p>
        </div>
    `, utils.PARAMS{
        "title":   title,
        "summary": summary,
    })
}
```

Call it while building a page:

```go
cards := ""
for _, item := range items {
    cards += components.Card(item.Title, item.Summary)
}

return utils.HTParser(`<section>{cards}</section>`, utils.PARAMS{
    "cards": cards,
}), nil
```

## Composing components

Because every component returns a `string`, nesting is just string
concatenation or another `HTParser` call:

```go
func Layout(currentPage, content string) string {
    return utils.HTParser(`
        {header}
        <main>{content}</main>
        {footer}
    `, utils.PARAMS{
        "header":  Header(currentPage),
        "content": content,
        "footer":  Footer,
    })
}
```

## Guidelines

- **Produce a `string`.** A variable holds the markup directly; a function
  returns it. Components do not write to `ctx` or register routes.
- **One concern per component.** `Header` owns navigation; `Footer` owns the
  footer. Keep page-specific markup in the route.
- **Use `HTParser` when the component needs substitution.** A plain `var` can
  hold raw HTML as-is; a `func` that fills `{tokens}` uses `HTParser`.
- **Escape user input.** Never pass untrusted values straight into `HTParser`.
  Use `utils.Parser` (escaped) for text and user input, or escape with
  `html.EscapeString` before the component inserts it. Route parameters returned
  to `Render` are escaped automatically because `Render` uses `Parser`.
- **Accept data, not globals.** Pass values such as the current path as
  arguments instead of reading request state inside the component.

## Adding a component

1. Create `src/components/<name>.go` with `package components`.
2. Export a variable (static markup) or a function (takes parameters, returns a
   string).
3. Import `src/components` in the route and insert the value into its
   `HTParser` parameters.

No registration or configuration is required; the import is enough.
