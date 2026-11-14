package utils

import (
	"fmt"
	"html"
	"strings"

	"github.com/gin-gonic/gin"
)

// HTParser replaces every "{key}" token in body with the string form of the
// matching value in params. Values are substituted verbatim and are NOT
// HTML-escaped, so only pass trusted markup. Use Render for user-supplied data.
func HTParser(body string, params map[string]any) string {
	if len(params) == 0 {
		return body
	}

	pairs := make([]string, 0, len(params)*2)
	for key, value := range params {
		pairs = append(pairs, "{"+key+"}", fmt.Sprint(value))
	}

	return strings.NewReplacer(pairs...).Replace(body)
}

func template(body string, params map[string]any) string {
	if len(params) == 0 {
		return body
	}

	pairs := make([]string, 0, len(params)*2)
	for key, value := range params {
		pairs = append(pairs, "{"+key+"}", html.EscapeString(fmt.Sprint(value)))
	}

	return strings.NewReplacer(pairs...).Replace(body)
}

// Render wraps body in the full HTML document shell, substitutes params with
// HTML escaping, and writes the result as text/html. It is called automatically
// for handlers that return a non-empty body or a non-nil parameter map.
func Render(ctx *gin.Context, title string, body string, params map[string]any) {
	if title == "" {
		title = "Constellation"
	}

	_html := fmt.Sprintf(`<!DOCTYPE html>
		<html lang="en">
  		<head>
    		<meta charset="UTF-8">
    		<meta name="viewport" content="width=device-width, initial-scale=1">
    		<title>%s</title>
    		<script src="https://cdn.tailwindcss.com"></script>
        <link rel="stylesheet" href="/static/styles.css" />
  		</head>
  		<body>
    		<main class="w-dvw h-dvh">%s</main>
  		</body>
		</html>
		`, title, template(body, params))
	ctx.Data(200, "text/html; charset=utf-8", []byte(_html))
}
