package routes

import "github.com/gin-gonic/gin"

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users")
	incomingRoutes.GET("/users/:user_id")
}
