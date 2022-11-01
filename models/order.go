package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderData struct {
	OrderID         MyObjectID      `json:"orderid,omitempty" bson:"_id,omitempty"`
	Owner           string          `json:"owner"`
	OwnerID         string          `json:"ownerid,omitempty"`
	TimeRegistry    time.Time       `json:"timeregister,omitempty"`
	Status          string          `json:"status"`
	Address         AddressData     `json:"addred"`
	AddressDelivery AddressData     `json:"addressdelivery"`
	Packages        []PackageData   `json:"package"`
	Coordinates     CoordinatesData `json:"coordinates"`
	Description     string          `json:"description,omitempty"`
}

type MyObjectID string

func (id MyObjectID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	p, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return bsontype.Null, nil, err
	}

	return bson.MarshalValue(p)
}
