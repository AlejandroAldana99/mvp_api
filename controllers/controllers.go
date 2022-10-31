package controllers

import (
	"net/http"
	"strings"

	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/AlejandroAldana99/mvp_api/services"
	"github.com/labstack/echo/v4"
)

// CandidateData :
type ControllerData struct {
	ServiceOrder services.IOrderService
	ServiceUser  services.IUserService
}

// GetCandidateData :
func (controller ControllerData) GetOrderData(c echo.Context) error {
	orderID := strings.ToLower(c.Param("orderID"))
	data, err := controller.ServiceOrder.GetOrder(orderID)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (controller ControllerData) CreateOrderData(c echo.Context) error {
	dto := c.Get("dto").(models.OrderData)
	err := controller.ServiceOrder.CreateOrder(dto)
	if resp.StatusCode == http.StatusCreated {
		return c.JSON(resp.StatusCode, resp.Success)
	}
	return c.JSON(resp.StatusCode, resp.Errors)
}

// RefreshCandidateData :
func (controller ControllerData) UpdateOrderStatus(c echo.Context) error {
	orderID := strings.ToLower(c.QueryParam("orderid"))
	status := strings.ToLower(c.QueryParam("status"))
	err := controller.ServiceOrder.UpdateOrderStatus(orderID, status)
	if resp.StatusCode == http.StatusCreated {
		return c.JSON(resp.StatusCode, resp.Success)
	}
	return c.JSON(resp.StatusCode, resp.Errors)
}

// GetCandidateData :
func (controller ControllerData) GetUserData(c echo.Context) error {
	userID := strings.ToLower(c.Param("userID"))
	data, err := controller.ServiceUser.GetUser(userID)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (controller ControllerData) CreateUserData(c echo.Context) error {
	dto := c.Get("dto").(models.UserData)
	err := controller.ServiceUser.CreateUser(dto)
	if resp.StatusCode == http.StatusCreated {
		return c.JSON(resp.StatusCode, resp.Success)
	}
	return c.JSON(resp.StatusCode, resp.Errors)
}

func (controller ControllerData) Login(c echo.Context) error {
	email := strings.ToLower(c.QueryParam("email"))
	password := strings.ToLower(c.QueryParam("password"))
	err := controller.ServiceOrder.UpdateOrderStatus(orderID, status)
	if resp.StatusCode == http.StatusCreated {
		return c.JSON(resp.StatusCode, resp.Success)
	}
	return c.JSON(resp.StatusCode, resp.Errors)
}
