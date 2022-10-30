package services

import "github.com/AlejandroAldana99/mvp_api/models"

type IOrderService interface {
	GetOrder(userID string) (models.OrderData, error)
	CreateOrder(data models.OrderData) error
}
