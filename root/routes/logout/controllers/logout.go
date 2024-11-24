package logoutControllers

import (
	"net/http"

	"github.com/Garg-Shashwat/Go-User-Authentication/root/services"
	"github.com/labstack/echo/v4"
)

// LogoutUser logs user out and revokes the token
func LogoutUser(ctx echo.Context) error {
	userIDFloat64, ok := ctx.Get("userID").(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	userID := uint(userIDFloat64)
	err := services.LogoutUser(userID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "Successfully logged out")
}
