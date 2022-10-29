package middleware

import (
	"errors"
	"strings"

	er "github.com/AlejandroAldana99/mvp_api/errors"

	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/labstack/echo/v4"
)

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
