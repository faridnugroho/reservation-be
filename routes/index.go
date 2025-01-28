package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(route *gin.Engine) {
	route.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"responseCode":    "404",
			"responseMessage": "Invalid Routing",
		})
	})

	serviceCheck := route.Group("/check")
	api := route.Group("/api")
	v1 := api.Group("/v1")

	ServiceCheck(serviceCheck)

	TestRoute(v1)
	AuthRoute(v1)
}

func ServiceCheck(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"responseCode":    "200",
			"responseMessage": "Service is running",
		})
	})
}
