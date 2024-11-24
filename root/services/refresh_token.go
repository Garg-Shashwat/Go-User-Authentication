package services

import "github.com/Garg-Shashwat/Go-User-Authentication/root/utils"

// RefreshToken renews acces token in case of expiry
func RefreshToken(userID uint) (map[string]string, error) {
	accessToken, err := utils.GenerateJWTToken(userID)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token": accessToken,
	}, nil
}
