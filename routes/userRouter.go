package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sainikhilssn/JWTgo/controllers"
	"github.com/sainikhilssn/JWTgo/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// should authenticate before using the routes
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/user/:user_id", controllers.GetUser())

}
