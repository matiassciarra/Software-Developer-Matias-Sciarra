package models

import (
	"time"

	"gorm.io/gorm"
)

type DbItem struct {
	gorm.Model           // Esto incluye campos ID, CreatedAt, UpdatedAt y DeletedAt
	Ticker     string    `gorm:"type:varchar(100);not null"`
	TargetFrom float32   `gorm:;not null"`
	TargetTo   float32   `gorm:"not null"`
	Company    string    `gorm:"not null`
	Action     string    `gorm:"not null`
	Brokerage  string    `gorm:"not null`
	RatingFrom string    `gorm:"not null`
	RatingTo   string    `gorm:"not null`
	Time       time.Time `gorm:"not null`
}
