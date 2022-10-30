package services

import "github.com/AlejandroAldana99/mvp_api/models"

type IUserService interface {
	GetUser(userID string) (models.UserData, error)
	CreateUser(data models.UserData) error
}
