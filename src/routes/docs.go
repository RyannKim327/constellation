//go:generate go run ../../tools/md2go -docs ../../docs -out .

package routes

import (
	"html"

	"github.com/RyannKim327/constellation/config/middleware"
	"github.com/RyannKim327/constellation/config/utils"
	u "github.com/RyannKim327/constellation/src/utils"
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
		cards += `<a href="/docs/` + item.Slug + `" class="flex flex-col gap-2 rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-[1.4rem] transition duration-150 hover:-translate-y-0.5 hover:border-orange-400/50">
			<span class="text-[1.05rem] font-bold">` + html.EscapeString(item.Title) + `</span>
			<span class="text-[0.88rem] text-slate-400">` + html.EscapeString(item.Summary) + `</span>
			<span class="mt-auto pt-2 text-[0.8rem] font-semibold text-orange-400">Read document →</span>
		</a>`
	}

	body := `<div class="mx-auto w-[calc(100%_-_2rem)] max-w-[1120px] pt-8 pb-20 md:w-[calc(100%_-_3rem)]">
		<nav class="flex flex-col items-start justify-between gap-4 pt-4 pb-8 text-sm text-slate-400 md:flex-row md:items-center [&_a:hover]:text-slate-50">
			<a href="/">← Back to main</a>
			<div class="flex flex-wrap gap-5">
				<a href="/about">About</a>
				<a href="/samples">Samples</a>
				<a href="/api">API</a>
			</div>
		</nav>

		<header class="pt-12 pb-10">
			<span class="block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400">Documentation</span>
			<h1 class="mt-3 mb-4 text-[clamp(2rem,5vw,3.5rem)] leading-[1.08] tracking-[-0.035em]">Constellation handbook</h1>
			<p class="m-0 max-w-[48rem] text-[1.02rem] text-slate-400">
				Every guide below is authored in Markdown under <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">docs/</code> and converted
				into Go files in <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">src/routes/</code>. Edit the Markdown, run
				<code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">go generate ./src/routes</code>, and the site picks up the change.
			</p>
		</header>

		<section class="mb-10 grid grid-cols-[repeat(auto-fit,minmax(17rem,1fr))] gap-4">` + cards + `</section>

		<section class="rounded-[0.9rem] border border-orange-400/30 bg-orange-400/[0.07] px-[1.4rem] py-[1.2rem]">
			<p class="m-0 mb-[0.35rem] text-[0.8rem] font-bold uppercase tracking-[0.08em] text-orange-400">Keeping docs in sync</p>
			<p class="m-0 text-[0.9rem] text-slate-400">
				The generated files use the same package as this route, so documentation pages
				are regular Constellation routes. No files are read at request time.
			</p>
		</section>
	</div>`

	return body, nil
}

func docLayout(doc markdownDoc) string {
	body := `<div class="mx-auto w-[calc(100%_-_2rem)] max-w-[1120px] pt-8 pb-20 md:w-[calc(100%_-_3rem)]">
		<nav class="flex flex-col items-start justify-between gap-4 pt-4 pb-8 text-sm text-slate-400 md:flex-row md:items-center [&_a:hover]:text-slate-50">
			<a href="/docs">← All documentation</a>
			<div class="flex flex-wrap gap-5">
				<a href="/">Home</a>
				<a href="/about">About</a>
				<a href="/samples">Samples</a>
			</div>
		</nav>

		<div class="grid grid-cols-1 items-start gap-12 lg:grid-cols-[minmax(13rem,16rem)_minmax(0,1fr)]">
			<aside class="rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-[1.2rem] lg:sticky lg:top-8">` + docsNavigation(doc.Slug) + `</aside>

			<article class="min-w-0">
				<header class="mb-8 border-b border-white/10 pb-7">
					<span class="block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400">Documentation</span>
					<h1 class="mt-[0.65rem] mb-3 text-[clamp(1.9rem,4vw,2.8rem)] tracking-[-0.03em]">` + html.EscapeString(doc.Title) + `</h1>
					<p class="m-0 max-w-[48rem] text-[1.02rem] text-slate-400">` + html.EscapeString(doc.Summary) + `</p>
				</header>
				<div class="prose prose-invert max-w-none prose-headings:text-slate-50 prose-headings:tracking-[-0.02em] prose-h2:border-b prose-h2:border-white/10 prose-h2:pb-2 prose-p:text-slate-300 prose-li:text-slate-300 prose-a:text-orange-400 prose-strong:text-slate-50 prose-code:text-orange-300 prose-pre:overflow-x-auto prose-pre:rounded-xl prose-pre:border prose-pre:border-white/10 prose-pre:bg-[#07070a] prose-blockquote:border-l-[3px] prose-blockquote:border-orange-400 prose-blockquote:bg-orange-400/[0.06] prose-blockquote:not-italic prose-th:bg-white/5 prose-th:text-slate-50 prose-td:text-slate-300 prose-table:text-[0.85rem]">` + u.Markdown(doc.Markdown) + `</div>
			</article>
		</div>
	</div>`

	return body
}

func docsNavigation(active string) string {
	nav := `<p class="m-0 mb-3 text-[0.7rem] font-bold uppercase tracking-[0.16em] text-slate-400">Documents</p><ul class="m-0 flex list-none flex-col gap-0.5 p-0">`
	for _, item := range generatedDocs {
		class := "block rounded px-[0.6rem] py-[0.45rem] text-[0.85rem] text-slate-400 hover:text-slate-50"
		if item.Slug == active {
			class += " bg-orange-400/[0.12] font-semibold text-orange-400"
		}
		nav += `<li><a class="` + class + `" href="/docs/` + item.Slug + `">` +
			html.EscapeString(item.Title) + `</a></li>`
	}
	nav += `</ul>`
	return nav
}
