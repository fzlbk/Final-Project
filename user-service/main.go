package main

import (
	"user-service/controllers"
	"user-service/database"
	"user-service/middleware"
	"user-service/routes"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.New()
	r.Use(cors.Default())

	r.Use(middleware.LoggingMiddleware())

	// Аутентификация
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Пользователи
	routes.UserRoutes(r)

	r.Run(":8081")
}
