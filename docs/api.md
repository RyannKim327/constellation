# API

An API endpoint writes its own response instead of returning an HTML string. Set the `Api` handler on an endpoint to serve JSON from your Constellation app.

## Define an endpoint

An API handler has the signature `func(r *constellation.Request, w constellation.Response)`. `Request` exposes the standard `net/http` request fields, and `Response` is an `http.ResponseWriter`, so headers, status codes, and body writes all work as usual.

```go
constellation.AddEndpoint(constellation.Endpoint{
    Path:       "/api/hello",
    Method:     constellation.GET,
    Permission: constellation.Unauthenticated,
    Handler: constellation.HandlerSetup{
        Api: func(r *constellation.Request, w constellation.Response) {
            w.Header().Set("Content-Type", "application/json")
            w.Write([]byte(`{"message":"Hello"}`))
        },
    },
})
```

`AddEndpoint` fills in `/` for an empty path and `GET` for an empty method. Call it before `constellation.StartServer`. Routes are registered as exact `METHOD path` patterns, and `GET`, `POST`, `PUT`, `PATCH`, and `DELETE` are available.

`Title` is used by the HTML page wrapper, so it can be left out for API endpoints. If `Web` and `Api` are both set on the same endpoint, `Api` takes precedence and the HTML wrapper is skipped.

## Write JSON

`middleware.SendApi` encodes a value as JSON:

```go
import "constellation/lib/middleware"

middleware.SendApi(w, map[string]any{"message": "Hello"})
```

It uses `json.Encoder.Encode`, which appends a newline and does not set the content type. Set the header before calling it:

```go
w.Header().Set("Content-Type", "application/json")
middleware.SendApi(w, map[string]any{"message": "Hello"})
```

Write the body yourself when you need control over encoding errors or status codes:

```go
payload, err := json.Marshal(map[string]any{"message": "Created"})
if err != nil {
    http.Error(w, "encoding failed", http.StatusInternalServerError)
    return
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
w.Write(payload)
```

Call `WriteHeader` before writing the body; once bytes are written, the status can no longer be changed.

## Read request data

Query string and form values use the standard library API:

```go
name := r.URL.Query().Get("name")

if err := r.ParseForm(); err != nil {
    http.Error(w, "bad form", http.StatusBadRequest)
    return
}
name = r.FormValue("name")
```

Decode a JSON body from `r.Body`:

```go
var input struct {
    Name string `json:"name"`
}
if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
    http.Error(w, "invalid JSON", http.StatusBadRequest)
    return
}
```

Routes may contain standard library wildcards, which are read with `PathValue`:

```go
constellation.Endpoint{
    Path: "/api/items/{id}",
    Method: constellation.GET,
    Handler: constellation.HandlerSetup{
        Api: func(r *constellation.Request, w constellation.Response) {
            id := r.PathValue("id")
            w.Header().Set("Content-Type", "application/json")
            middleware.SendApi(w, map[string]any{"id": id})
        },
    },
}
```

## Complete example

```go
package main

import (
    "encoding/json"
    "net/http"

    constellation "constellation/lib"
    "constellation/lib/middleware"
)

func main() {
    constellation.AddEndpoint(constellation.Endpoint{
        Path:   "/api/hello",
        Method: constellation.GET,
        Handler: constellation.HandlerSetup{
            Api: func(r *constellation.Request, w constellation.Response) {
                w.Header().Set("Content-Type", "application/json")
                middleware.SendApi(w, map[string]any{"message": "Hello"})
            },
        },
    })

    constellation.AddEndpoint(constellation.Endpoint{
        Path:       "/api/echo",
        Method:     constellation.POST,
        Permission: constellation.Unauthenticated,
        Handler: constellation.HandlerSetup{
            Api: echo,
        },
    })

    constellation.StartServer(nil)
}

func echo(r *constellation.Request, w constellation.Response) {
    var input struct {
        Name string `json:"name"`
    }
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }
    if input.Name == "" {
        input.Name = "world"
    }

    w.Header().Set("Content-Type", "application/json")
    middleware.SendApi(w, map[string]any{"message": "Hello, " + input.Name})
}
```

Run the app with `go run .` and try the endpoints:

```bash
curl -i http://localhost:8000/api/hello
curl -i -X POST -H 'Content-Type: application/json' \
  -d '{"name":"Ada"}' http://localhost:8000/api/echo
```

## Current limits

Permission values are enforced before a handler runs. `Unauthenticated` is public, while `Authenticated` and `Authorized` depend on the checks configured in `lib/middleware/access-controller.go` and return `403 Forbidden` by default. There are no built-in helpers for authentication, CORS, or API versioning; use the standard `net/http` APIs inside a handler for those. `SendApi` only encodes JSON, so headers and status codes remain your responsibility.

## See also

- [Routing](routing.md) for endpoint fields and method constants.
- [Rendering](rendering.md) for the HTML handler alternative.
- [Getting started](getting-started.md) for the minimal app setup.
