package services

import (
	"net/http"
	"time"

	"github.com/AlejandroAldana99/mvp_api/config"
	"github.com/AlejandroAldana99/mvp_api/errors"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/AlejandroAldana99/mvp_api/repositories"
)

const millisecondsEq = 1000000.0

type IHealthCandidateData interface {
	CheckPod(chanHealth chan models.HealthComponentDetail)
}

type HealthCandidateData struct{}

// CandidateData :
type CandidateData struct {
	Repository repositories.IOrderRepository
}

// GetCandidateData :
func (service CandidateData) GetCandidateData(userID string) (models.CandidateData, error) {
	var data models.CandidateData
	var err error
	data, err = service.Repository.GetCandidateData(userID)

	if err != nil {
		var sqlErr error
		data, sqlErr = service.Repository.GetCandidateDataFromSQL(userID)

		if sqlErr != nil {
			logger.Error("services", "GetCandidateData", sqlErr.Error())
			return data, errors.HandleServiceError(sqlErr)
		}
		go service.RefreshCandidateData(userID, data)
		return data, err
	}
	data.CalculateAge()
	data.BuildPhotoURL(config.GetConfig().S3URL)
	return data, nil
}

// RefreshCandidateData :
func (service CandidateData) RefreshCandidateData(userID string, data models.CandidateData) error {

	data.UserID = userID

	err := service.Repository.RefreshCandidateData(data)
	if err != nil {
		return err
	}

	return nil
}

func updateHealthDetails(details *models.HealthComponentDetail, now time.Time, statusCode int) {
	details.Status = "fail"
	if statusCode == http.StatusOK {
		details.Status = "pass"
	}
	details.MetricValue = float32(time.Since(now).Nanoseconds()) / millisecondsEq
	details.MetricUnit = "ms"
}

// CheckPod :
func (service *HealthCandidateData) CheckPod(chanHealth chan models.HealthComponentDetail) {
	now := time.Now()
	details := models.HealthComponentDetail{
		Name: "instance",
		Type: "pod",
		Time: now,
	}

	statusCode := 200
	updateHealthDetails(&details, now, statusCode)
	chanHealth <- details
}
