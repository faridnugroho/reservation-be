package routes

import (
	"reservation/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(route *gin.RouterGroup) {
	auth := route.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/verify-user", controllers.VerifyUser)
		auth.POST("/refresh-token", controllers.RefreshToken)
	}
}
