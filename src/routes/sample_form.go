package routes

import (
	"github.com/RyannKim327/constellation/config/middleware"
	"github.com/RyannKim327/constellation/config/utils"

	"github.com/gin-gonic/gin"
)

// SampleForm renders a small form; SampleFormSubmit handles the POST and shows
// the submitted value through the escaped parameter map.
var (
	SampleForm = middleware.Route{
		Method:     utils.METHOD_GET,
		Path:       "/samples/form",
		Handler:    sampleForm,
		Permission: utils.PERMISSION_ALL,
		Title:      "Form sample · Constellation",
	}

	SampleFormSubmit = middleware.Route{
		Method:     utils.METHOD_POST,
		Path:       "/samples/form/submit",
		Handler:    sampleFormSubmit,
		Permission: utils.PERMISSION_ALL,
		Title:      "Form result · Constellation",
	}
)

func sampleForm(_ *gin.Context) (string, utils.PARAMS) {
	return `
		<div class="mx-auto w-[calc(100%_-_2rem)] max-w-[1120px] pt-8 pb-20 md:w-[calc(100%_-_3rem)]">
			<nav class="flex flex-col items-start justify-between gap-4 pt-4 pb-8 text-sm text-slate-400 md:flex-row md:items-center [&_a:hover]:text-slate-50">
				<a href="/samples">← Back to samples</a>
				<div class="flex flex-wrap gap-5">
					<a href="/">Home</a>
					<a href="/docs">Documentation</a>
				</div>
			</nav>

			<header class="pt-12 pb-10">
				<span class="block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400">Sample · GET + POST</span>
				<h1 class="mt-3 mb-4 text-[clamp(2rem,5vw,3.5rem)] leading-[1.08] tracking-[-0.035em]">Form submission</h1>
				<p class="m-0 max-w-[48rem] text-[1.02rem] text-slate-400">
					The form posts to <code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">/samples/form/submit</code>. The handler returns the
					name as a parameter, so it is HTML-escaped before rendering.
				</p>
			</header>

			<form class="flex max-w-md flex-col gap-3 rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-6" method="post" action="/samples/form/submit">
				<label class="text-[0.8rem] font-semibold uppercase tracking-[0.06em] text-slate-400" for="name">Your name</label>
				<input class="rounded-lg border border-white/10 bg-[#07070a] px-[0.8rem] py-[0.7rem] text-[0.95rem] text-slate-50 outline-none focus:border-orange-400" id="name" name="name" type="text" placeholder="Ada Lovelace" required />
				<button class="mt-1 inline-flex cursor-pointer self-start items-center justify-center rounded-lg border border-transparent bg-orange-500 px-[1.1rem] py-[0.7rem] text-sm font-semibold text-gray-900 transition duration-150 hover:-translate-y-px hover:bg-orange-400" type="submit">Send</button>
			</form>
		</div>
	`, nil
}

func sampleFormSubmit(ctx *gin.Context) (string, utils.PARAMS) {
	name := ctx.PostForm("name")
	if name == "" {
		name = "Anonymous"
	}

	return `
		<div class="mx-auto w-[calc(100%_-_2rem)] max-w-[1120px] pt-8 pb-20 md:w-[calc(100%_-_3rem)]">
			<nav class="flex flex-col items-start justify-between gap-4 pt-4 pb-8 text-sm text-slate-400 md:flex-row md:items-center [&_a:hover]:text-slate-50">
				<a href="/samples/form">← Try again</a>
				<div class="flex flex-wrap gap-5">
					<a href="/samples">Samples</a>
					<a href="/docs">Documentation</a>
				</div>
			</nav>

			<header class="pt-12 pb-10">
				<span class="block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400">Sample · POST result</span>
				<h1 class="mt-3 mb-4 text-[clamp(2rem,5vw,3.5rem)] leading-[1.08] tracking-[-0.035em]">Hello, {name}</h1>
				<p class="m-0 max-w-[48rem] text-[1.02rem] text-slate-400">
					This value arrived through a form field and was escaped by
					<code class="rounded border border-white/10 bg-white/[0.06] px-[0.35rem] py-[0.1rem] text-[0.85em] text-orange-300">utils.Render</code>.
				</p>
			</header>

			<a class="inline-flex cursor-pointer items-center justify-center rounded-lg border border-white/10 bg-transparent px-[1.1rem] py-[0.7rem] text-sm font-semibold text-slate-50 transition duration-150 hover:-translate-y-px hover:border-orange-400" href="/samples/form">Submit another</a>
		</div>
	`, utils.PARAMS{
		"name": name,
	}
}
