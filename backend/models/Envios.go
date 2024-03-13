package models

type Envio struct {
	EnvioID           int
	PedidoClienteID   int
	FechaEnvio        string
	Transportista     string
	Estado            string
	CodigoSeguimiento string
}
