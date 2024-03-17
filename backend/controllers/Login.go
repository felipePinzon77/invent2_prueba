package controllers

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/felipePinzon77/invent2_prueba/backend/db"
	"github.com/felipePinzon77/invent2_prueba/backend/models"
	"github.com/felipePinzon77/invent2_prueba/backend/utils"
	"github.com/gin-gonic/gin"
)

// The string "my_secret_key" is just an example and should be replaced with a secret key of sufficient length and complexity in a real-world scenario.

func Login(c *gin.Context) {

	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		log.Fatal("Not set the key")
	}

	conexionEsta, _ := db.ConexionDB()

	var user models.Usuario

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.Usuario

	query := "SELECT * FROM users WHERE email = ? LIMIT 1"
	stmt, _ := conexionEsta.Prepare(query)

	stmt.QueryRow(user.NombreUsuario).Scan(&existingUser)

	if existingUser.EmpleadoID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.NombreUsuario,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{"error": "couldn't generate token"})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged in"})
}
