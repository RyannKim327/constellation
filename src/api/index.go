package api

import (
	constellation "constellation/lib"
	"constellation/lib/middleware"
)

var SampleApi = constellation.Endpoint{
	Path:       "/api",
	Method:     constellation.GET,
	Permission: constellation.Unauthenticated,
	Handler: constellation.HandlerSetup{
		Api: api,
	},
}

func api(r *constellation.Request, w constellation.Response) {
	middleware.SendApi(w, map[string]any{
		"test": "Mode",
	})
}
