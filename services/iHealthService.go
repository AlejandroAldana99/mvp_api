package services

import "github.com/AlejandroAldana99/mvp_api/models"

type IHealthCandidateData interface {
	CheckPod(chanHealth chan models.HealthComponentDetail)
}
