package models

type PedidoCliente struct {
	PedidoClienteID int
	ClienteID       int
	FechaPedido     string
	Estado          string
	EmpleadoID      int
}
