package models

type Factura struct {
	FacturaID       int
	PedidoClienteID int
	FechaFactura    string
	Total           float64
}
