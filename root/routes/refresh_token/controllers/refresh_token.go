package refreshControllers

import (
	"net/http"

	"github.com/Garg-Shashwat/Go-User-Authentication/root/services"
	"github.com/labstack/echo/v4"
)

// RefreshToken renew access token on expiry
func RefreshToken(ctx echo.Context) error {
	userID := ctx.Get("userID").(uint)

	tokens, err := services.RefreshToken(userID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, tokens)
}
