package controllers

import (
	"music-store/database"
	"music-store/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	brand := c.Query("brand")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	db := database.DB.Limit(limit).Offset(offset)

	if brand != "" {
		db = db.Where("brand ILIKE ?", "%"+brand+"%")
	}
	if minPrice != "" {
		db = db.Where("price >= ?", minPrice)
	}
	if maxPrice != "" {
		db = db.Where("price <= ?", maxPrice)
	}
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Brand = input.Brand
	database.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	database.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func PatchProductPrice(c *gin.Context) {
	id := c.Param("id")
	var data struct {
		Price float64 `json:"price"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	database.DB.Model(&models.Product{}).Where("id = ?", id).Update("price", data.Price)
	c.JSON(http.StatusOK, gin.H{"message": "Price updated"})
}

func GetNewestProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Order("created_at desc").Limit(5).Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetExpensiveProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Order("price desc").Limit(5).Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetBrands(c *gin.Context) {
	var brands []string
	database.DB.Model(&models.Product{}).Distinct().Pluck("brand", &brands)
	c.JSON(http.StatusOK, brands)
}

func ProductStats(c *gin.Context) {
	type Result struct {
		Count int     `json:"count"`
		Avg   float64 `json:"average_price"`
	}
	var result Result
	database.DB.Model(&models.Product{}).
		Select("count(*) as count, avg(price) as avg").
		Scan(&result)
	c.JSON(http.StatusOK, result)
}
func GetProductsByBrand(c *gin.Context) {
	brand := c.Param("brand")

	var products []models.Product
	if err := database.DB.Where("brand = ?", brand).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products by brand"})
		return
	}

	c.JSON(http.StatusOK, products)
}
func GetProductCount(c *gin.Context) {
	var count int64
	if err := database.DB.Model(&models.Product{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}
func GetProductsByPriceRange(c *gin.Context) {
	minPrice := c.Query("min")
	maxPrice := c.Query("max")

	var products []models.Product
	db := database.DB

	if minPrice != "" {
		db = db.Where("price >= ?", minPrice)
	}
	if maxPrice != "" {
		db = db.Where("price <= ?", maxPrice)
	}

	if err := db.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
func SearchProducts(c *gin.Context) {
	query := c.Query("query")
	var products []models.Product

	if err := database.DB.Where("name ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%").
		Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
