package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sql-api/initializers"
	"github.com/sql-api/routes"
)

// func getFunc(c *gin.Context) {

// 	c.Status(http.StatusOK)
// }

func init() {
	initializers.LoadEnv()
	initializers.Init()
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := router.Group("/")
	v1 := router.Group("/v1")

	routes.HealthCheck(health)
	routes.AddUserRoutes(v1)
	router.Run(":" + port)

}
