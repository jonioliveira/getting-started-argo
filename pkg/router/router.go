package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jonioliveira/getting-started-argo/pkg/ping"
	"github.com/jonioliveira/getting-started-argo/pkg/user"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

const (
	v1 = "/v1"
	v2 = "/v2"
)

func BuildRouter() *gin.Engine {
	router := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	apiV1 := router.Group(v1)
	apiV2 := router.Group(v2)
	ping.AddPingRoutes(apiV1)
	user.AddUserRoutesV1(apiV1)
	user.AddUserRoutesV2(apiV2)

	return router
}
