package repositories

import (
	"context"
	"time"

	"github.com/AlejandroAldana99/mvp_api/config"
	"github.com/AlejandroAldana99/mvp_api/constants"
	"github.com/AlejandroAldana99/mvp_api/errors"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderData struct {
	Config          config.Configuration
	MongoCollection *mongo.Collection
}

func (repo OrderData) GetOrder(orderID string) (models.OrderData, error) {
	t := time.Now()
	var order models.OrderData
	cErr := repo.MongoCollection.FindOne(context.TODO(), bson.D{{"orderid", orderID}}).
		Decode(&order)

	if cErr != nil {
		logger.Error("repositories", "GetOrderData", cErr.Error())
		return models.OrderData{}, errors.HandleServiceError(cErr)
	}

	logger.Performance("repository", "GetCandidateData", t)

	return order, nil
}

func (repo OrderData) CreateOrder(data models.OrderData) error {

	_, err := repo.MongoCollection.InsertOne(context.TODO(), data)
	if err != nil {
		logger.Error("repositories", "CreateOrder", err.Error())
		return errors.HandleServiceError(err)
	}

	return nil
}

func (repo OrderData) RefreshCandidateData(data models.OrderData) error {
	//var err error
	//result := repo.MongoCollection.FindOne(context.TODO(), bson.D{{"userid", data.UserID}})

	_, dErr := repo.MongoCollection.DeleteOne(context.TODO(), bson.D{{"userid", data.OrderID}})

	if dErr != nil {
		logger.Error("repositories", "RefreshProfileCandidateData", dErr.Error())
		return errors.HandleServiceError(dErr)
	}

	_, iErr := repo.MongoCollection.InsertOne(context.TODO(), data)

	if iErr != nil {
		logger.Error("repositories", "RefreshProfileCandidateData", iErr.Error())
		return errors.HandleServiceError(iErr)
	}

	return nil
}

func completeSize(data models.PackageData) models.PackageData {
	if data.Weight < constants.LimitWeightS {
		data.Size = constants.MeasureS
		return data
	} else if data.Weight < constants.LimitWeightM {
		data.Size = constants.MeasureM
		return data
	} else if data.Weight < constants.LimitWeightL {
		data.Size = constants.MeasureL
		return data
	} else {
		data.Size = constants.MeasureSpecial
		data.Service = constants.NameSpecialService
	}

	return data
}
