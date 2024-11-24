package registerSerializers

import (
	"net/http"

	"github.com/Garg-Shashwat/Go-User-Authentication/config"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/utils"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// BindRegisterRequest deserializes the data to RegisterRequest
func BindRegisterRequest(ctx echo.Context) (RegisterRequest, error) {
	var registerRequest RegisterRequest
	if err := ctx.Bind(&registerRequest); err != nil {
		return registerRequest, utils.GetHTTPError(
			http.StatusBadRequest,
			"Could not parse data",
			err.Error(),
		)
	}

	if err := config.GetValidator().Struct(registerRequest); err != nil {
		return registerRequest, utils.GetHTTPError(
			http.StatusBadRequest,
			"Invalid data",
			err.Error(),
		)
	}

	return registerRequest, nil
}
