package routes

import (
	"net/http"

	"github.com/AlejandroAldana99/mvp_api/controllers"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/middleware"
	"github.com/AlejandroAldana99/mvp_api/server/di"
	"github.com/labstack/echo/v4"
)

// Route represents the route structure for the service
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc echo.HandlerFunc
}

// Routes represents a route collection
type Routes []Route

// ServiceRoutes is a route collection for service handling
var ServiceRoutes Routes

func init() {
	controllersProvider := di.BuildContainer()
	err := controllersProvider.Invoke(func(
		healthController *controllers.HealthController,
		candidateData *controllers.CandidateData,
	) {
		ServiceRoutes = Routes{
			Route{
				Method:      http.MethodGet,
				Pattern:     "/health",
				HandlerFunc: healthController.HealthCheck,
				Name:        "HealthCheck",
			},
			Route{
				Method:      http.MethodGet,
				Pattern:     "/health/dependencies",
				HandlerFunc: healthController.HealthCheckDependencies,
				Name:        "HealthCheckDependencies",
			},
			Route{
				Method:      http.MethodGet,
				Pattern:     "/mvp/order/:orderID",
				HandlerFunc: middleware.ParamsValidator(candidateData.GetCandidateData),
				Name:        "GetOrder",
			},
			Route{
				Method:      http.MethodPost,
				Pattern:     "/mvp/order",
				HandlerFunc: middleware.ParamsValidator(candidateData.RefreshCandidateData),
				Name:        "CreateOrder",
			},
			Route{
				Method:      http.MethodPatch,
				Pattern:     "/mvp/order/:orderID",
				HandlerFunc: middleware.ParamsValidator(candidateData.RefreshCandidateData),
				Name:        "UpdateStatus",
			},
			Route{
				Method:      http.MethodDelete,
				Pattern:     "/mvp/oder",
				HandlerFunc: middleware.ParamsValidator(candidateData.RefreshCandidateData),
				Name:        "CancelOrder",
			},
			Route{
				Method:      http.MethodPost,
				Pattern:     "/mvp/user",
				HandlerFunc: middleware.ParamsValidator(candidateData.RefreshCandidateData),
				Name:        "CreateUser",
			},
			Route{
				Method:      http.MethodPost,
				Pattern:     "/mvp/login",
				HandlerFunc: middleware.ParamsValidator(candidateData.RefreshCandidateData),
				Name:        "Login",
			},
		}
	})

	if err != nil {
		logger.Error("routes", "init", err.Error())
	}
}
