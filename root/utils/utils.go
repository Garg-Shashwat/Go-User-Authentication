package utils

import (
	"github.com/labstack/echo/v4"
)

// GetHTTPError parses the http code, error and message
func GetHTTPError(httpCode int, message string, error string) error {
	return echo.NewHTTPError(
		httpCode,
		map[string]string{
			"message": message,
			"error":   error,
		},
	)
}
