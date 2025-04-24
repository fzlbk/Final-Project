package main

import (
	"music-store/database"
	"music-store/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	routes.ProductRoutes(r)
	routes.UserRoutes(r)

	r.Run(":8080")
}
