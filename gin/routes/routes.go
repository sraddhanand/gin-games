package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sraddhanand/profitGames/app/v1"
)

func Routes(router *gin.Engine) {
	hcRoutes := router.Group("/ops")
	hcRoutes.GET("/hc", v1.HealthCheck)
	hcRoutes.GET("/echo/:msg", v1.GetEchoREST)
}
