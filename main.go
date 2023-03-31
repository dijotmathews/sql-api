package main

import (
	"os"

	"github.com/gin-contrib/cors"
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

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"POST"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "http://localhost:3000"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

	router.Use(cors.Default())

	health := router.Group("/")
	v1 := router.Group("/v1")

	routes.HealthCheck(health)
	routes.AddUserRoutes(v1)
	router.Run(":" + port)

}
