package routes

import (
	"github.com/RyannKim327/constellation/config/middleware"
	"github.com/RyannKim327/constellation/config/utils"

	"github.com/gin-gonic/gin"
)

// Index is the main page of the application. It introduces the project and
// links to the about, documentation, and sample routes.
var Index = middleware.Route{
	Method:     utils.METHOD_GET,
	Path:       "/",
	Handler:    layout,
	Permission: utils.PERMISSION_ALL,
	Title:      "Constellation · A small web framework for Go",
}

func layout(_ *gin.Context) (string, utils.PARAMS) {
	return utils.HTParser(`
		<div class="mx-auto w-[calc(100%_-_2rem)] max-w-[1120px] pt-8 pb-20 md:w-[calc(100%_-_3rem)]">
			<nav class="flex flex-col items-start justify-between gap-4 pt-4 pb-8 text-sm text-slate-400 md:flex-row md:items-center [&_a:hover]:text-slate-50">
				<a href="/" class="text-base font-bold uppercase tracking-[0.08em] text-slate-50">Constellation</a>
				<div class="flex flex-wrap gap-5">
					<a href="/about">About</a>
					<a href="/docs">Documentation</a>
					<a href="/samples">Samples</a>
					<a href="/api">API</a>
				</div>
			</nav>

			<header class="flex flex-col items-start pt-20 pb-16">
				<span class="block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400">Built with Go and Gin</span>
				<h1 class="my-5 mb-6 text-[clamp(2.75rem,7vw,5.5rem)] font-bold leading-[0.98] tracking-[-0.045em]">Small framework.<br />Clear conventions.</h1>
				<p class="mb-8 max-w-2xl text-[clamp(1rem,2vw,1.2rem)] text-slate-400">
					Constellation is a convention-over-configuration web framework for Go.
					Declare a route, return HTML or JSON, and let the framework handle CORS,
					access control, static files, logging, and the 404 page.
				</p>
				<div class="flex flex-wrap gap-3">
					<a class="inline-flex cursor-pointer items-center justify-center rounded-lg border border-transparent bg-orange-500 px-[1.1rem] py-[0.7rem] text-sm font-semibold text-gray-900 transition duration-150 hover:-translate-y-px hover:bg-orange-400" href="/docs">Read the docs</a>
					<a class="inline-flex cursor-pointer items-center justify-center rounded-lg border border-white/10 bg-transparent px-[1.1rem] py-[0.7rem] text-sm font-semibold text-slate-50 transition duration-150 hover:-translate-y-px hover:border-orange-400" href="/samples">Browse samples</a>
				</div>
			</header>

			<section class="mb-16 grid grid-cols-[repeat(auto-fit,minmax(15rem,1fr))] gap-4">
				<article class="rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-6">
					<span class="text-xs font-bold tracking-[0.12em] text-orange-400">01</span>
					<h2 class="mt-5 mb-2 text-[1.1rem]">Declarative routes</h2>
					<p class="m-0 text-[0.9rem] text-slate-400">
						Describe an endpoint with a <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">middleware.Route</code> value.
						Methods and paths are normalized automatically.
					</p>
				</article>
				<article class="rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-6">
					<span class="text-xs font-bold tracking-[0.12em] text-orange-400">02</span>
					<h2 class="mt-5 mb-2 text-[1.1rem]">Server-rendered pages</h2>
					<p class="m-0 text-[0.9rem] text-slate-400">
						Compose HTML with <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">{placeholder}</code> templating, or return
						JSON from the same handler signature.
					</p>
				</article>
				<article class="rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-6">
					<span class="text-xs font-bold tracking-[0.12em] text-orange-400">03</span>
					<h2 class="mt-5 mb-2 text-[1.1rem]">Batteries included</h2>
					<p class="m-0 text-[0.9rem] text-slate-400">
						CORS, scraper detection, cookie access control, request logging,
						custom 404s, and static file serving.
					</p>
				</article>
			</section>

			<section class="mb-12 grid grid-cols-1 items-center gap-8 border-y border-white/10 py-14 md:grid-cols-2">
				<div>
					<span class="block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400">Documentation</span>
					<h2 class="mt-[0.65rem] mb-4 text-[clamp(1.5rem,3vw,2.1rem)] tracking-[-0.02em]">Guides generated from Markdown</h2>
					<p class="mb-4 text-slate-400">
						The {docs} guides in <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">docs/</code> are converted into Go files in
						<code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">src/routes/</code>. The same source renders here as HTML, so the
						repository stays the single source of truth.
					</p>
					<a class="font-semibold text-orange-400 hover:underline" href="/docs">Open the documentation →</a>
				</div>

				<pre class="overflow-x-auto rounded-[0.9rem] border border-white/10 bg-[#07070a] p-[1.4rem] text-[0.82rem] leading-[1.65] text-slate-200"><code>var Hello = middleware.Route{
  Method: utils.METHOD_GET,
  Path:   "/hello",
  Handler: func(ctx *gin.Context) (string, map[string]any) {
    return `+"`"+`<h1>Hello {name}</h1>`+"`"+`,
      map[string]any{"name": "World"}
  },
}</code></pre>
			</section>

			<section class="flex flex-col">
				<p class="mb-3 block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400">Available pages</p>
				<a class="flex items-center justify-between border-b border-white/10 py-4 text-[0.95rem] text-slate-50 hover:text-orange-400" href="/about"><span>About the project</span><span>→</span></a>
				<a class="flex items-center justify-between border-b border-white/10 py-4 text-[0.95rem] text-slate-50 hover:text-orange-400" href="/docs"><span>Documentation index</span><span>→</span></a>
				<a class="flex items-center justify-between border-b border-white/10 py-4 text-[0.95rem] text-slate-50 hover:text-orange-400" href="/samples"><span>Sample routes</span><span>→</span></a>
				<a class="flex items-center justify-between border-b border-white/10 py-4 text-[0.95rem] text-slate-50 hover:text-orange-400" href="/api"><span>JSON API</span><span>→</span></a>
			</section>

			<footer class="mt-16 flex flex-col items-start justify-between gap-4 border-t border-white/10 pt-6 text-[0.8rem] text-slate-400 md:flex-row md:items-center">
				<span>Constellation</span>
				<span>Go · Gin · Markdown</span>
			</footer>
		</div>
	`, utils.PARAMS{
		"docs": len(generatedDocs),
	}), nil
}
