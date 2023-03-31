package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sql-api/controllers"
)

func HealthCheck(rg *gin.RouterGroup) {
	healthRouterGroup := rg.Group("/")
	healthRouterGroup.GET("/healthcheck", controllers.HealthCheck())
}
