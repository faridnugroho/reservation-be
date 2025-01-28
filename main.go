package main

import (
	"reservation/config"
	"reservation/middlewares"
	"reservation/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := InitServer()

	app.Run(":8080")
}

func InitServer() *gin.Engine {
	route := gin.Default()

	route.Use(middlewares.CORSMiddleware())
	route.Use(middlewares.CustomLogger())
	route.Use(gin.Recovery())

	config.Database()

	routes.SetupRoutes(route)

	return route
}
