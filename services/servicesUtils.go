package services

import (
	"time"

	"github.com/AlejandroAldana99/mvp_api/constants"
	"github.com/AlejandroAldana99/mvp_api/models"
)

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

func compareTime(start time.Time, orderTime time.Time) bool {
	end := start.Add(time.Minute * 2)
	return orderTime.After(start) && orderTime.Before(end)
}

func compareStatus(orderStatus string) bool {
	return orderStatus != constants.OnWayStatus && orderStatus != constants.DeliveredStatus
}
