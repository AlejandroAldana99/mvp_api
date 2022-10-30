package services

import (
	"github.com/AlejandroAldana99/mvp_api/errors"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/AlejandroAldana99/mvp_api/repositories"
)

const millisecondsEq = 1000000.0

// CandidateData :
type OrderService struct {
	Repository repositories.IOrderRepository
}

// GetCandidateData :
func (service OrderService) GetOrder(orderID string) (models.OrderData, error) {
	var order models.OrderData
	var err error
	order, err = service.Repository.GetOrder(orderID)

	if err != nil {
		logger.Error("services", "GetOrder", err.Error())
		return order, errors.HandleServiceError(err)
	}

	return order, nil
}

func (service OrderService) CreateOrder(data models.OrderData) error {
	data.Package[0] = completeSize(data.Package[0])
	err := service.Repository.CreateOrder(data)
	if err != nil {
		logger.Error("services", "CreateOrder", err.Error())
		return errors.HandleServiceError(err)
	}

	return nil
}

func (service OrderService) UpdateOrderStatus(orderID string, status string) error {
	err := service.Repository.UpdateOrderStatus(orderID, status)
	if err != nil {
		logger.Error("services", "UpdateOrderStatus", err.Error())
		return errors.HandleServiceError(err)
	}

	return nil
}
