package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var stylesheet string

// TODO: SetStylesheet provides the CSS processed by the Tailwind browser runtime.
func SetStylesheet(css string) {
	stylesheet = css
}

func SendApi(w http.ResponseWriter, data map[string]any) {
	json.NewEncoder(w).Encode(data)
}

func WebParser(w http.ResponseWriter, title string, html string) (int, error) {
	return WebPage(w, http.StatusOK, title, html)
}

func WebPage(w http.ResponseWriter, status int, title string, html string) (int, error) {
	if status == 0 {
		status = http.StatusOK
	}
	if title == "" {
		title = "My Constellation App"
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Powered-By", "Constellation")
	w.Header().Set("Developed-Under", "Master Piece of Paper")
	w.Header().Set("Developed-By", "Ryann Kim Sesgundo")
	w.WriteHeader(status)

	_html := fmt.Sprintf(`<!DOCTYPE html>
		<html lang="en">
		  <!-- Constellation -->
		  <!-- A framework developed with Go language -->
  		<!-- Developed by Ryann Kim Sesgundo -->
  		<head>
    		<meta charset="UTF-8">
    		<meta name="viewport" content="width=device-width, initial-scale=1">
    		<meta name="color-scheme" content="dark">
    	  <title>%s</title>
    		<script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>
        <style type="text/tailwindcss">%s</style>
  		</head>
  		<body>
    		<main>%s</main>
  		</body>
		</html>
		`, title, stylesheet, html)
	return w.Write([]byte(_html))
}
