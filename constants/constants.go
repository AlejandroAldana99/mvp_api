package constants

const (
	MeasureS            = "S"
	MeasureM            = "M"
	MeasureL            = "L"
	MeasureSpecial      = "Special"
	NameOrdinaryService = "Standard"
	NameSpecialService  = "Special"
	LimitWeightS        = 5.0
	LimitWeightM        = 15.0
	LimitWeightL        = 25.0
	CreateStatus        = "creado"
	PickupStatus        = "recolectado"
	StationStatus       = "en_estacion"
	OnWayStatus         = "en_ruta"
	DeliveredStatus     = "entregada"
	CancelStatus        = "cancelado"
	GeneralStatus       = "Done"
	GenericName         = "unknown"
	GenericHeader       = "Authorization"
	AdminRole           = "Admin"
)

var StatusList = map[string]bool{
	"creado":      true,
	"recolectado": true,
	"en_estacion": true,
	"en_ruta":     true,
	"entregada":   true,
	"cancelado":   true,
}
