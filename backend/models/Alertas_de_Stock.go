package models

type AlertaStock struct {
	AlertaID   int
	ProductoID int
	Umbral     int
	Mensaje    string
}
