package routes

import (
	"constellation/config/middleware"
	"constellation/config/utils"

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
		<div class="page-shell">
			<nav class="page-nav">
				<a href="/">← Back to main</a>
				<div class="page-nav-links">
					<a href="/docs">Documentation</a>
					<a href="/samples">Samples</a>
					<a href="/api">API</a>
				</div>
			</nav>

			<header class="page-header">
				<span class="eyebrow">About the project</span>
				<h1 class="page-title">What is Constellation?</h1>
				<p class="page-lead">
					Constellation is a small, convention-over-configuration web framework for
					Go, built on top of <a class="text-link" href="https://github.com/gin-gonic/gin">Gin</a>.
					It was created to explore how far a declarative route model and a tiny
					templating helper can go.
				</p>
			</header>

			<div class="content-grid">
				<section class="prose-block">
					<h2>Why it exists</h2>
					<p>
						The project borrows the component-oriented feel of React and applies it
						to server-rendered Go pages. Each route returns a body and a parameter
						map; the framework renders the document. JSON endpoints use the same
						signature and simply write their own response.
					</p>

					<h2>What is included</h2>
					<ul class="plain-list">
						<li>Declarative routes with automatic method and path normalization.</li>
						<li>CORS setup driven by environment variables.</li>
						<li>Bot and scraper detection on every registered route.</li>
						<li>Cookie-based access control for protected submissions.</li>
						<li>Request logging, custom 404 pages, and static file serving.</li>
						<li>Markdown documentation converted into Go routes.</li>
					</ul>

					<h2>How it is organized</h2>
					<p>
						<code>config/</code> is the framework: bootstrap, middleware, route
						registration, rendering, and server lifecycle. <code>src/</code> is the
						application: the entrypoint, routes, and static assets. The documentation
						in <code>docs/</code> is generated into <code>src/routes/docs_*.go</code>
						by <code>tools/md2go</code>.
					</p>

					<div class="callout">
						<p class="callout-title">Try it locally</p>
						<p class="callout-body">
							Run <code>go run ./src/index.go</code> and open
							<a class="text-link" href="http://localhost:8000">localhost:8000</a>.
							For live reload, run <code>air</code>.
						</p>
					</div>
				</section>

				<aside class="fact-panel">
					<p class="section-label">At a glance</p>
					<dl>
						<div><dt>Language</dt><dd>Go</dd></div>
						<div><dt>Router</dt><dd>Gin</dd></div>
						<div><dt>Rendering</dt><dd>HTML + Markdown</dd></div>
						<div><dt>Default port</dt><dd>8000</dd></div>
						<div><dt>Live reload</dt><dd>Air</dd></div>
					</dl>
					<a class="button button-primary" href="/docs">Read the docs</a>
				</aside>
			</div>
		</div>
	`, nil
}
