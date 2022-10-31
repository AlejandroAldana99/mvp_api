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

type OrderRepository struct {
	Config  config.Configuration
	MongoDB *mongo.Database
}

func (repo OrderRepository) GetOrder(orderID string) (models.OrderData, error) {
	t := time.Now()
	var order models.OrderData
	err := repo.MongoDB.Collection("orders").FindOne(context.TODO(), bson.D{{Key: "orderid", Value: orderID}}).
		Decode(&order)

	if err != nil {
		logger.Error("repositories", "GetOrderData", err.Error())
		return models.OrderData{}, errors.HandleServiceError(err)
	}

	logger.Performance("repository", "GetOrder", t)

	return order, nil
}

func (repo OrderRepository) CreateOrder(data models.OrderData) error {

	t := time.Now()
	_, err := repo.MongoDB.Collection("orders").InsertOne(context.TODO(), data)
	if err != nil {
		logger.Error("repositories", "CreateOrder", err.Error())
		return errors.HandleServiceError(err)
	}

	logger.Performance("repository", "CreateOrder", t)

	return nil
}

func (repo OrderRepository) UpdateOrderStatus(orderID string, status string) error {
	filter := bson.D{{Key: "orderid", Value: orderID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "status", Value: status},
	}}}

	_, err := repo.MongoDB.Collection("orders").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		logger.Error("repositories", "UpdateOrderStatus", err.Error())
		return errors.HandleServiceError(err)
	}

	return nil
}
