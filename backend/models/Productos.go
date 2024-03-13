package models

type Producto struct {
	ProductoID              int
	Nombre                  string
	Descripcion             string
	Precio                  float64
	UnidadMedida            string
	ProveedorID             int
	LicenciaCertificacionID int
}
