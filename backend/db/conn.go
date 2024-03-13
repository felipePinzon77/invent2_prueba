package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ConexionDB inicializa la conexión a la base de datos local
func ConexionDB() (*sql.DB, error) {
	Driver := "mysql"
	Usuario := "root"
	Contraseña := ""
	Nombre := "invent2"

	// instruccion
	conexion, err := sql.Open(Driver, Usuario+":"+Contraseña+"@tcp(127.0.0.1)/"+Nombre)

	if err != nil {
		return nil, err
	}

	log.Println("Database is running")
	return conexion, nil
}