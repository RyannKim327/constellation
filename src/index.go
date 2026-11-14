package main

import (
	c "constellation/config"
	mw "constellation/config/middleware"
	"constellation/src/backend"
	routes "constellation/src/routes"
)

// TODO: Collect all the Routes here
// NOTE: I suggest to use routes (folder) to make it organize
var _routes = []mw.Route{
	backend.API,

	routes.Index,
}

func main() {
	// TODO: This code will going to initiate everything
	c.Constellation(_routes)
}
