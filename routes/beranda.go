package routes

import (
	"reservation/controllers"
	"reservation/middlewares"

	"github.com/gin-gonic/gin"
)

func BerandaRoute(route *gin.RouterGroup) {
	beranda := route.Group("/beranda", middlewares.Auth())
	{
		carousel := beranda.Group("/carousel")
		{
			carousel.POST("/upload", controllers.UploadCarousel)
			carousel.GET("/update-carousel-status/:id", controllers.UpdateCarouselStatus)
		}
	}
}
