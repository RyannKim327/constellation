package constellation

import "net/http"

type (
	HandlerWeb     func(r *Request) string
	HandlerApi     func(r *Request, w Response)
	MethodType     string
	PermissionType string
	Request        http.Request
	Response       http.ResponseWriter
	APIResponse    map[string]any
)

type HandlerSetup struct {
	Web HandlerWeb
	Api HandlerApi
}

const (
	GET    MethodType = "GET"
	POST   MethodType = "POST"
	PUT    MethodType = "PUT"
	PATCH  MethodType = "PATCH"
	DELETE MethodType = "DELETE"
)

const (
	Unauthenticated PermissionType = "__public__"
	Authenticated   PermissionType = "__private__"
	Authorized      PermissionType = "__protected__"
)

type WebComponent struct {
	Class      string
	Attributes map[string]string
	Children   []any
	Text       string
}

type Endpoint struct {
	Path       string
	Method     MethodType
	Handler    HandlerSetup
	Permission PermissionType
	Title      string
}

var Endpoints []Endpoint

func AddEndpoint(endpoint Endpoint) {
	if endpoint.Path == "" {
		endpoint.Path = "/"
	}

	if endpoint.Method == "" {
		endpoint.Method = GET
	}

	Endpoints = append(Endpoints, endpoint)
}
