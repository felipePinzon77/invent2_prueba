package main

import (
	"fmt"
	"net/http"
	"github.com/felipePinzon77/invent2_prueba/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    // Manejador para la ruta "/"
	// http.HandleFunc("/", welcome)

	// http.HandleFunc("/crear_Producto", controllers.Post)
	// http.HandleFunc("/actualizar_Producto", controllers.Update)
	// http.HandleFunc("/borrar_Producto", controllers.Delete)
	// http.HandleFunc("/select_Producto", controllers.Front)
    
    r := gin.Default()
    
    routes.AuthRoutes(r)

    // Definir el puerto en el que el servidor escuchará las solicitudes
    port := ":8080"
    
    // Iniciar el servidor
    fmt.Printf("Servidor escuchando en el puerto %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        fmt.Printf("Error al iniciar el servidor: %s\n", err)
    }

}
