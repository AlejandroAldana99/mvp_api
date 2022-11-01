package errors

import (
	"fmt"
	"net/http"

	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/labstack/echo/v4"
)

const (
	// UserNotFound :
	UserNotFound = iota + 1
	// DataSourceException :
	DataSourceException
	// InvalidParameters :
	InvalidParameters
)

// ServiceErrors :
var ServiceErrors map[int]string = map[int]string{
	UserNotFound:        "User not found",
	DataSourceException: "Data source exception",
	InvalidParameters:   "Invalid parameters",
}

// NewAPIErrorResponse :
func NewAPIErrorResponse(errors ...models.ErrorResponse) models.APIErrorResponse {
	return models.APIErrorResponse{
		Errors: errors,
	}
}

// MapErrorCode :
func MapErrorCode(code int) string {
	return ServiceErrors[code]
}

// ErrorCodeString :
func ErrorCodeString(code int) string {
	return fmt.Sprintf("CDS-%d", code)
}

func HandleServiceError(err error) error {
	var (
		status, code int
	)
	switch err.Error() {
	case "invalid parameters":
		status = http.StatusBadRequest
		code = InvalidParameters
		break
	case "mongo: no documents in result":
		status = http.StatusNotFound
		code = UserNotFound
		break
	default:
		status = http.StatusInternalServerError
		code = DataSourceException
	}
	return echo.NewHTTPError(
		status,
		NewAPIErrorResponse(
			models.ErrorResponse{
				Code:    ErrorCodeString(code),
				Message: MapErrorCode(code),
			},
		))
}
