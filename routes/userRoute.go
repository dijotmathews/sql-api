package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sql-api/controllers"
	"github.com/sql-api/middlewares"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	onboardingRoutes := rg.Group("/api")
	onboardingRoutes.POST("/run", middlewares.Validate(""), controllers.RunQuery())
}
