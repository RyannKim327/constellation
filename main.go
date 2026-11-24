package main

import (
	"embed"

	constellation "constellation/lib"
	"constellation/lib/docs"
	api "constellation/src/api"
	web "constellation/src/web"
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

	constellation.AddEndpoint(web.Index)
	constellation.AddEndpoint(api.SampleApi)

	css, err := StaticFS.ReadFile("static/styles.css")
	if err != nil {
		panic(err)
	}

	constellation.StartServer(css)
}
