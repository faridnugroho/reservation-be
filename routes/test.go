package routes

import (
	"reservation/middlewares"

	"github.com/gin-gonic/gin"
)

func TestRoute(route *gin.RouterGroup) {
	test := route.Group("/test", middlewares.Auth())
	{
		test.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		test.GET("/testing", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "Berhasil get data",
			})
		})
	}
}
