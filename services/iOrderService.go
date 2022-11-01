package services

import "github.com/AlejandroAldana99/mvp_api/models"

type IOrderService interface {
	GetOrder(orderID string) (models.OrderData, error)
	CreateOrder(data models.OrderData) error
	UpdateOrderStatus(orderID string, status string, role string, userID string) error
}
