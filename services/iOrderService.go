package services

import "github.com/AlejandroAldana99/mvp_api/models"

type IOrderService interface {
	GetCandidateData(userID string) (models.CandidateData, error)
	RefreshCandidateData(userID string, data models.CandidateData) error
}
