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
}

// GetCandidateData :
func (controller ControllerData) GetOrderData(c echo.Context) error {
	userID := strings.ToLower(c.Param("orderID"))
	data, err := controller.ServiceOrder.GetCandidateData(userID)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

// RefreshCandidateData :
func (controller ControllerData) RefreshCandidateData(c echo.Context) error {
	var data models.CandidateData
	userID := c.Param("userID")

	if e := c.Bind(&data); e != nil {
		return e
	}

	err := controller.ServiceOrder.RefreshCandidateData(userID, data)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}
