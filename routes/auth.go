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
		auth.POST("/refresh-token", controllers.RefreshToken)

		emailVerification := auth.Group("/email-verification")
		{
			emailVerification.POST("", controllers.VerifyUser)
			emailVerification.POST("/resend", controllers.ResendEmailVerification)
		}

		resetPassword := auth.Group("/reset-password")
		{
			resetPassword.POST("/send", controllers.SendForgotPasswordRequest)
			resetPassword.POST("", controllers.ResetPassword)
		}
	}
}
