package controllers_test

import (
	"bytes"
	"encoding/json"
	"music-store/controllers"
	"music-store/database"
	"music-store/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRouter() *gin.Engine {
	database.Connect()
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductByID)
	return r
}

func TestRegister(t *testing.T) {
	r := SetupTestRouter()
	w := httptest.NewRecorder()

	body := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLogin(t *testing.T) {
	r := SetupTestRouter()
	w := httptest.NewRecorder()

	body := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateProduct(t *testing.T) {
	r := SetupTestRouter()
	w := httptest.NewRecorder()

	product := models.Product{
		Name:        "Guitar",
		Description: "Electric guitar",
		Price:       299.99,
		Brand:       "Yamaha",
	}
	jsonValue, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProducts(t *testing.T) {
	r := SetupTestRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/products", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProductByID(t *testing.T) {
	r := SetupTestRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/products/1", nil)
	r.ServeHTTP(w, req)
	assert.True(t, w.Code == http.StatusOK || w.Code == http.StatusNotFound)
}

func TestRegisterInvalidJSON(t *testing.T) {
	r := SetupTestRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/register", bytes.NewBufferString("invalid_json"))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginInvalidJSON(t *testing.T) {
	r := SetupTestRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString("invalid_json"))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateProductInvalidJSON(t *testing.T) {
	r := SetupTestRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/products", bytes.NewBufferString("invalid_json"))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetNonExistentProduct(t *testing.T) {
	r := SetupTestRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/products/999999", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
