package routes

import (
	"user-service/controllers"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	router.GET("/profile", controllers.GetProfile)
	router.GET("/me", middleware.AuthMiddleware(), controllers.GetMe)
}
