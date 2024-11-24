package repositories

import (
	"time"

	"github.com/Garg-Shashwat/Go-User-Authentication/config"
	"github.com/Garg-Shashwat/Go-User-Authentication/constants"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/models"
	"github.com/google/uuid"
)

// RevokePreviousTokens revokes old refresh token
func RevokePreviousTokens(userID uint) error {
	var tokenCount int64

	err := config.GetDB().Model(&models.UserToken{}).
		Where("user_id = ? AND is_revoked = false", userID).
		Count(&tokenCount).Error
	if err != nil {
		return err
	}

	if tokenCount != 0 {
		return config.GetDB().Model(&models.UserToken{}).
			Where("user_id = ? AND is_revoked = false", userID).
			Update("is_revoked", true).Error
	}

	return nil
}

// CreateRefreshToken generates a refresh token and stores it in the DB.
func CreateRefreshToken(userID uint) (string, error) {
	token := uuid.New().String()

	userToken := models.UserToken{
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(constants.RefreshExpiry * time.Hour),
		IsRevoked: false,
	}

	if err := config.GetDB().Create(&userToken).Error; err != nil {
		return "", err
	}

	return token, nil
}
