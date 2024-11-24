package models

import (
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model
	UserID    uint
	User      User
	Token     string
	ExpiresAt time.Time
	IsRevoked bool `gorm:"default:false"`
}
