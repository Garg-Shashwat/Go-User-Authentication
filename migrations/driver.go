package migrations

import (
	"github.com/Garg-Shashwat/Go-User-Authentication/root/models"
	"gorm.io/gorm"
)

// Migrate drives all DB migrations
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{},
		&models.UserToken{},
	); err != nil {
		return err
	}

	return nil
}
