package main

import (
	"os"

	"github.com/gin-gonic/gin"
	routes "github.com/techbrolakes/golang-jwt/routes"
)

func main() {
	port := os.Getenv("")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)

	router.Run(":" + port)
}
