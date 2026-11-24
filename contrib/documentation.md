# Documentation workflow

There are two Markdown trees:

- `docs/` — user guides, published by the web site at `/docs`.
- `contrib/` — contributor notes, kept in the repository only.

## Adding a user guide

1. Create `docs/<slug>.md`. The filename becomes `/docs/<slug>`.
2. Start the file with a single `# Title` heading. It becomes the page title.
3. Put a short plain paragraph after the heading. It becomes the card summary.
4. Link other guides with relative `.md` links, for example
   `[Styling](styling.md)`. They are rewritten to `/docs/styling`.
5. Rebuild and restart. The docs are embedded at build time.

No code generation step is needed. `docs.Register` in `lib/docs/web.go` creates
the endpoints from the embedded files.

## Supported Markdown

The renderer in `lib/docs/markdown.go` supports:

- headings `#` through `######`
- paragraphs and horizontal rules
- fenced code blocks
- unordered and ordered lists
- blockquotes
- tables with a header separator row
- inline code, bold, emphasis, and links

Keep guides short. The `.markdown` styles in `static/styles.css` handle the
rendered HTML.

## Adding contributor notes

Add a Markdown file under `contrib/` and link it from `contrib/README.md`.
These files are not embedded or published; they describe how to work on the
framework itself.
