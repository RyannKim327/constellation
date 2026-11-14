package backend

import (
	"github.com/RyannKim327/constellation/config/middleware"
	"github.com/RyannKim327/constellation/config/utils"

	"github.com/gin-gonic/gin"
)

// SampleJSON demonstrates a handler that writes its own JSON response.
var SampleJSON = middleware.Route{
	Method:     utils.METHOD_GET,
	Path:       "/samples/json",
	Handler:    sampleJSON,
	Permission: utils.PERMISSION_ALL,
	Title:      "JSON sample · Constellation",
}

func sampleJSON(ctx *gin.Context) (string, utils.PARAMS) {
	ctx.JSON(200, gin.H{
		"message": "Hello from a Constellation sample",
		"method":  ctx.Request.Method,
		"path":    ctx.Request.URL.Path,
	})
	return "", nil
}
