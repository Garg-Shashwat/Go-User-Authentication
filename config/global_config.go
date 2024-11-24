package config

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Config struct {
	env       *viper.Viper
	database  *gorm.DB
	validator *validator.Validate
}

var globalConfigInstance Config
var configInstanceOnce sync.Once

// InitiateConfig initialises the global config instance
func InitiateConfig() {
	configInstanceOnce.Do(func() {
		env := loadConfig()
		db := loadDB(env)
		globalConfigInstance = Config{
			env:       env,
			database:  db,
			validator: validator.New(),
		}
	})
}

// GetEnv returns environment instance
func GetEnv() *viper.Viper {
	return globalConfigInstance.env
}

// GetDB returns database instance
func GetDB() *gorm.DB {
	return globalConfigInstance.database
}

// GetValidator returns validator instance
func GetValidator() *validator.Validate {
	return globalConfigInstance.validator
}
