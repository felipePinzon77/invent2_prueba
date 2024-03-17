package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/felipePinzon77/invent2_prueba/backend/controllers"

)

func AuthRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)
    r.POST("/signup", controllers.Singup)
    r.GET("/", controllers.Home)
    r.GET("/Admin", controllers.Premium)
    r.GET("/logout", controllers.Logout)
}