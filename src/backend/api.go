package backend

import (
	"constellation/config/middleware"
	"constellation/config/utils"

	"github.com/gin-gonic/gin"
)

var API = middleware.Route{
	Path:    "/api",
	Method:  utils.METHOD_GET,
	Handler: api,
}

func api(ctx *gin.Context) (string, utils.PARAMS) {
	ctx.JSON(200, gin.H{
		"message": "Hello",
	})
	return "", nil
}
