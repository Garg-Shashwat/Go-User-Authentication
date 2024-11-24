package loginControllers

import (
	"net/http"

	loginSerializers "github.com/Garg-Shashwat/Go-User-Authentication/root/routes/login/serializers"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/services"
	"github.com/labstack/echo/v4"
)

// LoginUser handles user authentications and token generation
func LoginUser(ctx echo.Context) error {
	loginRequest, err := loginSerializers.BindLoginRequest(ctx)
	if err != nil {
		return err
	}

	tokens, err := services.LoginUser(loginRequest)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, tokens)
}
