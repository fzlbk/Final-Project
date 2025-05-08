package utils

import (
	"time"
	"user-service/models"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret_key")

func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
