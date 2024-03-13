package models

type Devolucion struct {
	DevolucionID int
	Tipo         string
	ProductoID   int
	Cantidad     int
	Fecha        string
	Motivo       string
}
