package models

type PagoProveedor struct {
	PagoProveedorID int
	OrdenCompraID   int
	FechaPago       string
	Monto           float64
}
