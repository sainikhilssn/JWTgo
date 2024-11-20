package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sainikhilssn/JWTgo/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true})
	})

	router.Run(":" + port)
}
