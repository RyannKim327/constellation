package routes

import (
	"github.com/RyannKim327/constellation/config/middleware"
	"github.com/RyannKim327/constellation/config/utils"

	"github.com/gin-gonic/gin"
)

// About describes the project, its goals, and how the repository is organized.
var About = middleware.Route{
	Method:     utils.METHOD_GET,
	Path:       "/about",
	Handler:    about,
	Permission: utils.PERMISSION_ALL,
	Title:      "About · Constellation",
}

func about(_ *gin.Context) (string, utils.PARAMS) {
	return `
		<div class="mx-auto w-[calc(100%_-_2rem)] max-w-[1120px] pt-8 pb-20 md:w-[calc(100%_-_3rem)]">
			<nav class="flex flex-col items-start justify-between gap-4 pt-4 pb-8 text-sm text-slate-400 md:flex-row md:items-center [&_a:hover]:text-slate-50">
				<a href="/">← Back to main</a>
				<div class="flex flex-wrap gap-5">
					<a href="/docs">Documentation</a>
					<a href="/samples">Samples</a>
					<a href="/api">API</a>
				</div>
			</nav>

			<header class="pt-12 pb-10">
				<span class="block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400">About the project</span>
				<h1 class="mt-3 mb-4 text-[clamp(2rem,5vw,3.5rem)] leading-[1.08] tracking-[-0.035em]">What is Constellation?</h1>
				<p class="m-0 max-w-[48rem] text-[1.02rem] text-slate-400">
					Constellation is a small, convention-over-configuration web framework for
					Go, built on top of <a class="font-semibold text-orange-400 hover:underline" href="https://github.com/gin-gonic/gin">Gin</a>.
					It was created to explore how far a declarative route model and a tiny
					templating helper can go.
				</p>
			</header>

			<div class="grid grid-cols-1 items-start gap-12 lg:grid-cols-[minmax(0,2fr)_minmax(16rem,1fr)]">
				<section>
					<h2 class="mt-0 mb-4 text-[1.45rem] tracking-[-0.02em]">Why it exists</h2>
					<p class="mb-4 text-slate-400">
						The project borrows the component-oriented feel of React and applies it
						to server-rendered Go pages. Each route returns a body and a parameter
						map; the framework renders the document. JSON endpoints use the same
						signature and simply write their own response.
					</p>

					<h2 class="mt-9 mb-4 text-[1.45rem] tracking-[-0.02em]">What is included</h2>
					<ul class="mb-6 list-disc pl-5 text-slate-400">
						<li class="mb-2">Declarative routes with automatic method and path normalization.</li>
						<li class="mb-2">CORS setup driven by environment variables.</li>
						<li class="mb-2">Bot and scraper detection on every registered route.</li>
						<li class="mb-2">Cookie-based access control for protected submissions.</li>
						<li class="mb-2">Request logging, custom 404 pages, and static file serving.</li>
						<li class="mb-2">Markdown documentation converted into Go routes.</li>
					</ul>

					<h2 class="mt-9 mb-4 text-[1.45rem] tracking-[-0.02em]">How it is organized</h2>
					<p class="mb-4 text-slate-400">
						<code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">config/</code> is the framework: bootstrap, middleware, route
						registration, rendering, and server lifecycle. <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">src/</code> is the
						application: the entrypoint, routes, and static assets. The documentation
						in <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">docs/</code> is generated into <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">src/routes/docs_*.go</code>
						by <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">tools/md2go</code>.
					</p>

					<div class="rounded-[0.9rem] border border-orange-400/30 bg-orange-400/[0.07] px-[1.4rem] py-[1.2rem]">
						<p class="m-0 mb-[0.35rem] text-[0.8rem] font-bold uppercase tracking-[0.08em] text-orange-400">Try it locally</p>
						<p class="m-0 text-[0.9rem] text-slate-400">
							Run <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">go run ./src/index.go</code> and open
							<a class="font-semibold text-orange-400 hover:underline" href="http://localhost:8000">localhost:8000</a>.
							For live reload, run <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">air</code>.
						</p>
					</div>
				</section>

				<aside class="rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-6">
					<p class="mb-3 block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400">At a glance</p>
					<dl class="my-4 mb-6">
						<div class="flex items-center justify-between gap-4 border-b border-white/10 py-[0.6rem] text-[0.85rem]"><dt class="text-slate-400">Language</dt><dd class="m-0 font-semibold">Go</dd></div>
						<div class="flex items-center justify-between gap-4 border-b border-white/10 py-[0.6rem] text-[0.85rem]"><dt class="text-slate-400">Router</dt><dd class="m-0 font-semibold">Gin</dd></div>
						<div class="flex items-center justify-between gap-4 border-b border-white/10 py-[0.6rem] text-[0.85rem]"><dt class="text-slate-400">Rendering</dt><dd class="m-0 font-semibold">HTML + Markdown</dd></div>
						<div class="flex items-center justify-between gap-4 border-b border-white/10 py-[0.6rem] text-[0.85rem]"><dt class="text-slate-400">Default port</dt><dd class="m-0 font-semibold">8000</dd></div>
						<div class="flex items-center justify-between gap-4 border-b border-white/10 py-[0.6rem] text-[0.85rem]"><dt class="text-slate-400">Live reload</dt><dd class="m-0 font-semibold">Air</dd></div>
					</dl>
					<a class="inline-flex cursor-pointer items-center justify-center rounded-lg border border-transparent bg-orange-500 px-[1.1rem] py-[0.7rem] text-sm font-semibold text-gray-900 transition duration-150 hover:-translate-y-px hover:bg-orange-400" href="/docs">Read the docs</a>
				</aside>
			</div>
		</div>
	`, nil
}
