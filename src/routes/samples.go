package routes

import (
	"constellation/config/middleware"
	"constellation/config/utils"

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

func samples(_ *gin.Context) (string, utils.PARAMS) {
	return `
		<div class="page-shell">
			<nav class="page-nav">
				<a href="/">← Back to main</a>
				<div class="page-nav-links">
					<a href="/about">About</a>
					<a href="/docs">Documentation</a>
					<a href="/api">API</a>
				</div>
			</nav>

			<header class="page-header">
				<span class="eyebrow">Sample files</span>
				<h1 class="page-title">Small routes, real behavior</h1>
				<p class="page-lead">
					These examples show the two handler styles: write the response yourself,
					or return HTML and let Constellation render the page.
				</p>
			</header>

			<section class="docs-grid">
				<a href="/samples/json" class="docs-card">
					<span class="docs-card-title">JSON endpoint</span>
					<span class="docs-card-summary">
						A handler that writes JSON and returns an empty body.
					</span>
					<span class="docs-card-link">Open /samples/json →</span>
				</a>
				<a href="/samples/form" class="docs-card">
					<span class="docs-card-title">Form page</span>
					<span class="docs-card-summary">
						A GET form and a POST handler that renders escaped user input.
					</span>
					<span class="docs-card-link">Open /samples/form →</span>
				</a>
				<a href="/api" class="docs-card">
					<span class="docs-card-title">Backend API</span>
					<span class="docs-card-summary">
						The baseline JSON route registered from the backend package.
					</span>
					<span class="docs-card-link">Open /api →</span>
				</a>
			</section>

			<section class="prose-block">
				<h2>Handler contract</h2>
				<p>
					Every handler returns <code>(string, utils.PARAMS)</code>. Returning
					<code>("", nil)</code> tells the framework that the handler already wrote
					the response. Any other result is wrapped in the HTML document shell.
				</p>
			</section>

			<footer class="page-footer">
				<span>Constellation</span>
				<span>Examples live in src/routes/sample_*.go</span>
			</footer>
		</div>
	`, nil
}
