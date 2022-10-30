package services

import (
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