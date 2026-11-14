//go:generate go run ../../tools/md2go -docs ../../docs -out .

package routes

import (
	"html"

	"constellation/config/middleware"
	"constellation/config/utils"

	"github.com/gin-gonic/gin"
)

// markdownDoc is the shape emitted by tools/md2go for each Markdown file in
// docs/. See docs_manifest.go for the generated list.
type markdownDoc struct {
	Slug     string
	Title    string
	Summary  string
	Markdown string
}

// Docs is the documentation landing page. Individual documents are exposed by
// DocsPages at /docs/<slug>.
var Docs = middleware.Route{
	Method:  utils.METHOD_GET,
	Path:    "/docs",
	Title:   "Documentation · Constellation",
	Handler: docsIndex,
}

// DocsPages is built from the generated documentation manifest. Register it
// alongside the other frontend routes in src/index.go.
var DocsPages = buildDocsPages()

func buildDocsPages() []middleware.Route {
	pages := make([]middleware.Route, 0, len(generatedDocs))

	for _, item := range generatedDocs {
		item := item
		pages = append(pages, middleware.Route{
			Method: utils.METHOD_GET,
			Path:   "/docs/" + item.Slug,
			Title:  item.Title + " · Constellation",
			Handler: func(_ *gin.Context) (string, utils.PARAMS) {
				return docLayout(item), nil
			},
		})
	}

	return pages
}

func docsIndex(_ *gin.Context) (string, utils.PARAMS) {
	cards := ""
	for _, item := range generatedDocs {
		cards += `<a href="/docs/` + item.Slug + `" class="docs-card">
			<span class="docs-card-title">` + html.EscapeString(item.Title) + `</span>
			<span class="docs-card-summary">` + html.EscapeString(item.Summary) + `</span>
			<span class="docs-card-link">Read document →</span>
		</a>`
	}

	body := `<div class="page-shell">
		<nav class="page-nav">
			<a href="/">← Back to main</a>
			<div class="page-nav-links">
				<a href="/about">About</a>
				<a href="/samples">Samples</a>
				<a href="/api">API</a>
			</div>
		</nav>

		<header class="page-header">
			<span class="eyebrow">Documentation</span>
			<h1 class="page-title">Constellation handbook</h1>
			<p class="page-lead">
				Every guide below is authored in Markdown under <code>docs/</code> and converted
				into Go files in <code>src/routes/</code>. Edit the Markdown, run
				<code>go generate ./src/routes</code>, and the site picks up the change.
			</p>
		</header>

		<section class="docs-grid">` + cards + `</section>

		<section class="callout">
			<p class="callout-title">Keeping docs in sync</p>
			<p class="callout-body">
				The generated files use the same package as this route, so documentation pages
				are regular Constellation routes. No files are read at request time.
			</p>
		</section>
	</div>`

	return body, nil
}

func docLayout(doc markdownDoc) string {
	body := `<div class="page-shell">
		<nav class="page-nav">
			<a href="/docs">← All documentation</a>
			<div class="page-nav-links">
				<a href="/">Home</a>
				<a href="/about">About</a>
				<a href="/samples">Samples</a>
			</div>
		</nav>

		<div class="docs-layout">
			<aside class="docs-sidebar">` + docsNavigation(doc.Slug) + `</aside>

			<article class="docs-article">
				<header class="docs-article-header">
					<span class="eyebrow">Documentation</span>
					<h1>` + html.EscapeString(doc.Title) + `</h1>
					<p class="page-lead">` + html.EscapeString(doc.Summary) + `</p>
				</header>
				<div class="markdown">` + utils.Markdown(doc.Markdown) + `</div>
			</article>
		</div>
	</div>`

	return body
}

func docsNavigation(active string) string {
	nav := `<p class="docs-sidebar-title">Documents</p><ul class="docs-sidebar-list">`
	for _, item := range generatedDocs {
		class := "docs-sidebar-link"
		if item.Slug == active {
			class += " docs-sidebar-link-active"
		}
		nav += `<li><a class="` + class + `" href="/docs/` + item.Slug + `">` +
			html.EscapeString(item.Title) + `</a></li>`
	}
	nav += `</ul>`
	return nav
}
