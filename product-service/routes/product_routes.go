// файл: product-service/routes/product_routes.go

package routes

import (
	"product-service/controllers"
	"product-service/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	// ==== Продукты ====
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProductByID)
	router.POST("/products", controllers.CreateProduct)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)

	router.GET("/products/count", controllers.GetProductCount)
	router.GET("/products/price-range", controllers.GetProductsByPriceRange)
	router.GET("/products/search", controllers.SearchProducts)
	router.PATCH("/products/:id/price", controllers.PatchProductPrice)
	router.GET("/products/newest", controllers.GetNewestProducts)
	router.GET("/products/expensive", controllers.GetExpensiveProducts)
	router.GET("/brands", controllers.GetBrands)
	router.GET("/brands/:brand/products", controllers.GetProductsByBrand)
	router.GET("/stats/products", controllers.ProductStats)

	// ==== Прокси на user-service ====
	router.GET("/profile", middleware.AuthMiddleware(), controllers.GetProfileProxy)
}
