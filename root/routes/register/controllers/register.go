package controllers

import (
	"net/http"

	"github.com/Garg-Shashwat/Go-User-Authentication/root/routes/register/serializers"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/services"
	"github.com/labstack/echo/v4"
)

// RegisterUser handles user registration
func RegisterUser(ctx echo.Context) error {
	registerRequest, err := serializers.BindRegisterRequest(ctx)
	if err != nil {
		return err
	}

	_, err = services.RegisterUser(registerRequest)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, "User Registered Successfully")
}
