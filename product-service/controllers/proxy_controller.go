package controllers

import (
	"net/http"
	"product-service/utils"

	"github.com/gin-gonic/gin"
)

func GetProfileProxy(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	resp, err := utils.UserClient.R().
		SetHeader("Authorization", auth).
		Get("/me")
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}
