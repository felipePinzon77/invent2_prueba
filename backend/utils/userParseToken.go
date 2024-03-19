package utils

import (
	"log"
	"os"

	"github.com/felipePinzon77/invent2_prueba/backend/models"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenString string) (claims *models.Claims, err error) {

	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		log.Fatal("Not set the key")
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}
