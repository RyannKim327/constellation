package main

import (
	routes "github.com/RyannKim327/constellation/src/routes"

	c "github.com/RyannKim327/constellation"
	mw "github.com/RyannKim327/constellation/config/middleware"
	"github.com/RyannKim327/constellation/src/backend"
)

// TODO: Collect all the Routes here
// NOTE: I suggest to use routes (folder) to make it organize
var _routes = append([]mw.Route{
	backend.API,
	backend.SampleJSON,

	routes.Index,
	routes.About,
	routes.Docs,
	routes.Samples,
	routes.SampleForm,
	routes.SampleFormSubmit,
}, routes.DocsPages...)

func main() {
	// TODO: This code will going to initiate everything
	c.Constellation(_routes)
}
