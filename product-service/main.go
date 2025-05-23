package main

import (
	"log"

	"product-service/database"
	"product-service/models"
	"product-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Product{})

	r := gin.Default()
	r.Use(cors.Default())

	r.Use(func(c *gin.Context) {
		log.Println("Request received:", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	routes.ProductRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run product-service: %v", err)
	}
}
