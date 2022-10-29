package models

import "time"

type OrderData struct {
	OrderID         string        `json:"orderid"`
	Owner           string        `json:"owner"`
	OwnerID         string        `json:"ownerid"`
	TimeRegistry    time.Time     `json:"timeregister"`
	Status          string        `json:"status"`
	Address         AddressData   `json:"addred"`
	AddressDelivery AddressData   `json:"addressdelivery"`
	Package         []PackageData `json:"package"`
}
