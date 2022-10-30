package services

import (
	"github.com/AlejandroAldana99/mvp_api/errors"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/AlejandroAldana99/mvp_api/repositories"
)

type UserService struct {
	Repository repositories.IUserRepository
}

// GetCandidateData :
func (service UserService) Getuser(userID string) (models.UserData, error) {
	var user models.UserData
	var err error
	user, err = service.Repository.GetUser(userID)

	if err != nil {
		logger.Error("services", "GetUser", err.Error())
		return user, errors.HandleServiceError(err)
	}

	return user, nil
}

func (service UserService) Createuser(data models.UserData) error {
	err := service.Repository.CreateUser(data)
	if err != nil {
		logger.Error("services", "Createuser", err.Error())
		return errors.HandleServiceError(err)
	}

	return nil
}
