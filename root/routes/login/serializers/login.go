package loginSerializers

import (
	"net/http"

	"github.com/Garg-Shashwat/Go-User-Authentication/config"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/utils"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// BindLoginRequest deserializes the data to LoginRequest
func BindLoginRequest(ctx echo.Context) (LoginRequest, error) {
	var loginRequest LoginRequest
	if err := ctx.Bind(&loginRequest); err != nil {
		return loginRequest, utils.GetHTTPError(
			http.StatusBadRequest,
			"Could not parse data",
			err.Error(),
		)
	}

	if err := config.GetValidator().Struct(loginRequest); err != nil {
		return loginRequest, utils.GetHTTPError(
			http.StatusBadRequest,
			"Invalid data",
			err.Error(),
		)
	}

	return loginRequest, nil
}
