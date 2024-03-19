package controllers

import (
	"github.com/felipePinzon77/invent2_prueba/backend/db"
	"github.com/felipePinzon77/invent2_prueba/backend/models"
	"github.com/felipePinzon77/invent2_prueba/backend/utils"

	"github.com/gin-gonic/gin"
)

func Singup(c *gin.Context) {

	var user models.Usuario

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.Usuario

	db.DB.Where("NombreUsuario = ?", user.NombreUsuario).First(&existingUser)

	if existingUser.UsuarioID != 0 {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	db.DB.Create(&user)

	c.JSON(200, gin.H{"success": "user created"})
}
