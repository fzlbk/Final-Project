package main

import (
	"log"

	"product-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Локальные middleware
	r.Use(func(c *gin.Context) {
		// Пример базового middleware с логом
		log.Println("Request received:", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// Регистрируем только свои маршруты
	routes.ProductRoutes(r)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to run product-service: %v", err)
	}
}
