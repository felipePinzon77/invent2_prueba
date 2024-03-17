package models

import "github.com/dgrijalva/jwt-go"

type Usuario struct {
	UsuarioID     int
	EmpleadoID    int
	NombreUsuario string
	Role 		  string
	Password      string
}

type Claims struct {
    Role string `json:"role"`
    jwt.StandardClaims
}