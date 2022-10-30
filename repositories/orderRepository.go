package repositories

import (
	"context"
	"time"

	"github.com/AlejandroAldana99/mvp_api/config"
	"github.com/AlejandroAldana99/mvp_api/errors"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderData struct {
	Config  config.Configuration
	MongoDB *mongo.Database
}

func (repo OrderData) GetOrder(orderID string) (models.OrderData, error) {
	t := time.Now()
	var order models.OrderData
	err := repo.MongoDB.Collection("orders").FindOne(context.TODO(), bson.D{{"orderid", orderID}}).
		Decode(&order)

	if err != nil {
		logger.Error("repositories", "GetOrderData", err.Error())
		return models.OrderData{}, errors.HandleServiceError(err)
	}

	logger.Performance("repository", "GetOrder", t)

	return order, nil
}

func (repo OrderData) CreateOrder(data models.OrderData) error {

	t := time.Now()
	_, err := repo.MongoDB.Collection("orders").InsertOne(context.TODO(), data)
	if err != nil {
		logger.Error("repositories", "CreateOrder", err.Error())
		return errors.HandleServiceError(err)
	}

	logger.Performance("repository", "CreateOrder", t)

	return nil
}

func (repo OrderData) UpdateOrderStatus(orderID string, status string) error {

	filter := bson.D{{"status", status}}
	update := bson.D{{"$set", bson.D{{"orderid", orderID}}}}

	_, err := repo.MongoDB.Collection("orders").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		logger.Error("repositories", "UpdateOrderStatus", err.Error())
		return errors.HandleServiceError(err)
	}

	return nil
}
