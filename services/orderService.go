package services

import (
	e "errors"
	"time"

	"github.com/AlejandroAldana99/mvp_api/constants"
	"github.com/AlejandroAldana99/mvp_api/errors"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/AlejandroAldana99/mvp_api/repositories"
)

const millisecondsEq = 1000000.0

type OrderService struct {
	Repository repositories.IOrderRepository
}

func (service OrderService) GetOrder(orderID string) (models.OrderData, error) {
	order, err := service.Repository.GetOrder(orderID)
	if err != nil {
		logger.Error("services", "GetOrder", err.Error())
		return order, errors.HandleServiceError(err)
	}

	return order, nil
}

func (service OrderService) CreateOrder(data models.OrderData) error {
	// Add time
	data.TimeRegistry = time.Now()
	// Valid status
	if !constants.StatusList[data.Status] {
		err := e.New("Invalid Status")
		logger.Error("services", "UpdateOrderStatus", err.Error())
		return errors.HandleServiceError(err)
	}
	// Valid coordinates
	if !validCoordinates(data.Coordinates.Latitude, data.Coordinates.Longitude) {
		err := e.New("Invalid Coordinates")
		logger.Error("services", "UpdateOrderStatus", err.Error())
		return errors.HandleServiceError(err)
	}
	// Read packages and bussines rules
	for i := range data.Packages {
		sp := false
		data.Packages[i], sp = completeSize(data.Packages[i])
		if sp {
			data.Description = "Special Package: Special deal is required."
		}
	}

	err := service.Repository.CreateOrder(data)
	if err != nil {
		logger.Error("services", "CreateOrder", err.Error())
		return errors.HandleServiceError(err)
	}

	return nil
}

func (service OrderService) UpdateOrderStatus(orderID string, status string, role string) error {
	if !constants.StatusList[status] {
		err := e.New("Invalid Status")
		logger.Error("services", "UpdateOrderStatus", err.Error())
		return errors.HandleServiceError(err)
	}

	// Cancelation statement
	if status == constants.CancelStatus {
		data, dErr := service.Repository.GetOrder(orderID)
		if dErr != nil {
			logger.Error("services", "CancelOrder", dErr.Error())
			return errors.HandleServiceError(dErr)
		}

		if !compareTime(data.TimeRegistry, time.Now()) && !compareStatus(data.Status) {
			err := e.New("Non-cancellable Order")
			logger.Error("services", "CancelOrder", err.Error())
			return errors.HandleServiceError(err)
		}

		err := service.Repository.UpdateOrderStatus(orderID, status)
		if err != nil {
			logger.Error("services", "CancelOrder", err.Error())
			return errors.HandleServiceError(err)
		}
		// Update statement
	} else {
		if role != constants.AdminRole {
			err := e.New("Invalid Role")
			logger.Error("services", "UpdateOrderStatus", err.Error())
			return errors.HandleServiceError(err)
		}
		err := service.Repository.UpdateOrderStatus(orderID, status)
		if err != nil {
			logger.Error("services", "UpdateOrderStatus", err.Error())
			return errors.HandleServiceError(err)
		}
	}

	return nil
}
