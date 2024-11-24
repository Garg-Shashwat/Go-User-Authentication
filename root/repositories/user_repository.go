package repositories

import (
	"github.com/Garg-Shashwat/Go-User-Authentication/config"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/models"
)

// CreateUser creates a new user in the database
func CreateUser(user *models.User) error {
	return config.GetDB().Create(user).Error
}

// GetUserByEmail fetches a user by email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := config.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
