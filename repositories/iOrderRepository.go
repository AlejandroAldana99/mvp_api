package repositories

import "github.com/AlejandroAldana99/mvp_api/models"

type IOrderRepository interface {
	GetOrder(userID string) (models.OrderData, error)
	CreateOrder(data models.OrderData) error
	RefreshCandidateData(data models.CandidateData) error
}
