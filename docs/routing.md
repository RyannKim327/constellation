# Routing

Routes are declared as `constellation.Endpoint` values and registered with
`AddEndpoint` before the server starts.

## Endpoint fields

```go
type Endpoint struct {
    Path       string
    Method     MethodType
    Handler    HandlerSetup
    Permission PermissionType
    Title      string
}

type HandlerSetup struct {
    Web HandlerWeb // func(r *Request) string
    Api HandlerApi // func(r *Request, w Response)
}
```

| Field        | Required | Description                                      |
| ------------ | -------- | ------------------------------------------------ |
| `Path`       | no       | URL path. Defaults to `/`.                       |
| `Method`     | no       | HTTP method. Defaults to `GET`.                  |
| `Handler`    | yes      | Web and/or Api callbacks for the endpoint.       |
| `Permission` | no       | Access label. Defaults to empty when omitted.    |
| `Title`      | no       | Page title. Falls back to `Constellation`.       |

## Example

```go
constellation.AddEndpoint(constellation.Endpoint{
    Path:       "/hello",
    Method:     constellation.GET,
    Permission: constellation.Unauthenticated,
    Title:      "Hello",
    Handler: constellation.HandlerSetup{
        Web: func(r *constellation.Request) string {
            return "<h1>Hello</h1>"
        },
    },
})
```

The server registers each endpoint as an exact `METHOD path` pattern, for
example `GET /hello`. The root path is registered as `GET /{$}` so other paths
reach the 404 handler. `POST`, `PUT`, `PATCH`, and `DELETE` are also available.
See the [API guide](api.md) for a full walkthrough of JSON endpoints.

## Web and Api handlers

`HandlerSetup` combines both handler styles. Set `Web` to return an HTML body,
which is wrapped by `middleware.WebParser`, or set `Api` to write the response
yourself. If both are set, `Api` takes precedence.

```go
constellation.AddEndpoint(constellation.Endpoint{
    Path:   "/api/hello",
    Method: constellation.GET,
    Handler: constellation.HandlerSetup{
        Api: func(r *constellation.Request, w constellation.Response) {
            w.Header().Set("Content-Type", "application/json")
            w.Write([]byte(`{"message":"Hello"}`))
        },
    },
})
```

## Defaults, permissions, and error pages

`AddEndpoint` fills in `/` for an empty path, `GET` for an empty method, and
`Unauthenticated` for an empty permission. Permission values are checked before
handlers run: public endpoints are always allowed, while `Authenticated` and
`Authorized` endpoints depend on `middleware.AuthenticatedCheck` and
`middleware.AuthorizedCheck` and return `403 Forbidden` when denied.

Unmatched paths return `404 Not Found`. Both status codes render an HTML body
wrapped by the page parser. Register the built-in pages, or your own
`HandlerWeb` functions, before starting the server:

```go
import (
    constellation "constellation/lib"
    "constellation/lib/status"
)

constellation.SetErrorPages(constellation.ErrorPages{
    NotFound:  status.Status404,
    Forbidden: status.Status403,
})
```

`NotFound` receives unmatched requests and `Forbidden` receives permission
denials. Each handler gets the current request and returns the HTML body.

## Reading request data

A handler receives a `*constellation.Request`, which is an alias for
`*http.Request`. Query values and form data use the usual standard library API:

```go
name := r.URL.Query().Get("name")
if name == "" {
    name = "world"
}
```
