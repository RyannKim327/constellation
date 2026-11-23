package main

import (
	"embed"

	sample_api "constellation/sample/api"
	sample_web "constellation/sample/web"
	constellation "constellation/src"
	"constellation/src/docs"
)

// StaticFS contains the assets under static, embedded at build time.
//
//go:embed all:static
var StaticFS embed.FS

// DocsFS contains the Markdown documentation, embedded at build time.
//
//go:embed docs
var DocsFS embed.FS

func main() {
	docs.Register(DocsFS)

	constellation.AddEndpoint(sample_web.Index)
	constellation.AddEndpoint(sample_api.SampleApi)

	css, err := StaticFS.ReadFile("static/styles.css")
	if err != nil {
		panic(err)
	}

	constellation.StartServer(css)
}
