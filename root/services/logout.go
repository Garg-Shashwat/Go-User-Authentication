package services

import (
	"net/http"

	"github.com/Garg-Shashwat/Go-User-Authentication/root/repositories"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/utils"
)

// LogoutUser logs user out
func LogoutUser(userID uint) error {
	err := repositories.RevokePreviousTokens(userID)
	if err != nil {
		return utils.GetHTTPError(
			http.StatusInternalServerError,
			"Error while revoking Token",
			err.Error(),
		)
	}

	return nil
}
