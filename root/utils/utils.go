package utils

import (
	"time"

	"github.com/Garg-Shashwat/Go-User-Authentication/config"
	"github.com/Garg-Shashwat/Go-User-Authentication/constants"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// GetHTTPError parses the http code, error and message
func GetHTTPError(httpCode int, message string, error string) error {
	return echo.NewHTTPError(
		httpCode,
		map[string]string{
			"message": message,
			"error":   error,
		},
	)
}

// HashPassword creates a password hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash verifies password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWTToken generates Accesstokens for user
func GenerateJWTToken(userID uint) (string, error) {
	currTime := time.Now()
	expiryTime := currTime.Add(constants.AccessExpiry * time.Hour)
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": expiryTime.Unix(),
		"iat": currTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetEnv().GetString("JWT_SECRET")))
}
