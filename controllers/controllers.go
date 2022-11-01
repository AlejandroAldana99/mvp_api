package controllers

import (
	"net/http"
	"strings"

	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/AlejandroAldana99/mvp_api/services"
	"github.com/labstack/echo/v4"
)

type ControllerData struct {
	ServiceOrder services.IOrderService
	ServiceUser  services.IUserService
}

func (controller ControllerData) GetOrderData(c echo.Context) error {
	orderID := strings.ToLower(c.Param("id"))
	data, err := controller.ServiceOrder.GetOrder(orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, data)
}

func (controller ControllerData) CreateOrderData(c echo.Context) error {
	dto := c.Get("dto").(models.OrderData)
	err := controller.ServiceOrder.CreateOrder(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Success")
}

func (controller ControllerData) UpdateOrderStatus(c echo.Context) error {
	orderID := strings.ToLower(c.QueryParam("orderid"))
	status := strings.ToLower(c.QueryParam("status"))
	err := controller.ServiceOrder.UpdateOrderStatus(orderID, status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Success")
}

func (controller ControllerData) GetUserData(c echo.Context) error {
	userID := strings.ToLower(c.Param("id"))
	data, err := controller.ServiceUser.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, data)
}

func (controller ControllerData) CreateUserData(c echo.Context) error {
	dto := c.Get("dto").(models.UserData)
	err := controller.ServiceUser.CreateUser(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Success")
}

func (controller ControllerData) Login(c echo.Context) error {
	user := strings.ToLower(c.QueryParam("user"))
	password := strings.ToLower(c.QueryParam("password"))
	login, err := controller.ServiceUser.Login(user, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, login)
}
