package routes

import (
	"github.com/RyannKim327/constellation/config/middleware"
	"github.com/RyannKim327/constellation/config/utils"
	"github.com/RyannKim327/constellation/src/components"

	"github.com/gin-gonic/gin"
)

// Samples is the landing page for the example routes. The examples themselves
// live in sample_json.go and sample_form.go.
var Samples = middleware.Route{
	Method:     utils.METHOD_GET,
	Path:       "/samples",
	Handler:    samples,
	Permission: utils.PERMISSION_ALL,
	Title:      "Samples · Constellation",
}

func samples(ctx *gin.Context) (string, utils.PARAMS) {
	currentPage := ctx.Request.URL.Path

	return utils.HTParser(`
		<div class="mx-auto w-[calc(100%_-_2rem)] max-w-[1120px] pt-8 pb-20 md:w-[calc(100%_-_3rem)]">
			{navigation}

			<header class="pt-12 pb-10">
				<span class="block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400">Sample files</span>
				<h1 class="mt-3 mb-4 text-[clamp(2rem,5vw,3.5rem)] leading-[1.08] tracking-[-0.035em]">Small routes, real behavior</h1>
				<p class="m-0 max-w-[48rem] text-[1.02rem] text-slate-400">
					These examples show the two handler styles: write the response yourself,
					or return HTML and let Constellation render the page.
				</p>
			</header>

			<section class="mb-10 grid grid-cols-[repeat(auto-fit,minmax(17rem,1fr))] gap-4">
				<a href="/samples/json" class="flex flex-col gap-2 rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-[1.4rem] transition duration-150 hover:-translate-y-0.5 hover:border-orange-400/50">
					<span class="text-[1.05rem] font-bold">JSON endpoint</span>
					<span class="text-[0.88rem] text-slate-400">
						A handler that writes JSON and returns an empty body.
					</span>
					<span class="mt-auto pt-2 text-[0.8rem] font-semibold text-orange-400">Open /samples/json →</span>
				</a>
				<a href="/samples/form" class="flex flex-col gap-2 rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-[1.4rem] transition duration-150 hover:-translate-y-0.5 hover:border-orange-400/50">
					<span class="text-[1.05rem] font-bold">Form page</span>
					<span class="text-[0.88rem] text-slate-400">
						A GET form and a POST handler that renders escaped user input.
					</span>
					<span class="mt-auto pt-2 text-[0.8rem] font-semibold text-orange-400">Open /samples/form →</span>
				</a>
				<a href="/api" class="flex flex-col gap-2 rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-[1.4rem] transition duration-150 hover:-translate-y-0.5 hover:border-orange-400/50">
					<span class="text-[1.05rem] font-bold">Backend API</span>
					<span class="text-[0.88rem] text-slate-400">
						The baseline JSON route registered from the backend package.
					</span>
					<span class="mt-auto pt-2 text-[0.8rem] font-semibold text-orange-400">Open /api →</span>
				</a>
			</section>

			<section>
				<h2 class="mt-0 mb-4 text-[1.45rem] tracking-[-0.02em]">Handler contract</h2>
				<p class="mb-4 text-slate-400">
					Every handler returns <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">(string, utils.PARAMS)</code>. Returning
					<code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">("", nil)</code> tells the framework that the handler already wrote
					the response. Any other result is wrapped in the HTML document shell.
				</p>
			</section>

			<footer class="mt-16 flex flex-col items-start justify-between gap-4 border-t border-white/10 pt-6 text-[0.8rem] text-slate-400 md:flex-row md:items-center">
				<span>Constellation</span>
				<span>Examples live in src/routes/sample_*.go</span>
			</footer>
		</div>
	`, utils.PARAMS{
		"navigation": components.Header(currentPage),
	}), nil
}
