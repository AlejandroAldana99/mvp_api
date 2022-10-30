package middleware

import (
	"errors"
	"strings"

	er "github.com/AlejandroAldana99/mvp_api/errors"
	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/go-playground/validator"

	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

// ParamsValidator :
func ParamsValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) (e error) {
		userID := strings.TrimSpace(context.Param("userID"))

		if userID == "undefined" || userID == "null" || userID == "" {
			logger.Error("middleware", "ParamsValidator", "Instance param cannot be null")
			return er.HandleServiceError(errors.New("invalid parameters"))
		}

		e = next(context)
		return e
	}

}

//ValidateInputToSendGrid :
func ValidateOrderBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (e error) {
		dto, err := orderBodyToStruct(c)
		if err != nil {
			logger.Error("middleware", "ValidateBody", err.Error())
			return er.HandleServiceError(errors.New("invalid body"))
		}
		errValidation := validateOrderModel(dto)
		if errValidation != nil {
			return er.HandleServiceError(errValidation)
		}

		c.Set("dto", dto)
		return next(c)
	}
}

//SendgridPayloadToStruct :
func orderBodyToStruct(c echo.Context) (models.OrderData, error) {
	dto := new(models.OrderData)
	err := c.Bind(dto)
	return *dto, err
}

//ValidateInputModelSendgrid :
func validateOrderModel(dto models.OrderData) error {
	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errs.Errors = AppendErrors(validationErrors)
	}
	return errs
}
