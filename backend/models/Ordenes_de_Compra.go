package models

type OrdenCompra struct {
	OrdenCompraID int
	ProveedorID   int
	ProductoID    int
	Cantidad      int
	FechaOrden    string
	Estado        string
	EmpleadoID    int
}
