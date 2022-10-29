package di

import (
	"github.com/AlejandroAldana99/mvp_api/config"
	"github.com/AlejandroAldana99/mvp_api/controllers"
	"github.com/AlejandroAldana99/mvp_api/database"
	"github.com/AlejandroAldana99/mvp_api/repositories"
	"github.com/AlejandroAldana99/mvp_api/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

// Repositories
func newCandidateDataRepository(collection *mongo.Collection) *repositories.CandidateData {
	return &repositories.CandidateData{
		MongoCollection: collection,
		Config:          config.GetConfig(),
	}
}

// Services
func newCandidateDataService(repo *repositories.CandidateData) *services.CandidateData {
	return &services.CandidateData{Repository: repo}
}

// Controllers
func newCandidateDataController(service *services.CandidateData) *controllers.CandidateData {
	return &controllers.CandidateData{
		Service: service,
	}
}

func newHealthController(client *mongo.Client, service *services.HealthCandidateData) *controllers.HealthController {
	return &controllers.HealthController{
		Configuration: config.GetConfig(),
		MongoClient:   client,
		ServiceHealth: service,
	}
}

func newHealthUserCandidateService() *services.HealthCandidateData {
	return &services.HealthCandidateData{}
}

// BuildContainer :
func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(database.NewMongoDBClient)
	_ = container.Provide(database.NewMongoCollection)
	_ = container.Provide(newCandidateDataRepository)
	_ = container.Provide(newHealthUserCandidateService)
	_ = container.Provide(newCandidateDataService)
	_ = container.Provide(newHealthController)
	_ = container.Provide(newCandidateDataController)

	return container
}
