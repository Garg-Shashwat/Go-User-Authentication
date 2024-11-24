package logoutControllers

import (
	"net/http"

	"github.com/Garg-Shashwat/Go-User-Authentication/root/services"
	"github.com/labstack/echo/v4"
)

// LogoutUser logs user out and revokes the token
func LogoutUser(ctx echo.Context) error {
	userID := ctx.Get("userID").(uint)
	err := services.LogoutUser(userID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "Successfully logged out")
}