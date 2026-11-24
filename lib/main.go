package constellation

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"constellation/lib/middleware"
)

func NewHandler(css []byte) http.Handler {
	serve := http.NewServeMux()

	for _, endpoint := range Endpoints {
		method := endpoint.Method
		if method == "" {
			method = GET
		}
		path := endpoint.Path
		if path == "" {
			path = "/"
		}
		pattern := fmt.Sprintf("%s %s", method, path)
		if path == "/" {
			pattern = fmt.Sprintf("%s /{$}", method)
		}

		serve.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			request := (*Request)(r)

			if !middleware.AccessController(endpoint.Permission, r) {
				handler := errorPages.Forbidden
				if handler == nil {
					handler = forbiddenPage
				}
				middleware.WebPage(w, http.StatusForbidden, "403 · Forbidden", handler(request))
				return
			}

			switch {
			case endpoint.Handler.Api != nil:
				endpoint.Handler.Api(request, w)
			case endpoint.Handler.Web != nil:
				middleware.WebParser(w, endpoint.Title, endpoint.Handler.Web(request))
			}
		})
	}

	serve.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := errorPages.NotFound
		if handler == nil {
			handler = notFoundPage
		}
		middleware.WebPage(w, http.StatusNotFound, "404 · Page Not Found", handler((*Request)(r)))
	})

	middleware.SetStylesheet(string(css))

	return serve
}

func StartServer(css []byte) {
	serve := NewHandler(css)

	println("System Powering Up")
	println("Powered by Go")
	println("Developed under Master Piece of Paper")
	println("Welcome to Constellation Server")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(http.ListenAndServe(":"+port, serve))
}

func notFoundPage(_ *Request) string {
	return defaultErrorPage("404", "Page Not Found", "The page you requested does not exist.")
}

func forbiddenPage(_ *Request) string {
	return defaultErrorPage("403", "Forbidden", "You do not have permission to access this page.")
}

func defaultErrorPage(code, title, message string) string {
	return fmt.Sprintf(
		`<section class="page-shell mx-auto max-w-[1120px] px-6 py-24"><span class="eyebrow block">%s</span><h1 class="page-title">%s</h1><p class="page-lead">%s</p></section>`,
		html.EscapeString(code),
		html.EscapeString(title),
		html.EscapeString(message),
	)
}
