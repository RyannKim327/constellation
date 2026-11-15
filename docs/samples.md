# Samples

The sample routes live in `src/routes/` and exercise both supported handler
styles.

## JSON endpoint

`sample_json.go` writes its own response and returns an empty body:

```go
func sampleJSON(ctx *gin.Context) (string, utils.PARAMS) {
    ctx.JSON(200, gin.H{"message": "Hello"})
    return "", nil
}
```

Visit `/samples/json` to see the result.

## Form page

`sample_form.go` pairs a `GET` page with a `POST` handler. The `POST` route sets
`Permission: utils.PERMISSION_ALL` so it does not require the `temporary`
cookie.

The handler reads the `name` field and returns it as a parameter:

```go
return `<h1>Hello, {name}</h1>`, utils.PARAMS{"name": ctx.PostForm("name")}
```

Because `name` is rendered as a parameter, it is HTML-escaped. Visit
`/samples/form` to try it.

## The backend API

`src/backend/api.go` is the smallest possible JSON route:

```go
var API = middleware.Route{
    Method:  utils.METHOD_GET,
    Path:    "/api",
    Handler: api,
}
```

All samples are registered in `src/main.go` together with the frontend routes.
