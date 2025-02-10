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
			carousel.POST("", controllers.UploadCarousel)
			carousel.GET("", controllers.GetCarousels)
			carousel.PATCH("/:id", controllers.UpdateCarousel)
			carousel.DELETE("/:id", controllers.DeleteCarousel)
			carousel.GET("/update-carousel-status/:id", controllers.UpdateCarouselStatus)
		}
	}
}
