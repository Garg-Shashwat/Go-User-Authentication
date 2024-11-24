package config

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	validator *validator.Validate
}

var globalConfigInstance Config
var configInstanceOnce sync.Once

func InitiateConfig() {
	configInstanceOnce.Do(func() {
		globalConfigInstance = Config{
			validator: validator.New(),
		}
	})
}

func Validator() *validator.Validate {
	return globalConfigInstance.validator
}
