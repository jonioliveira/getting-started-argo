package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPingRoutes(rg *gin.RouterGroup) {
	pingRouter := rg.Group("/ping")

	pingRouter.GET("/", pongHandler)
}

func pongHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
