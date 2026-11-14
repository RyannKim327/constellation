package routes

import (
	"constellation/config/middleware"
	"constellation/config/utils"

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
		<div class="page-shell">
			<nav class="page-nav">
				<a href="/" class="brand">Constellation</a>
				<div class="page-nav-links">
					<a href="/about">About</a>
					<a href="/docs">Documentation</a>
					<a href="/samples">Samples</a>
					<a href="/api">API</a>
				</div>
			</nav>

			<header class="hero">
				<span class="eyebrow">Built with Go and Gin</span>
				<h1 class="hero-title">Small framework.<br />Clear conventions.</h1>
				<p class="hero-lead">
					Constellation is a convention-over-configuration web framework for Go.
					Declare a route, return HTML or JSON, and let the framework handle CORS,
					access control, static files, logging, and the 404 page.
				</p>
				<div class="hero-actions">
					<a class="button button-primary" href="/docs">Read the docs</a>
					<a class="button button-secondary" href="/samples">Browse samples</a>
				</div>
			</header>

			<section class="feature-grid">
				<article class="feature-card">
					<span class="feature-index">01</span>
					<h2>Declarative routes</h2>
					<p>
						Describe an endpoint with a <code>middleware.Route</code> value.
						Methods and paths are normalized automatically.
					</p>
				</article>
				<article class="feature-card">
					<span class="feature-index">02</span>
					<h2>Server-rendered pages</h2>
					<p>
						Compose HTML with <code>{placeholder}</code> templating, or return
						JSON from the same handler signature.
					</p>
				</article>
				<article class="feature-card">
					<span class="feature-index">03</span>
					<h2>Batteries included</h2>
					<p>
						CORS, scraper detection, cookie access control, request logging,
						custom 404s, and static file serving.
					</p>
				</article>
			</section>

			<section class="split-section">
				<div class="prose-block">
					<span class="eyebrow">Documentation</span>
					<h2>Guides generated from Markdown</h2>
					<p>
						The {docs} guides in <code>docs/</code> are converted into Go files in
						<code>src/routes/</code>. The same source renders here as HTML, so the
						repository stays the single source of truth.
					</p>
					<a class="text-link" href="/docs">Open the documentation →</a>
				</div>

				<pre class="code-block"><code>var Hello = middleware.Route{
  Method: utils.METHOD_GET,
  Path:   "/hello",
  Handler: func(ctx *gin.Context) (string, map[string]any) {
    return `+"`"+`<h1>Hello {name}</h1>`+"`"+`,
      map[string]any{"name": "World"}
  },
}</code></pre>
			</section>

			<section class="route-list">
				<p class="section-label">Available pages</p>
				<a href="/about"><span>About the project</span><span>→</span></a>
				<a href="/docs"><span>Documentation index</span><span>→</span></a>
				<a href="/samples"><span>Sample routes</span><span>→</span></a>
				<a href="/api"><span>JSON API</span><span>→</span></a>
			</section>

			<footer class="page-footer">
				<span>Constellation</span>
				<span>Go · Gin · Markdown</span>
			</footer>
		</div>
	`, utils.PARAMS{
		"docs": len(generatedDocs),
	}), nil
}
