package main

import (
	"fmt"
	"net/http"
	"github.com/felipePinzon77/invent2_prueba/backend/api"

)

func main() {
    // Manejador para la ruta "/"
	http.HandleFunc("/", welcome)

	http.HandleFunc("/crear_Producto", api.Post)
	http.HandleFunc("/actualizar_Producto", api.Update)
	http.HandleFunc("/borrar_Producto", api.Delete)
	http.HandleFunc("/select_Producto", api.Front)

    // Definir el puerto en el que el servidor escuchará las solicitudes
    port := ":8080"

    // Iniciar el servidor
    fmt.Printf("Servidor escuchando en el puerto %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        fmt.Printf("Error al iniciar el servidor: %s\n", err)
    }
}

func welcome(http.ResponseWriter, *http.Request){
	fmt.Println("Hola")
}
