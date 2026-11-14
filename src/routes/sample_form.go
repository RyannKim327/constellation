package routes

import (
	"constellation/config/middleware"
	"constellation/config/utils"

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
		<div class="page-shell">
			<nav class="page-nav">
				<a href="/samples">← Back to samples</a>
				<div class="page-nav-links">
					<a href="/">Home</a>
					<a href="/docs">Documentation</a>
				</div>
			</nav>

			<header class="page-header">
				<span class="eyebrow">Sample · GET + POST</span>
				<h1 class="page-title">Form submission</h1>
				<p class="page-lead">
					The form posts to <code>/samples/form/submit</code>. The handler returns the
					name as a parameter, so it is HTML-escaped before rendering.
				</p>
			</header>

			<form class="sample-form" method="post" action="/samples/form/submit">
				<label for="name">Your name</label>
				<input id="name" name="name" type="text" placeholder="Ada Lovelace" required />
				<button class="button button-primary" type="submit">Send</button>
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
		<div class="page-shell">
			<nav class="page-nav">
				<a href="/samples/form">← Try again</a>
				<div class="page-nav-links">
					<a href="/samples">Samples</a>
					<a href="/docs">Documentation</a>
				</div>
			</nav>

			<header class="page-header">
				<span class="eyebrow">Sample · POST result</span>
				<h1 class="page-title">Hello, {name}</h1>
				<p class="page-lead">
					This value arrived through a form field and was escaped by
					<code>utils.Render</code>.
				</p>
			</header>

			<a class="button button-secondary" href="/samples/form">Submit another</a>
		</div>
	`, utils.PARAMS{
		"name": name,
	}
}
