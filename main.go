package main

import (
	"os"
	"reservation/config"
	"reservation/middlewares"
	"reservation/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := InitServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Run(":" + port)
}

func InitServer() *gin.Engine {
	route := gin.Default()

	route.Use(middlewares.CORSMiddleware())
	route.Use(middlewares.Logger())
	route.Use(gin.Recovery())

	config.Database()

	routes.SetupRoutes(route)

	return route
}
