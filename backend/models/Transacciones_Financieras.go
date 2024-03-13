package models

type TransaccionFinanciera struct {
	TransaccionID   int
	Tipo            string
	Monto           float64
	Fecha           string
	OrdenCompraID   int
	PedidoClienteID int
}
