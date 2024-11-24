package services

import (
	"errors"
	"net/http"

	"github.com/Garg-Shashwat/Go-User-Authentication/root/models"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/repositories"
	loginSerializers "github.com/Garg-Shashwat/Go-User-Authentication/root/routes/login/serializers"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/utils"
	"gorm.io/gorm"
)

// LoginUser verfies user details and returns the tokens
func LoginUser(loginRequest loginSerializers.LoginRequest) (map[string]string, error) {
	var user *models.User
	user, err := repositories.GetUserByEmail(loginRequest.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.GetHTTPError(
			http.StatusBadRequest,
			"User not registered",
			err.Error(),
		)
	} else if err != nil {
		return nil, utils.GetHTTPError(
			http.StatusInternalServerError,
			"Error while finding User",
			err.Error(),
		)
	}

	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		return nil, utils.GetHTTPError(
			http.StatusBadRequest,
			"Invalid Credentials",
			"Password not matched",
		)
	}

	accessToken, err := utils.GenerateJWTToken(user.ID)
	if err != nil {
		return nil, err
	}

	err = repositories.RevokePreviousTokens(user.ID)
	if err != nil {
		return nil, utils.GetHTTPError(
			http.StatusInternalServerError,
			"Error while revoking Token",
			err.Error(),
		)
	}

	refreshToken, err := repositories.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, utils.GetHTTPError(
			http.StatusInternalServerError,
			"Error while generating Token",
			err.Error(),
		)
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}
