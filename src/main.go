package constellation

import (
	"fmt"
	"net/http"

	"constellation/src/middleware"
)

func StartServer(css []byte) {
	serve := http.NewServeMux()

	for _, endpoint := range Endpoints {
		serve.HandleFunc(fmt.Sprintf("%s %s", endpoint.Method, endpoint.Path), func(w http.ResponseWriter, r *http.Request) {
			request := (*Request)(r)

			switch {
			case endpoint.Handler.Api != nil:
				endpoint.Handler.Api(request, w)
			case endpoint.Handler.Web != nil:
				middleware.WebParser(w, endpoint.Title, endpoint.Handler.Web(request))
			}
		})
	}
	println("System Powering Up")
	println("Powered by Go")
	println("Developed under Master Piece of Paper")
	println("Welcome to Constellation Server")

	middleware.SetStylesheet(string(css))

	http.ListenAndServe(":8000", serve)
}
