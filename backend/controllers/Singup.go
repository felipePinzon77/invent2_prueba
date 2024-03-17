package controllers

import (
	"github.com/felipePinzon77/invent2_prueba/backend/db"
	"github.com/felipePinzon77/invent2_prueba/backend/models"
	"github.com/felipePinzon77/invent2_prueba/backend/utils"

	"github.com/gin-gonic/gin"
)

// The string "my_secret_key" is just an example and should be replaced with a secret key of sufficient length and complexity in a real-world scenario.


func Singup(c *gin.Context) {

    conexionEsta, _ := db.ConexionDB()

	var user models.Usuario

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    var existingUser models.Usuario

	query := "SELECT * FROM users WHERE NombreUsuario = ? LIMIT 1"
	stmt, _ := conexionEsta.Prepare(query)

    stmt.QueryRow(user.NombreUsuario).Scan(&existingUser)

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

	querys := "INSERT INTO users (NombreUsuario) VALUES (?)"
	stmts, _ := conexionEsta.Prepare(querys)

    stmts.Exec(&user)

    c.JSON(200, gin.H{"success": "user created"})
}