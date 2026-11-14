# Routing

Routes are declared as `middleware.Route` values and collected in
`src/index.go`.

## The `Route` type

```go
type Route struct {
    Method     string
    Path       string
    Permission string
    Handler    utils.HandlerSetup
}
```

| Field        | Required | Description                                         |
| ------------ | -------- | --------------------------------------------------- |
| `Method`     | no       | HTTP method. Defaults to `GET`. Case-insensitive.   |
| `Path`       | yes      | URL path. Lowercased automatically.                 |
| `Permission` | no       | Access level. Defaults by method (see below).       |
| `Handler`    | yes      | The `func(ctx *gin.Context) (string, map[string]any)`. |

## Declaring a route

```go
package routes

import (
    "constellation/config/middleware"
    "constellation/config/utils"

    "github.com/gin-gonic/gin"
)

var About = middleware.Route{
    Method: utils.METHOD_GET,
    Path:   "/about",
    Handler: func(ctx *gin.Context) (string, map[string]any) {
        return `<h1>About {app}</h1>`, map[string]any{"app": "Constellation"}
    },
}
```

Register it in `src/index.go`:

```go
var _routes = []mw.Route{
    backend.API,
    routes.Index,
    routes.About,
}
```

## Registration behavior

`mw.Register` normalizes each route before attaching it to the engine:

- **Method** — defaults to `GET` when empty, then uppercased.
- **Path** — lowercased.
- **POST suffix** — a `POST` route whose permission is *not* `ALL` and whose path
  does not already end in `/submit` gets `/submit` appended. This gives
  form-submission endpoints a consistent shape.
- **Permission default** — `POST` defaults to `COOKIE`; every other method
  defaults to `ALL`.

## Available constants

Defined in `config/utils/statics.go`:

```go
// Methods
METHOD_GET, METHOD_POST, METHOD_PUT, METHOD_PATCH, METHOD_DELETE

// Permissions
PERMISSION_ALL     // no auth
PERMISSION_COOKIE  // requires a valid `temporary` cookie
PERMISSION_ADMIN   // reserved; not enforced yet
```

## Example: JSON endpoint

A handler can write the response itself and return `(nil, nil)` to skip HTML
rendering:

```go
var API = middleware.Route{
    Method:  utils.METHOD_GET,
    Path:    "/api",
    Handler: func(ctx *gin.Context) (string, map[string]any) {
        ctx.JSON(200, gin.H{"message": "Hello"})
        return "", nil
    },
}
```
