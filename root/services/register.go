package services

import (
	"errors"
	"net/http"

	"github.com/Garg-Shashwat/Go-User-Authentication/root/models"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/repositories"
	registerSerializers "github.com/Garg-Shashwat/Go-User-Authentication/root/routes/register/serializers"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/utils"
	"gorm.io/gorm"
)

// RegisterUser checks and registers user in system
func RegisterUser(registerRequest registerSerializers.RegisterRequest) (*models.User, error) {
	if usrRecord, err := repositories.GetUserByEmail(registerRequest.Email); usrRecord != nil {
		return nil, utils.GetHTTPError(
			http.StatusBadRequest,
			"User already registered",
			"User already registered",
		)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.GetHTTPError(
			http.StatusInternalServerError,
			"Error while finding User",
			err.Error(),
		)
	}

	hashedPassword, err := utils.HashPassword(registerRequest.Password)
	if err != nil {
		return nil, utils.GetHTTPError(
			http.StatusInternalServerError,
			"Error while hashing password",
			err.Error(),
		)
	}

	user := models.User{
		Email:    registerRequest.Email,
		Password: hashedPassword,
	}

	if err := repositories.CreateUser(&user); err != nil {
		return nil, utils.GetHTTPError(
			http.StatusInternalServerError,
			"Error while creating User",
			err.Error(),
		)
	}

	return &user, nil
}
