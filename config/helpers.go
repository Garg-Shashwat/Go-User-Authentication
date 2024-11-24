package config

import (
	"fmt"
	"log"

	"github.com/Garg-Shashwat/Go-User-Authentication/constants"
	"github.com/Garg-Shashwat/Go-User-Authentication/migrations"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// loadConfig loads configurations for project
func loadConfig() *viper.Viper {
	environment := viper.New()
	environment.SetConfigFile(constants.EnvFileName)
	environment.SetConfigType("env")

	environment.WatchConfig()
	err := environment.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	return environment
}

// loadDB loads and connects with DB
func loadDB(env *viper.Viper) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.GetString("db_host"),
		env.GetString("db_user"),
		env.GetString("db_password"),
		env.GetString("db_name"),
		env.GetString("db_port"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Printf("DB connected successfully!")

	return db
}
