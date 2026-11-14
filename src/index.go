package main

import (
	c "constellation"
	mw "constellation/config/middleware"
	"constellation/src/backend"
	routes "constellation/src/routes"
)

// TODO: Collect all the Routes here
// NOTE: I suggest to use routes (folder) to make it organize
var _routes = append([]mw.Route{
	backend.API,

	routes.Index,
	routes.About,
	routes.Docs,
	routes.Samples,
	routes.SampleJSON,
	routes.SampleForm,
	routes.SampleFormSubmit,
}, routes.DocsPages...)

func main() {
	// TODO: This code will going to initiate everything
	c.Constellation(_routes)
}
